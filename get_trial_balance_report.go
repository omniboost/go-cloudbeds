package cloudbeds

import (
	"context"
	"net/http"
	"net/url"
)

func (c *Client) NewGetTrialBalanceReportRequest() GetTrialBalanceReportRequest {
	return GetTrialBalanceReportRequest{
		client:      c,
		queryParams: c.NewGetTrialBalanceReportQueryParams(),
		pathParams:  c.NewGetTrialBalanceReportPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetTrialBalanceReportRequestBody(),
	}
}

type GetTrialBalanceReportRequest struct {
	client      *Client
	queryParams *GetTrialBalanceReportQueryParams
	pathParams  *GetTrialBalanceReportPathParams
	method      string
	headers     http.Header
	requestBody GetTrialBalanceReportRequestBody
}

func (c *Client) NewGetTrialBalanceReportQueryParams() *GetTrialBalanceReportQueryParams {
	return &GetTrialBalanceReportQueryParams{}
}

type GetTrialBalanceReportQueryParams struct {
	Date Date `schema:"date"`
}

func (p GetTrialBalanceReportQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTrialBalanceReportRequest) QueryParams() *GetTrialBalanceReportQueryParams {
	return r.queryParams
}

func (c *Client) NewGetTrialBalanceReportPathParams() *GetTrialBalanceReportPathParams {
	return &GetTrialBalanceReportPathParams{}
}

type GetTrialBalanceReportPathParams struct {
}

func (p *GetTrialBalanceReportPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTrialBalanceReportRequest) PathParams() *GetTrialBalanceReportPathParams {
	return r.pathParams
}

func (r *GetTrialBalanceReportRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTrialBalanceReportRequest) Method() string {
	return r.method
}

func (s *Client) NewGetTrialBalanceReportRequestBody() GetTrialBalanceReportRequestBody {
	return GetTrialBalanceReportRequestBody{}
}

type GetTrialBalanceReportRequestBody struct {
}

func (r *GetTrialBalanceReportRequest) RequestBody() *GetTrialBalanceReportRequestBody {
	return &r.requestBody
}

func (r *GetTrialBalanceReportRequest) SetRequestBody(body GetTrialBalanceReportRequestBody) {
	r.requestBody = body
}

func (r *GetTrialBalanceReportRequest) NewResponseBody() *GetTrialBalanceReportResponseBody {
	return &GetTrialBalanceReportResponseBody{}
}

type GetTrialBalanceReportResponseBody struct {
	TrialBalanceId string `json:"trialBalanceId"`
	Summary        struct {
		HotelOpeningBalance          float64 `json:"hotelOpeningBalance"`
		DepositActivity              float64 `json:"depositActivity"`
		GuestLedgerActivity          float64 `json:"guestLedgerActivity"`
		ArActivity                   float64 `json:"arActivity"`
		TotalActivity                float64 `json:"totalActivity"`
		HotelClosingBalance          float64 `json:"hotelClosingBalance"`
		GuestLedgerOpeningBalance    float64 `json:"guestLedgerOpeningBalance"`
		GuestLedgerTransactionsTotal float64 `json:"guestLedgerTransactionsTotal"`
		GuestLedgerDepositTransfers  float64 `json:"guestLedgerDepositTransfers"`
		GuestLedgerArTransfers       float64 `json:"guestLedgerArTransfers"`
		GuestLedgerClosingBalance    float64 `json:"guestLedgerClosingBalance"`
		OpeningBalance               float64 `json:"openingBalance"`
		TransactionsTotalAmount      float64 `json:"transactionsTotalAmount"`
		LedgerActivity               float64 `json:"ledgerActivity"`
		ArPayments                   float64 `json:"arPayments"`
		DepositTransfers             float64 `json:"depositTransfers"`
		ClosingBalance               float64 `json:"closingBalance"`
		ArTransfers                  float64 `json:"arTransfers"`
	} `json:"summary"`
	LedgerBalances struct {
		DepositLedger []struct {
			Code        *string `json:"code"`
			Description string  `json:"description"`
			Amount      float64 `json:"amount"`
		} `json:"depositLedger"`
		GuestLedger []struct {
			Code        *string `json:"code"`
			Description string  `json:"description"`
			Amount      float64 `json:"amount"`
		} `json:"guestLedger"`
		AccountsReceivable []struct {
			Code        *string `json:"code"`
			Description string  `json:"description"`
			Amount      float64 `json:"amount"`
		} `json:"accountsReceivable"`
	} `json:"ledgerBalances"`
	GuestLedger struct {
		Charges []struct {
			Code        string  `json:"code"`
			Description string  `json:"description"`
			Amount      float64 `json:"amount"`
		} `json:"charges"`
		Taxes []struct {
			Code        string  `json:"code"`
			Description string  `json:"description"`
			Amount      float64 `json:"amount"`
		} `json:"taxes"`
		Payments []interface{} `json:"payments"`
	} `json:"guestLedger"`
}

func (r *GetTrialBalanceReportRequest) URL() url.URL {
	return r.client.GetEndpointURL("accounting/v1.0/trial-balance/report", r.PathParams())
}

func (r *GetTrialBalanceReportRequest) Do(ctx context.Context) (GetTrialBalanceReportResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
