package cloudbeds

import (
	"context"
	"net/http"
	"net/url"
)

// GL - GetCustomGeneralLedgerCodes
// Retrieve all custom general ledger (GL) codes configured for a property.

func (c *Client) NewGetCustomGeneralLedgerCodesRequest() GetCustomGeneralLedgerCodesRequest {
	return GetCustomGeneralLedgerCodesRequest{
		client:      c,
		queryParams: c.NewGetCustomGeneralLedgerCodesQueryParams(),
		pathParams:  c.NewGetCustomGeneralLedgerCodesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetCustomGeneralLedgerCodesRequestBody(),
	}
}

type GetCustomGeneralLedgerCodesRequest struct {
	client      *Client
	queryParams *GetCustomGeneralLedgerCodesQueryParams
	pathParams  *GetCustomGeneralLedgerCodesPathParams
	method      string
	headers     http.Header
	requestBody GetCustomGeneralLedgerCodesRequestBody
}

func (c *Client) NewGetCustomGeneralLedgerCodesQueryParams() *GetCustomGeneralLedgerCodesQueryParams {
	return &GetCustomGeneralLedgerCodesQueryParams{}
}

type GetCustomGeneralLedgerCodesQueryParams struct {
}

func (p GetCustomGeneralLedgerCodesQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCustomGeneralLedgerCodesRequest) QueryParams() *GetCustomGeneralLedgerCodesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCustomGeneralLedgerCodesPathParams() *GetCustomGeneralLedgerCodesPathParams {
	return &GetCustomGeneralLedgerCodesPathParams{}
}

type GetCustomGeneralLedgerCodesPathParams struct {
}

func (p *GetCustomGeneralLedgerCodesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCustomGeneralLedgerCodesRequest) PathParams() *GetCustomGeneralLedgerCodesPathParams {
	return r.pathParams
}

func (r *GetCustomGeneralLedgerCodesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomGeneralLedgerCodesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCustomGeneralLedgerCodesRequestBody() GetCustomGeneralLedgerCodesRequestBody {
	return GetCustomGeneralLedgerCodesRequestBody{}
}

type GetCustomGeneralLedgerCodesRequestBody struct{}

func (r *GetCustomGeneralLedgerCodesRequest) RequestBody() *GetCustomGeneralLedgerCodesRequestBody {
	return nil
}

func (r *GetCustomGeneralLedgerCodesRequest) SetRequestBody(body GetCustomGeneralLedgerCodesRequestBody) {
	r.requestBody = body
}

func (r *GetCustomGeneralLedgerCodesRequest) NewResponseBody() *GetCustomGeneralLedgerCodesResponseBody {
	return &GetCustomGeneralLedgerCodesResponseBody{}
}

type GetCustomGeneralLedgerCodesResponseBody CustomGeneralLedgerCodes

func (r *GetCustomGeneralLedgerCodesRequest) URL() url.URL {
	return r.client.GetEndpointURL("accounting/v1.0/custom-general-ledger-codes", r.PathParams())
}

func (r *GetCustomGeneralLedgerCodesRequest) Do(ctx context.Context) (GetCustomGeneralLedgerCodesResponseBody, error) {
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
