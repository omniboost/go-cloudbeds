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

func (c *Client) NewAccountingTrialBalanceConfigurationStatusGetRequest() AccountingTrialBalanceConfigurationStatusGetRequest {
	return AccountingTrialBalanceConfigurationStatusGetRequest{
		client:      c,
		queryParams: c.NewAccountingTrialBalanceConfigurationStatusGetQueryParams(),
		pathParams:  c.NewAccountingTrialBalanceConfigurationStatusGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewAccountingTrialBalanceConfigurationStatusGetRequestBody(),
	}
}

type AccountingTrialBalanceConfigurationStatusGetRequest struct {
	client      *Client
	queryParams *AccountingTrialBalanceConfigurationStatusGetQueryParams
	pathParams  *AccountingTrialBalanceConfigurationStatusGetPathParams
	method      string
	headers     http.Header
	requestBody AccountingTrialBalanceConfigurationStatusGetRequestBody
}

func (c *Client) NewAccountingTrialBalanceConfigurationStatusGetQueryParams() *AccountingTrialBalanceConfigurationStatusGetQueryParams {
	return &AccountingTrialBalanceConfigurationStatusGetQueryParams{}
}

type AccountingTrialBalanceConfigurationStatusGetQueryParams struct{}

func (p AccountingTrialBalanceConfigurationStatusGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountingTrialBalanceConfigurationStatusGetRequest) QueryParams() *AccountingTrialBalanceConfigurationStatusGetQueryParams {
	return r.queryParams
}

func (c *Client) NewAccountingTrialBalanceConfigurationStatusGetPathParams() *AccountingTrialBalanceConfigurationStatusGetPathParams {
	return &AccountingTrialBalanceConfigurationStatusGetPathParams{}
}

type AccountingTrialBalanceConfigurationStatusGetPathParams struct {
}

func (p *AccountingTrialBalanceConfigurationStatusGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountingTrialBalanceConfigurationStatusGetRequest) PathParams() *AccountingTrialBalanceConfigurationStatusGetPathParams {
	return r.pathParams
}

func (r *AccountingTrialBalanceConfigurationStatusGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountingTrialBalanceConfigurationStatusGetRequest) Method() string {
	return r.method
}

func (s *Client) NewAccountingTrialBalanceConfigurationStatusGetRequestBody() AccountingTrialBalanceConfigurationStatusGetRequestBody {
	return AccountingTrialBalanceConfigurationStatusGetRequestBody{}
}

type AccountingTrialBalanceConfigurationStatusGetRequestBody struct{}

func (r *AccountingTrialBalanceConfigurationStatusGetRequest) RequestBody() *AccountingTrialBalanceConfigurationStatusGetRequestBody {
	return &r.requestBody
}

func (r *AccountingTrialBalanceConfigurationStatusGetRequest) SetRequestBody(body AccountingTrialBalanceConfigurationStatusGetRequestBody) {
	r.requestBody = body
}

func (r *AccountingTrialBalanceConfigurationStatusGetRequest) NewResponseBody() *AccountingTrialBalanceConfigurationStatusGetResponseBody {
	return &AccountingTrialBalanceConfigurationStatusGetResponseBody{}
}

type AccountingTrialBalanceConfigurationStatusGetResponseBody struct {
	Configured   bool     `json:"configured"`
	ConfiguredAt DateTime `json:"configuredAt"`
}

func (r *AccountingTrialBalanceConfigurationStatusGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/accounting/v1.0/trial-balance/configuration/status", r.PathParams())
}

func (r *AccountingTrialBalanceConfigurationStatusGetRequest) Do(ctx context.Context) (AccountingTrialBalanceConfigurationStatusGetResponseBody, error) {
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
