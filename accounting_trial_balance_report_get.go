package cloudbeds

import (
	"context"
	"net/http"
	"net/url"
)

// Payment - getTransactions
// Get a list of transactions for a property, or list of properties, for the
// date range specified. If no date range or reservation is specified, it will
// return the transactions for the last 7 days, unless stated otherwise.

func (c *Client) NewAccountingTrialBalanceReportGetRequest() AccountingTrialBalanceReportGetRequest {
	return AccountingTrialBalanceReportGetRequest{
		client:      c,
		queryParams: c.NewAccountingTrialBalanceReportGetQueryParams(),
		pathParams:  c.NewAccountingTrialBalanceReportGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewAccountingTrialBalanceReportGetRequestBody(),
	}
}

type AccountingTrialBalanceReportGetRequest struct {
	client      *Client
	queryParams *AccountingTrialBalanceReportGetQueryParams
	pathParams  *AccountingTrialBalanceReportGetPathParams
	method      string
	headers     http.Header
	requestBody AccountingTrialBalanceReportGetRequestBody
}

func (c *Client) NewAccountingTrialBalanceReportGetQueryParams() *AccountingTrialBalanceReportGetQueryParams {
	return &AccountingTrialBalanceReportGetQueryParams{}
}

type AccountingTrialBalanceReportGetQueryParams struct {
	Date Date `schema:"date"`
}

func (p AccountingTrialBalanceReportGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountingTrialBalanceReportGetRequest) QueryParams() *AccountingTrialBalanceReportGetQueryParams {
	return r.queryParams
}

func (c *Client) NewAccountingTrialBalanceReportGetPathParams() *AccountingTrialBalanceReportGetPathParams {
	return &AccountingTrialBalanceReportGetPathParams{}
}

type AccountingTrialBalanceReportGetPathParams struct {
}

func (p *AccountingTrialBalanceReportGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountingTrialBalanceReportGetRequest) PathParams() *AccountingTrialBalanceReportGetPathParams {
	return r.pathParams
}

func (r *AccountingTrialBalanceReportGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountingTrialBalanceReportGetRequest) Method() string {
	return r.method
}

func (s *Client) NewAccountingTrialBalanceReportGetRequestBody() AccountingTrialBalanceReportGetRequestBody {
	return AccountingTrialBalanceReportGetRequestBody{}
}

type AccountingTrialBalanceReportGetRequestBody struct {}

func (r *AccountingTrialBalanceReportGetRequest) RequestBody() *AccountingTrialBalanceReportGetRequestBody {
	return &r.requestBody
}

func (r *AccountingTrialBalanceReportGetRequest) SetRequestBody(body AccountingTrialBalanceReportGetRequestBody) {
	r.requestBody = body
}

func (r *AccountingTrialBalanceReportGetRequest) NewResponseBody() *AccountingTrialBalanceReportGetResponseBody {
	return &AccountingTrialBalanceReportGetResponseBody{}
}

type AccountingTrialBalanceReportGetResponseBody struct {
}

func (r *AccountingTrialBalanceReportGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/accounting/v1.0/trial-balance/report", r.PathParams())
}

func (r *AccountingTrialBalanceReportGetRequest) Do(ctx context.Context) (AccountingTrialBalanceReportGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r.Method(), r.URL(), nil)
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
