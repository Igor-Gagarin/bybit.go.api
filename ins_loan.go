package bybit_connector

import (
	"context"
	"github.com/Igor-Gagarin/bybit.go.api/handlers"
	"net/http"
)

func (s *BybitClientRequest) GetInsLoanInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/ins-loan/product-infos",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetInsMarginCoinInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/ins-loan/ensure-tokens-convert",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetInsLoanOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/ins-loan/loan-order",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetInsRepayOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/ins-loan/repaid-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) GetInsLoanToValue(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/ins-loan/ltv-convert",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

func (s *BybitClientRequest) AssociateInsLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/ins-loan/association-uid",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// RepayInsLoan repays the institutional loan
func (s *BybitClientRequest) RepayInsLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/ins-loan/repay-loan",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Crypto Loan

// Deprecated: BorrowCryptoLoan is deprecated.
func (s *BybitClientRequest) BorrowCryptoLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan/borrow",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: RepayCryptoLoan is deprecated.
func (s *BybitClientRequest) RepayCryptoLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan/repay",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: AdjustCryptoLoanToValue is deprecated.
func (s *BybitClientRequest) AdjustCryptoLoanToValue(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/crypto-loan/adjust-ltv",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: GetCryptoLoanCollateralInfo is deprecated.
func (s *BybitClientRequest) GetCryptoLoanCollateralInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan/collateral-data",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: GetCryptoLoanBorrowInfo is deprecated.
func (s *BybitClientRequest) GetCryptoLoanBorrowInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan/loanable-data",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: GetCryptoLoanBorrowLimit is deprecated.
func (s *BybitClientRequest) GetCryptoLoanBorrowLimit(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan/borrowable-collateralisable-number",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: GetCryptoLoanUnpaidLoans is deprecated.
func (s *BybitClientRequest) GetCryptoLoanUnpaidLoans(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan/ongoing-orders",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: GetCryptoLoanRepaymentHistory is deprecated.
func (s *BybitClientRequest) GetCryptoLoanRepaymentHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan/repayment-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: GetCryptoLoanBorrowHistory is deprecated.
func (s *BybitClientRequest) GetCryptoLoanBorrowHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan/borrow-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: GetCryptoLoanMaxCollateral is deprecated.
func (s *BybitClientRequest) GetCryptoLoanMaxCollateral(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan/borrow-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: GetCryptoLoanAdjustHistory is deprecated.
func (s *BybitClientRequest) GetCryptoLoanAdjustHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan/adjustment-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: GetCryptoLoanCompletedHistory is deprecated.
func (s *BybitClientRequest) GetCryptoLoanCompletedHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan/borrow-history",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// C2C Loan End

// Deprecated: GetC2cLendingCoinInfo is deprecated.
func (s *BybitClientRequest) GetC2cLendingCoinInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/lending/info",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: GetC2cLendingOrders is deprecated.
func (s *BybitClientRequest) GetC2cLendingOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/lending/history-order",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: GetC2cLendingAccountInfo is deprecated.
func (s *BybitClientRequest) GetC2cLendingAccountInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/lending/account",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: C2cDepositFunds is deprecated.
func (s *BybitClientRequest) C2cDepositFunds(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/lending/purchase",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: C2cRedeemFunds is deprecated.
func (s *BybitClientRequest) C2cRedeemFunds(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/lending/redeem",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}

// Deprecated: C2cCancelRedeemFunds is deprecated.
func (s *BybitClientRequest) C2cCancelRedeemFunds(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/lending/redeem-cancel",
		secType:  secTypeSigned,
	}
	data, headers, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data, headers)
}
