package bybit_connector

import (
	"context"
	"net/http"
)

// Crypto Loan

func (s *BybitClientRequest) GetBorrowableCoins(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-common/loanable-data",
		secType:  secTypeNone,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetCollateralCoins(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-common/collateral-data",
		secType:  secTypeNone,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetMaxCollateralAmount(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-common/max-collateral-amount",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) AdjustCollateralAmount(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan-common/adjust-ltv",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetCollateralAdjustHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-common/adjustment-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetCryptoLoanPosition(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-common/position",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Flexible loan

func (s *BybitClientRequest) BorrowFlexibleLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan-flexible/borrow",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) RepayFlexibleLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan-flexible/repay",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetFlexibleLoans(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-flexible/ongoing-coin",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetFlexibleBorrowHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-flexible/borrow-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetFlexibleRepaymentHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-flexible/repayment-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Fixed Loan

func (s *BybitClientRequest) GetFixedSupplyingMarket(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-fixed/supply-order-quote",
		secType:  secTypeNone,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetFixedBorrowingMarket(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-fixed/borrow-order-quote",
		secType:  secTypeNone,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetFixedBorrowHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-fixed/borrow-order-info",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetFixedSupplyHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-fixed/supply-order-info",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetFixedBorrowContract(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-fixed/borrow-contract-info",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetFixedSupplyContract(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-fixed/supply-contract-info",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) FixedLoanRepaymentHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-fixed/repayment-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) BorrowFixedLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan-fixed/borrow",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) RepayFixedLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan-fixed/fully-repay",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) CancelBorrowFixedLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan-fixed/borrow-order-cancel",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) CancelSupplyFixedLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan-fixed/supply-order-cancel",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) SupplyFixedLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan-fixed/supply",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// RepayCollateral
func (s *BybitClientRequest) RepayCollateral(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan-flexible/repay-collateral",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// RepayFixedCollateral
func (s *BybitClientRequest) RepayFixedCollateral(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan-fixed/repay-collateral",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// RenewBorrowOrder
func (s *BybitClientRequest) RenewBorrowOrder(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan-fixed/renew",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetRenewOrderInfo
func (s *BybitClientRequest) GetRenewOrderInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-fixed/renew-info",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// GetCryptoLoanMaxCollateralAmount gets the max collateral amount for crypto loan
func (s *BybitClientRequest) GetCryptoLoanMaxCollateralAmount(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan/max-collateral-amount",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}
