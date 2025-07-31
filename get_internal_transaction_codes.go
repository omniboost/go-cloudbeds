package cloudbeds

import (
	"net/http"
	"net/url"
)

// Payment - GetInternalTransactionCodes
// Get a list of transactions for a property, or list of properties, for the
// date range specified. If no date range or reservation is specified, it will
// return the transactions for the last 7 days, unless stated otherwise.

func (c *Client) NewGetInternalTransactionCodesRequest() GetInternalTransactionCodesRequest {
	return GetInternalTransactionCodesRequest{
		client:      c,
		queryParams: c.NewGetInternalTransactionCodesQueryParams(),
		pathParams:  c.NewGetInternalTransactionCodesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetInternalTransactionCodesRequestBody(),
	}
}

type GetInternalTransactionCodesRequest struct {
	client      *Client
	queryParams *GetInternalTransactionCodesQueryParams
	pathParams  *GetInternalTransactionCodesPathParams
	method      string
	headers     http.Header
	requestBody GetInternalTransactionCodesRequestBody
}

func (c *Client) NewGetInternalTransactionCodesQueryParams() *GetInternalTransactionCodesQueryParams {
	return &GetInternalTransactionCodesQueryParams{}
}

type GetInternalTransactionCodesQueryParams struct {
}

func (p GetInternalTransactionCodesQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInternalTransactionCodesRequest) QueryParams() *GetInternalTransactionCodesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetInternalTransactionCodesPathParams() *GetInternalTransactionCodesPathParams {
	return &GetInternalTransactionCodesPathParams{}
}

type GetInternalTransactionCodesPathParams struct {
}

func (p *GetInternalTransactionCodesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetInternalTransactionCodesRequest) PathParams() *GetInternalTransactionCodesPathParams {
	return r.pathParams
}

func (r *GetInternalTransactionCodesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInternalTransactionCodesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetInternalTransactionCodesRequestBody() GetInternalTransactionCodesRequestBody {
	return GetInternalTransactionCodesRequestBody{}
}

type GetInternalTransactionCodesRequestBody struct{}

func (r *GetInternalTransactionCodesRequest) RequestBody() *GetInternalTransactionCodesRequestBody {
	return nil
}

func (r *GetInternalTransactionCodesRequest) SetRequestBody(body GetInternalTransactionCodesRequestBody) {
	r.requestBody = body
}

func (r *GetInternalTransactionCodesRequest) NewResponseBody() *GetInternalTransactionCodesResponseBody {
	return &GetInternalTransactionCodesResponseBody{}
}

type GetInternalTransactionCodesResponseBody struct {
	Content InternalTransactionCodes `json:"content"`
}

func (r *GetInternalTransactionCodesRequest) URL() url.URL {
	return r.client.GetEndpointURL("accounting/v1.0/internal-transaction-codes", r.PathParams())
}

func (r *GetInternalTransactionCodesRequest) Do() (GetInternalTransactionCodesResponseBody, error) {
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
