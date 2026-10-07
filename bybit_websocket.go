package bybit_connector

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type MessageHandler func(message string) error

// DebugWSRequest enables tracing of WebSocket request payloads: subscription
// args and order parameters. Request credentials are never traced. It is off
// by default, because a library should not write request payloads to stdout
// unless asked, and those payloads carry live trading parameters.
var DebugWSRequest = false

var monitorTickInterval = 5 * time.Second

func debugWSRequest(format string, a ...interface{}) {
	if DebugWSRequest {
		fmt.Println(fmt.Sprintf(format, a...))
	}
}

func (b *WebSocket) handleIncomingMessages(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			return
		}
		b.connMu.Lock()
		conn := b.conn
		changed := b.connChanged
		b.connMu.Unlock()
		if conn == nil {
			select {
			case <-ctx.Done():
				return
			case <-changed:
				continue
			}
		}
		_, message, err := conn.ReadMessage()
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			b.connMu.Lock()
			current := b.conn
			changed = b.connChanged
			b.connMu.Unlock()
			if current != conn {
				continue
			}
			fmt.Println("Error reading:", err)
			b.setConnected(false)
			select {
			case <-ctx.Done():
				return
			case <-changed:
			}
			continue
		}

		if b.onMessage != nil {
			err := b.onMessage(string(message))
			if err != nil {
				fmt.Println("Error handling message:", err)
				b.setConnected(false)
			}
		}
	}
}

func (b *WebSocket) monitorConnection(ctx context.Context) {
	ticker := time.NewTicker(monitorTickInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
		if !b.getConnected() && ctx.Err() == nil {
			fmt.Println("Attempting to reconnect...")
			if err := b.redial(); err != nil {
				fmt.Println("Reconnection failed:")
			}
		}
	}
}

func (b *WebSocket) SetMessageHandler(handler MessageHandler) {
	b.onMessage = handler
}

type WebSocket struct {
	conn         *websocket.Conn
	connChanged  chan struct{}
	url          string
	apiKey       string
	apiSecret    string
	maxAliveTime string
	pingInterval int
	onMessage    MessageHandler
	ctx          context.Context
	cancel       context.CancelFunc
	isConnected  bool
	started      bool
	connMu       sync.Mutex
	writeMu      sync.Mutex
	redialMu     sync.Mutex
	mu           sync.RWMutex
	lifeWG       sync.WaitGroup
}

// setConnected / getConnected guard isConnected, which is accessed concurrently
// from Connect, Disconnect, handleIncomingMessages and monitorConnection.
func (b *WebSocket) setConnected(v bool) {
	b.mu.Lock()
	b.isConnected = v
	b.mu.Unlock()
}

func (b *WebSocket) getConnected() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.isConnected
}

func (b *WebSocket) getConn() *websocket.Conn {
	b.connMu.Lock()
	defer b.connMu.Unlock()
	return b.conn
}

func (b *WebSocket) swapConn(conn *websocket.Conn) {
	b.connMu.Lock()
	old := b.conn
	b.conn = conn
	b.connMu.Unlock()
	if old != nil {
		_ = old.Close()
	}
}

type WebsocketOption func(*WebSocket)

func WithPingInterval(pingInterval int) WebsocketOption {
	return func(c *WebSocket) {
		c.pingInterval = pingInterval
	}
}

func WithMaxAliveTime(maxAliveTime string) WebsocketOption {
	return func(c *WebSocket) {
		c.maxAliveTime = maxAliveTime
	}
}

func NewBybitPrivateWebSocket(url, apiKey, apiSecret string, handler MessageHandler, options ...WebsocketOption) *WebSocket {
	c := &WebSocket{
		url:          url,
		apiKey:       apiKey,
		apiSecret:    apiSecret,
		maxAliveTime: "",
		pingInterval: 20,
		onMessage:    handler,
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.connChanged = make(chan struct{})

	// Apply the provided options
	for _, opt := range options {
		opt(c)
	}

	return c
}

func NewBybitPublicWebSocket(url string, handler MessageHandler) *WebSocket {
	c := &WebSocket{
		url:          url,
		pingInterval: 20, // default is 20 seconds
		onMessage:    handler,
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.connChanged = make(chan struct{})

	return c
}

func (b *WebSocket) Connect() *WebSocket {
	b.redialMu.Lock()
	defer b.redialMu.Unlock()

	b.mu.Lock()
	if b.ctx.Err() != nil {
		b.ctx, b.cancel = context.WithCancel(context.Background())
	}
	b.mu.Unlock()

	if err := b.establish(); err != nil {
		return nil
	}
	b.startLifecycle()
	return b
}

func (b *WebSocket) establish() error {
	wssUrl := b.url
	if b.maxAliveTime != "" {
		wssUrl += "?max_alive_time=" + b.maxAliveTime
	}
	conn, _, err := websocket.DefaultDialer.Dial(wssUrl, nil)
	if err != nil {
		fmt.Printf("connect to %q error: %s\n", wssUrl, err)
		return err
	}
	b.swapConn(conn)
	if b.ctx.Err() != nil {
		b.swapConn(nil)
		return fmt.Errorf("connect to %q aborted: websocket closed", wssUrl)
	}
	if b.requiresAuthentication() {
		if err = b.sendAuth(); err != nil {
			fmt.Println("failed authentication:", fmt.Sprintf("%v", err))
			return err
		}
	}
	b.setConnected(true)
	return nil
}

func (b *WebSocket) redial() error {
	b.redialMu.Lock()
	defer b.redialMu.Unlock()
	return b.establish()
}

func (b *WebSocket) startLifecycle() {
	b.mu.Lock()
	if b.started {
		b.mu.Unlock()
		return
	}
	b.started = true
	ctx := b.ctx
	b.mu.Unlock()

	b.lifeWG.Add(3)
	go func() {
		defer b.lifeWG.Done()
		b.handleIncomingMessages(ctx)
	}()
	go func() {
		defer b.lifeWG.Done()
		b.monitorConnection(ctx)
	}()
	go func() {
		defer b.lifeWG.Done()
		ping(b, ctx)
	}()
}

func (b *WebSocket) SendSubscription(args []string) (*WebSocket, error) {
	reqID := uuid.New().String()
	subMessage := map[string]interface{}{
		"req_id": reqID,
		"op":     "subscribe",
		"args":   args,
	}
	debugWSRequest("subscribe msg: %v", subMessage["args"])
	if err := b.sendAsJson(subMessage); err != nil {
		fmt.Println("Failed to send subscription:", err)
		return b, err
	}
	fmt.Println("Subscription sent successfully.")
	return b, nil
}

// SendRequest sendRequest sends a custom request over the WebSocket connection.
func (b *WebSocket) SendRequest(op string, args map[string]interface{}, headers map[string]string, reqId ...string) (*WebSocket, error) {
	finalReqId := uuid.New().String()
	if len(reqId) > 0 && reqId[0] != "" {
		finalReqId = reqId[0]
	}

	request := map[string]interface{}{
		"reqId":  finalReqId,
		"header": headers,
		"op":     op,
		"args":   []interface{}{args},
	}
	debugWSRequest("request op channel: %v", request["op"])
	debugWSRequest("request msg: %v", request["args"])
	if err := b.sendAsJson(request); err != nil {
		fmt.Println("Failed to send websocket trade request:", err)
		return b, err
	}
	fmt.Println("Successfully sent websocket trade request.")
	return b, nil
}

func (b *WebSocket) SendTradeRequest(tradeTruest map[string]interface{}) (*WebSocket, error) {
	debugWSRequest("trade request op channel: %v", tradeTruest["op"])
	debugWSRequest("trade request msg: %v", tradeTruest["args"])
	if err := b.sendAsJson(tradeTruest); err != nil {
		fmt.Println("Failed to send websocket trade request:", err)
		return b, err
	}
	fmt.Println("Successfully sent websocket trade request.")
	return b, nil
}

func ping(b *WebSocket, ctx context.Context) {
	if b.pingInterval <= 0 {
		fmt.Println("Ping interval is set to a non-positive value.")
		return
	}

	ticker := time.NewTicker(time.Duration(b.pingInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			currentTime := time.Now().Unix()
			pingMessage := map[string]string{
				"op":     "ping",
				"req_id": fmt.Sprintf("%d", currentTime),
			}
			jsonPingMessage, err := json.Marshal(pingMessage)
			if err != nil {
				fmt.Println("Failed to marshal ping message:", err)
				continue
			}
			conn := b.getConn()
			if conn == nil {
				continue
			}
			b.writeMu.Lock()
			if err := conn.WriteMessage(websocket.TextMessage, jsonPingMessage); err != nil {
				b.writeMu.Unlock()
				fmt.Println("Failed to send ping:", err)
				continue
			}
			b.writeMu.Unlock()
			fmt.Println("Ping sent with UTC time:", currentTime)

		case <-ctx.Done():
			fmt.Println("Ping context closed, stopping ping.")
			return
		}
	}
}

func (b *WebSocket) Disconnect() error {
	b.mu.Lock()
	cancel := b.cancel
	b.started = false
	b.mu.Unlock()

	if cancel != nil {
		cancel()
	}
	b.setConnected(false)
	conn := b.getConn()
	var closeErr error
	if conn != nil {
		closeErr = conn.Close()
	}
	finished := make(chan struct{})
	go func() {
		b.lifeWG.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	case <-time.After(2 * time.Second):
	}
	return closeErr
}

func (b *WebSocket) requiresAuthentication() bool {
	return b.url == WEBSOCKET_PRIVATE_MAINNET || b.url == WEBSOCKET_PRIVATE_TESTNET ||
		b.url == WEBSOCKET_TRADE_MAINNET || b.url == WEBSOCKET_TRADE_TESTNET ||
		b.url == WEBSOCKET_TRADE_DEMO || b.url == WEBSOCKET_PRIVATE_DEMO
	// v3 offline
	/*
		b.url == V3_CONTRACT_PRIVATE ||
			b.url == V3_UNIFIED_PRIVATE ||
			b.url == V3_SPOT_PRIVATE
	*/
}

func (b *WebSocket) sendAuth() error {
	// Get current Unix time in milliseconds
	expires := time.Now().UnixNano()/1e6 + 10000
	val := fmt.Sprintf("GET/realtime%d", expires)

	h := hmac.New(sha256.New, []byte(b.apiSecret))
	h.Write([]byte(val))

	// Convert to hexadecimal instead of base64
	signature := hex.EncodeToString(h.Sum(nil))

	authMessage := map[string]interface{}{
		"req_id": uuid.New(),
		"op":     "auth",
		"args":   []interface{}{b.apiKey, expires, signature},
	}
	return b.sendAsJson(authMessage)
}

func (b *WebSocket) sendAsJson(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return b.send(string(data))
}

func (b *WebSocket) send(message string) error {
	conn := b.getConn()
	if conn == nil {
		return fmt.Errorf("websocket: not connected")
	}
	b.writeMu.Lock()
	defer b.writeMu.Unlock()
	return conn.WriteMessage(websocket.TextMessage, []byte(message))
}
