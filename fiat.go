package bybit_connector

import (
	"context"
	"net/http"
)

// GetFiatCoinList queries the list of supported fiat coins
func (s *BybitClientRequest) GetFiatCoinList(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/fiat/query-coin-list",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetFiatReferencePrice gets the fiat reference price
func (s *BybitClientRequest) GetFiatReferencePrice(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/fiat/reference-price",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// ApplyFiatQuote applies for a fiat quote
func (s *BybitClientRequest) ApplyFiatQuote(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/fiat/quote-apply",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// ExecuteFiatTrade executes a fiat trade
func (s *BybitClientRequest) ExecuteFiatTrade(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/fiat/trade-execute",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// QueryFiatTrade queries a fiat trade
func (s *BybitClientRequest) QueryFiatTrade(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/fiat/trade-query",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetFiatTradeHistory gets the fiat trade history
func (s *BybitClientRequest) GetFiatTradeHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/fiat/query-trade-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetFiatBalance queries the fiat balance
func (s *BybitClientRequest) GetFiatBalance(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/fiat/balance-query",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}
