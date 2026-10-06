package bybit_connector

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// The tests in this file reproduce known WebSocket reconnect defects that this
// fork inherited from e03d95a. They are skipped unless REPRO_RACE is set, so
// `go test .` and `go test -race .` stay green. To reproduce:
//
//	REPRO_RACE=1 go test -race -run Race -v .
//
// Expected: a DATA RACE report between the b.conn assignment in Connect()
// (under writeMu) and the unlocked b.conn read in handleIncomingMessages.
func requireRepro(t *testing.T) {
	t.Helper()
	if os.Getenv("REPRO_RACE") == "" {
		t.Skip("known reconnect defects; set REPRO_RACE=1 to reproduce under -race")
	}
}

// startStuckWSServer upgrades every request, then holds the connection open
// until the test finishes, so client readers stay blocked in ReadMessage.
func startStuckWSServer(t *testing.T) string {
	t.Helper()

	done := make(chan struct{})
	upgrader := websocket.Upgrader{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		<-done
		_ = conn.Close()
	}))
	t.Cleanup(func() {
		close(done)
		srv.Close()
	})

	return "ws" + strings.TrimPrefix(srv.URL, "http")
}

// TestRaceConnectSwapsConnWhileReaderBlocked reproduces the unlocked read of
// b.conn in handleIncomingMessages racing against Connect()'s write of the
// same pointer: a second Connect() while the first reader is still blocked in
// ReadMessage on the previous connection.
func TestRaceConnectSwapsConnWhileReaderBlocked(t *testing.T) {
	requireRepro(t)

	ws := NewBybitPublicWebSocket(startStuckWSServer(t), func(string) error { return nil })
	if ws.Connect() == nil {
		t.Fatal("first connect failed")
	}
	// Let the reader evaluate b.conn and block inside ReadMessage.
	time.Sleep(200 * time.Millisecond)

	if ws.Connect() == nil {
		t.Fatal("second connect failed")
	}
	t.Cleanup(func() { _ = ws.Disconnect() })
	time.Sleep(200 * time.Millisecond)
}

// TestRaceMonitorReconnectsWhileReaderAlive reproduces the internal reconnect
// path: monitorConnection() calls Connect() on its 5s tick once isConnected
// goes false, while a reader can still be alive. Connect() itself spawns a
// reader, and monitorConnection() then spawns a second one, so the pool of
// readers on a single connection grows with every cycle.
func TestRaceMonitorReconnectsWhileReaderAlive(t *testing.T) {
	requireRepro(t)

	ws := NewBybitPublicWebSocket(startStuckWSServer(t), func(string) error { return nil })
	if ws.Connect() == nil {
		t.Fatal("connect failed")
	}
	time.Sleep(200 * time.Millisecond)

	// isConnected false while the reader is still blocked in ReadMessage —
	// the state the duplicate reader of the reconnect path leaves behind.
	ws.setConnected(false)

	// monitorConnection ticks every 5s and dials by itself.
	time.Sleep(7 * time.Second)

	_ = ws.Disconnect()
}
