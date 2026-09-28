package bybit_connector

import (
	"context"
	"net/http"

	"github.com/Igor-Gagarin/bybit.go.api/handlers"
)

func (s *BybitClientRequest) GetTransactionLog(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	var endpoint string
	if s.isUta {
		endpoint = "/v5/account/transaction-log"
	} else {
		endpoint = "/v5/account/contract-transaction-log"
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetFeeRates(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/fee-rate",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetAccountWallet(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/wallet-balance",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetBorrowHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/borrow-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetCoinGreeks(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/asset/coin-greeks",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetCollateralInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/collateral-info",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetAccountInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/info",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetMMPState(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/mmp-state",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetTransferableAmount(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/withdrawal",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) SetSpotHedgeMode(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/set-hedging-mode",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) UpgradeToUTA(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/account/upgrade-to-uta",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) SetCollateralCoin(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/account/set-collateral-switch",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) BatchSetCollateralCoin(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/account/set-collateral-switch-batch",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) SetMarginMode(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/account/set-margin-mode",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) SetMarketMakerProtection(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/account/mmp-modify",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) ResetMarketMakerProtection(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/account/mmp-reset",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) RepayLiability(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/account/quick-repayment",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetDisconnectProtectionInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/query-dcp-info",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetSelfMarketProtectionGroup(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/smp-group",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetAccountInstrumentsInfo
func (s *BybitClientRequest) GetAccountInstrumentsInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/instruments-info",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// ManualBorrow
func (s *BybitClientRequest) ManualBorrow(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/account/borrow",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetMaxBorrowableAmount
func (s *BybitClientRequest) GetMaxBorrowableAmount(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-margin-trade/max-borrowable",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetPositionTiers
func (s *BybitClientRequest) GetPositionTiers(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-margin-trade/position-tiers",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetCoinState
func (s *BybitClientRequest) GetCoinState(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-margin-trade/coinstate",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetAvailableAmountToRepay
func (s *BybitClientRequest) GetAvailableAmountToRepay(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-margin-trade/repayment-available-amount",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// ManualRepay
func (s *BybitClientRequest) ManualRepay(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/account/repay",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// ManualRepayWithoutAssetConversion
func (s *BybitClientRequest) ManualRepayWithoutAssetConversion(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/account/no-convert-repay",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// SetLimitPriceAction sets the limit order price action when it exceeds the price limit
func (s *BybitClientRequest) SetLimitPriceAction(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/account/set-limit-px-action",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}
