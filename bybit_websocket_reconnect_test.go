package bybit_connector

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// Regression tests for the WebSocket reconnect mechanics: a second Connect()
// with a live reader, monitor-driven redials, message handler failures and
// goroutine stability across reconnect cycles. TestMain shortens
// monitorTickInterval so every cycle finishes in milliseconds; the tests fail
// under `go test -race` if any of the formerly documented defects return.

func TestMain(m *testing.M) {
	monitorTickInterval = 50 * time.Millisecond
	os.Exit(m.Run())
}

// holdWSServer upgrades every request, counts accepted connections and holds
// each one open until the client closes it or the test finishes. It never
// writes, so client readers stay blocked in ReadMessage. If message is not
// empty, it is pushed once per connection before holding.
func holdWSServer(t *testing.T, message string) (string, *atomic.Int64) {
	t.Helper()

	var conns atomic.Int64
	done := make(chan struct{})
	upgrader := websocket.Upgrader{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		conns.Add(1)
		closed := make(chan struct{})
		go func() {
			select {
			case <-done:
				_ = conn.Close()
			case <-closed:
			}
		}()
		if message != "" {
			_ = conn.WriteMessage(websocket.TextMessage, []byte(message))
		}
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				close(closed)
				return
			}
		}
	}))
	t.Cleanup(func() {
		close(done)
		srv.Close()
	})

	return "ws" + strings.TrimPrefix(srv.URL, "http"), &conns
}

func waitFor(t *testing.T, timeout time.Duration, what string, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatalf("timed out waiting for %s", what)
}

// TestConnectWhileReaderBlocked does a second Connect() while the first reader
// is blocked in ReadMessage: the old connection must be closed, exactly one
// connection must be live on the server and the goroutine count must not grow.
func TestConnectWhileReaderBlocked(t *testing.T) {
	url, conns := holdWSServer(t, "")
	ws := NewBybitPublicWebSocket(url, func(string) error { return nil })
	if ws.Connect() == nil {
		t.Fatal("first connect failed")
	}
	waitFor(t, 3*time.Second, "first connection", func() bool { return conns.Load() >= 1 })
	time.Sleep(200 * time.Millisecond)

	before := runtime.NumGoroutine()
	if ws.Connect() == nil {
		t.Fatal("second connect failed")
	}
	time.Sleep(300 * time.Millisecond)

	if got := conns.Load(); got != 2 {
		t.Fatalf("expected exactly 2 server connections, got %d", got)
	}
	if !ws.getConnected() {
		t.Fatal("socket not connected after second connect")
	}
	if after := runtime.NumGoroutine(); after > before+2 {
		t.Fatalf("goroutines grew across Connect(): %d -> %d", before, after)
	}
	_ = ws.Disconnect()
}

// TestMonitorRedialsDisconnectedSocket clears isConnected while the reader is
// still alive; the monitor must redial exactly once and no more.
func TestMonitorRedialsDisconnectedSocket(t *testing.T) {
	url, conns := holdWSServer(t, "")
	ws := NewBybitPublicWebSocket(url, func(string) error { return nil })
	if ws.Connect() == nil {
		t.Fatal("connect failed")
	}
	waitFor(t, 3*time.Second, "first connection", func() bool { return conns.Load() >= 1 })

	ws.setConnected(false)
	waitFor(t, 3*time.Second, "monitor redial", func() bool { return conns.Load() >= 2 })
	time.Sleep(300 * time.Millisecond)

	if got := conns.Load(); got != 2 {
		t.Fatalf("expected exactly 1 redial, got %d server connections", got)
	}
	if !ws.getConnected() {
		t.Fatal("socket not connected after redial")
	}
	_ = ws.Disconnect()
}

// TestHandlerErrorTriggersRedial: a failing message handler must mark the
// connection disconnected instead of silently ending the reader, so the
// monitor redials.
func TestHandlerErrorTriggersRedial(t *testing.T) {
	url, conns := holdWSServer(t, "boom")
	ws := NewBybitPublicWebSocket(url, func(string) error { return errors.New("handler failed") })
	if ws.Connect() == nil {
		t.Fatal("connect failed")
	}
	waitFor(t, 3*time.Second, "redial after handler failure", func() bool { return conns.Load() >= 2 })
	_ = ws.Disconnect()
}

// TestDisconnectWithoutConnect: Disconnect and SendSubscription before any
// Connect() must return instead of panicking.
func TestDisconnectWithoutConnect(t *testing.T) {
	ws := NewBybitPublicWebSocket("ws://127.0.0.1:1", nil)
	if err := ws.Disconnect(); err != nil {
		t.Fatalf("Disconnect before Connect returned %v", err)
	}
	if _, err := ws.SendSubscription([]string{"tickers.BTCUSDT"}); err == nil {
		t.Fatal("SendSubscription before Connect should fail")
	}
}

// TestRedialCyclesKeepGoroutinesStable: repeated monitor redials must not
// accumulate readers, monitors or pings, and Disconnect must release them.
func TestRedialCyclesKeepGoroutinesStable(t *testing.T) {
	beforeConnect := runtime.NumGoroutine()

	url, conns := holdWSServer(t, "")
	ws := NewBybitPublicWebSocket(url, func(string) error { return nil })
	if ws.Connect() == nil {
		t.Fatal("connect failed")
	}
	waitFor(t, 3*time.Second, "first connection", func() bool { return conns.Load() >= 1 })
	time.Sleep(200 * time.Millisecond)

	base := runtime.NumGoroutine()
	for i := 1; i <= 5; i++ {
		ws.setConnected(false)
		want := int64(1 + i)
		waitFor(t, 3*time.Second, "redial", func() bool { return conns.Load() >= want })
		time.Sleep(150 * time.Millisecond)
		if got := conns.Load(); got != want {
			t.Fatalf("cycle %d: expected %d server connections, got %d", i, want, got)
		}
	}
	if after := runtime.NumGoroutine(); after > base+3 {
		t.Fatalf("goroutines leaked across redials: %d -> %d", base, after)
	}

	_ = ws.Disconnect()
	time.Sleep(150 * time.Millisecond)
	if end := runtime.NumGoroutine(); end > beforeConnect+2 {
		t.Fatalf("goroutines not released after Disconnect: %d -> %d", beforeConnect, end)
	}
}
