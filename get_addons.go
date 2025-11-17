package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetAddonsRequest() GetAddonsRequest {
	return GetAddonsRequest{
		client:      c,
		queryParams: c.NewGetAddonsQueryParams(),
		pathParams:  c.NewGetAddonsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetAddonsRequestBody(),
	}
}

type GetAddonsRequest struct {
	client      *Client
	queryParams *GetAddonsQueryParams
	pathParams  *GetAddonsPathParams
	method      string
	headers     http.Header
	requestBody GetAddonsRequestBody
}

func (c *Client) NewGetAddonsQueryParams() *GetAddonsQueryParams {
	return &GetAddonsQueryParams{}
}

type GetAddonsQueryParams struct {
	Limit  int `schema:"limit,omitempty"`
	Offset int `schema:"offset,omitempty"`
}

func (p GetAddonsQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAddonsRequest) QueryParams() *GetAddonsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetAddonsPathParams() *GetAddonsPathParams {
	return &GetAddonsPathParams{}
}

type GetAddonsPathParams struct {
}

func (p *GetAddonsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAddonsRequest) PathParams() *GetAddonsPathParams {
	return r.pathParams
}

func (r *GetAddonsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAddonsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetAddonsRequestBody() GetAddonsRequestBody {
	return GetAddonsRequestBody{}
}

type GetAddonsRequestBody struct {
}

func (r *GetAddonsRequest) RequestBody() *GetAddonsRequestBody {
	return &r.requestBody
}

func (r *GetAddonsRequest) SetRequestBody(body GetAddonsRequestBody) {
	r.requestBody = body
}

func (r *GetAddonsRequest) NewResponseBody() *GetAddonsResponseBody {
	return &GetAddonsResponseBody{}
}

type GetAddonsResponseBody struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Data   Addons `json:"data"`
}

func (r *GetAddonsRequest) URL() url.URL {
	return r.client.GetEndpointURL("addons/v1/addons", r.PathParams())
}

func (r *GetAddonsRequest) Do() (GetAddonsResponseBody, error) {
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
