package bybit_connector

import (
	"context"
	"net/http"

	"github.com/Igor-Gagarin/bybit.go.api/handlers"
)

func (s *BybitClientRequest) GetRFQConfig(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/rfq/config",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) CreateRFQ(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/rfq/create-rfq",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// CancelRFQ
// Added in 2025-10-10: Cancel an RFQ inquiry
func (s *BybitClientRequest) CancelRFQ(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/rfq/cancel-rfq",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// CancelAllRFQ
// Added in 2025-10-10: Cancel an RFQ inquiry
func (s *BybitClientRequest) CancelAllRFQ(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/rfq/cancel-all-rfq",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetRFQList
// Added in 2025-10-10: Query RFQ list
func (s *BybitClientRequest) GetRFQList(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/rfq/list",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetRFQRealtimePrice
// Added in 2025-10-10: Get RFQ realtime price
func (s *BybitClientRequest) GetRFQRealtimePrice(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/rfq/realtime",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetRFQQuoteRealtime
// Added in 2025-10-10: Get RFQ quote realtime information
func (s *BybitClientRequest) GetRFQQuoteRealtime(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/rfq/quote-realtime",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// CreateRFQQuote
// Added in 2025-10-10: Create/Apply RFQ quote (quoter operation)
func (s *BybitClientRequest) CreateRFQQuote(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/rfq/quote-apply",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// ExecuteRFQQuote
// Added in 2025-10-10: Execute RFQ quote (inquirer operation)
func (s *BybitClientRequest) ExecuteRFQQuote(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/rfq/quote-execute",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// CancelRFQQuote
// Added in 2025-10-10: Cancel RFQ quote (quoter operation)
func (s *BybitClientRequest) CancelRFQQuote(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/rfq/quote-cancel",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetRFQQuoteList
// Added in 2025-10-10: Query RFQ quote list
func (s *BybitClientRequest) GetRFQQuoteList(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/rfq/quote-list",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetRFQHistory
// Added in 2025-10-10: Get RFQ history
func (s *BybitClientRequest) GetRFQHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/rfq/history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetRFQPublicTrades
// Added in 2025-10-10: Get RFQ public trades
func (s *BybitClientRequest) GetRFQPublicTrades(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/rfq/public-trades",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetRFQTradeList
// Added in 2025-10-10: Get RFQ trade list
func (s *BybitClientRequest) GetRFQTradeList(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/rfq/trade-list",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}
