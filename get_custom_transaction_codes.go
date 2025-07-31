package cloudbeds

import (
	"net/http"
	"net/url"
)

// Payment - GetCustomTransactionCodes
// Get a list of transactions for a property, or list of properties, for the
// date range specified. If no date range or reservation is specified, it will
// return the transactions for the last 7 days, unless stated otherwise.

func (c *Client) NewGetCustomTransactionCodesRequest() GetCustomTransactionCodesRequest {
	return GetCustomTransactionCodesRequest{
		client:      c,
		queryParams: c.NewGetCustomTransactionCodesQueryParams(),
		pathParams:  c.NewGetCustomTransactionCodesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetCustomTransactionCodesRequestBody(),
	}
}

type GetCustomTransactionCodesRequest struct {
	client      *Client
	queryParams *GetCustomTransactionCodesQueryParams
	pathParams  *GetCustomTransactionCodesPathParams
	method      string
	headers     http.Header
	requestBody GetCustomTransactionCodesRequestBody
}

func (c *Client) NewGetCustomTransactionCodesQueryParams() *GetCustomTransactionCodesQueryParams {
	return &GetCustomTransactionCodesQueryParams{}
}

type GetCustomTransactionCodesQueryParams struct {
}

func (p GetCustomTransactionCodesQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCustomTransactionCodesRequest) QueryParams() *GetCustomTransactionCodesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCustomTransactionCodesPathParams() *GetCustomTransactionCodesPathParams {
	return &GetCustomTransactionCodesPathParams{}
}

type GetCustomTransactionCodesPathParams struct {
}

func (p *GetCustomTransactionCodesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCustomTransactionCodesRequest) PathParams() *GetCustomTransactionCodesPathParams {
	return r.pathParams
}

func (r *GetCustomTransactionCodesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomTransactionCodesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCustomTransactionCodesRequestBody() GetCustomTransactionCodesRequestBody {
	return GetCustomTransactionCodesRequestBody{}
}

type GetCustomTransactionCodesRequestBody struct{}

func (r *GetCustomTransactionCodesRequest) RequestBody() *GetCustomTransactionCodesRequestBody {
	return nil
}

func (r *GetCustomTransactionCodesRequest) SetRequestBody(body GetCustomTransactionCodesRequestBody) {
	r.requestBody = body
}

func (r *GetCustomTransactionCodesRequest) NewResponseBody() *GetCustomTransactionCodesResponseBody {
	return &GetCustomTransactionCodesResponseBody{}
}

type GetCustomTransactionCodesResponseBody CustomTransactionCodes

func (r *GetCustomTransactionCodesRequest) URL() url.URL {
	return r.client.GetEndpointURL("accounting/v1.0/custom-transaction-codes", r.PathParams())
}

func (r *GetCustomTransactionCodesRequest) Do() (GetCustomTransactionCodesResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
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
