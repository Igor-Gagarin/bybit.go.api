package bybit_connector

import (
	"errors"
	"net"
	"testing"
	"time"
)

// Lifecycle coverage for Done, CloseError and the autoReconnect option: a
// broken connection with autoReconnect disabled must close Done and report the
// reason, with autoReconnect enabled the monitor must redial while Done stays
// open, a silent connection is ended by the read deadline, and a user-initiated
// Disconnect closes Done without an error. TestMain in
// bybit_websocket_reconnect_test.go shortens monitorTickInterval so redials
// finish in milliseconds.

func waitDone(t *testing.T, ws *WebSocket, what string) {
	t.Helper()
	select {
	case <-ws.Done():
	case <-time.After(3 * time.Second):
		t.Fatalf("timed out waiting for Done: %s", what)
	}
}

func assertDoneOpen(t *testing.T, ws *WebSocket, what string) {
	t.Helper()
	select {
	case <-ws.Done():
		t.Fatalf("Done must stay open: %s", what)
	default:
	}
}

// TestDoneReportsReadFailureWithoutAutoReconnect: killing the underlying
// connection must close Done, report the read error through CloseError and
// leave the monitor with nothing to redial.
func TestDoneReportsReadFailureWithoutAutoReconnect(t *testing.T) {
	url, conns := holdWSServer(t, "")
	ws := NewBybitPublicWebSocket(url, func(string) error { return nil }, WithAutoReconnect(false))
	if ws.Connect() == nil {
		t.Fatal("connect failed")
	}
	waitFor(t, 3*time.Second, "connection", func() bool { return conns.Load() >= 1 })

	assertDoneOpen(t, ws, "healthy connection")
	if ws.CloseError() != nil {
		t.Fatalf("CloseError on healthy connection: %v", ws.CloseError())
	}

	conn := ws.getConn()
	if conn == nil {
		t.Fatal("no connection to break")
	}
	_ = conn.Close()

	waitDone(t, ws, "read failure with autoReconnect disabled")
	if ws.CloseError() == nil {
		t.Fatal("CloseError must report the read failure")
	}

	time.Sleep(300 * time.Millisecond)
	if got := conns.Load(); got != 1 {
		t.Fatalf("autoReconnect=false must not redial, got %d server connections", got)
	}
	_ = ws.Disconnect()
}

// TestDoneStaysOpenWhileMonitorRedials: with autoReconnect enabled (the
// default) a broken connection ends the previous one with a reportable
// CloseError, but Done must stay open — the caller's selector keeps waiting
// while the monitor redials.
func TestDoneStaysOpenWhileMonitorRedials(t *testing.T) {
	url, conns := holdWSServer(t, "")
	ws := NewBybitPublicWebSocket(url, func(string) error { return nil })
	if ws.Connect() == nil {
		t.Fatal("connect failed")
	}
	waitFor(t, 3*time.Second, "connection", func() bool { return conns.Load() >= 1 })
	assertDoneOpen(t, ws, "healthy connection")

	conn := ws.getConn()
	if conn == nil {
		t.Fatal("no connection to break")
	}
	_ = conn.Close()

	waitFor(t, 3*time.Second, "redial after read failure", func() bool { return conns.Load() >= 2 })
	assertDoneOpen(t, ws, "autoReconnect redial")
	if !ws.getConnected() {
		t.Fatal("socket not connected after redial")
	}
	if ws.CloseError() == nil {
		t.Fatal("CloseError must report the failure that ended the previous connection")
	}
	_ = ws.Disconnect()
}

// TestReadIdleDeadlineEndsSilentConnection: a connection that stays silent
// past readIdleTimeout must surface a timeout through CloseError and close
// Done when autoReconnect is disabled, instead of hanging forever.
func TestReadIdleDeadlineEndsSilentConnection(t *testing.T) {
	url, conns := holdWSServer(t, "")
	ws := NewBybitPublicWebSocket(url, func(string) error { return nil }, WithAutoReconnect(false))
	ws.readIdleTimeout = 200 * time.Millisecond
	if ws.Connect() == nil {
		t.Fatal("connect failed")
	}
	waitFor(t, 3*time.Second, "connection", func() bool { return conns.Load() >= 1 })

	waitDone(t, ws, "read idle timeout")

	var ne net.Error
	if !errors.As(ws.CloseError(), &ne) || !ne.Timeout() {
		t.Fatalf("CloseError must be the read deadline error, got %v", ws.CloseError())
	}
	if got := conns.Load(); got != 1 {
		t.Fatalf("must not redial after idle timeout with autoReconnect=false, got %d server connections", got)
	}
	_ = ws.Disconnect()
}

// TestHandlerErrorEndsWithoutReconnect: with autoReconnect disabled a failing
// handler must close Done and surface its own error through CloseError.
func TestHandlerErrorEndsWithoutReconnect(t *testing.T) {
	handlerErr := errors.New("handler failed")
	url, conns := holdWSServer(t, "boom")
	ws := NewBybitPublicWebSocket(url, func(string) error { return handlerErr }, WithAutoReconnect(false))
	if ws.Connect() == nil {
		t.Fatal("connect failed")
	}

	waitDone(t, ws, "handler error with autoReconnect disabled")
	if !errors.Is(ws.CloseError(), handlerErr) {
		t.Fatalf("CloseError = %v, want the handler error", ws.CloseError())
	}

	time.Sleep(300 * time.Millisecond)
	if got := conns.Load(); got != 1 {
		t.Fatalf("autoReconnect=false must not redial after handler error, got %d server connections", got)
	}
	_ = ws.Disconnect()
}

// TestUserDisconnectClosesDoneWithoutError: Done is valid before Connect, and
// Disconnect is a deliberate close — Done fires but CloseError stays nil.
func TestUserDisconnectClosesDoneWithoutError(t *testing.T) {
	url, conns := holdWSServer(t, "")
	ws := NewBybitPublicWebSocket(url, func(string) error { return nil })

	done := ws.Done()
	select {
	case <-done:
		t.Fatal("Done must be open before Connect")
	default:
	}

	if ws.Connect() == nil {
		t.Fatal("connect failed")
	}
	waitFor(t, 3*time.Second, "connection", func() bool { return conns.Load() >= 1 })
	select {
	case <-done:
		t.Fatal("Done must stay open while connected")
	default:
	}

	_ = ws.Disconnect()

	waitDone(t, ws, "user Disconnect")
	if ws.CloseError() != nil {
		t.Fatalf("CloseError after user Disconnect: %v", ws.CloseError())
	}
}
