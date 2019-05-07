package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetCustomFieldsRequest() GetCustomFieldsRequest {
	return GetCustomFieldsRequest{
		client:      c,
		queryParams: c.NewGetCustomFieldsQueryParams(),
		pathParams:  c.NewGetCustomFieldsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetCustomFieldsRequestBody(),
	}
}

type GetCustomFieldsRequest struct {
	client      *Client
	queryParams *GetCustomFieldsQueryParams
	pathParams  *GetCustomFieldsPathParams
	method      string
	headers     http.Header
	requestBody GetCustomFieldsRequestBody
}

func (c *Client) NewGetCustomFieldsQueryParams() *GetCustomFieldsQueryParams {
	return &GetCustomFieldsQueryParams{}
}

type GetCustomFieldsQueryParams struct {
}

func (p GetCustomFieldsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCustomFieldsRequest) QueryParams() *GetCustomFieldsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCustomFieldsPathParams() *GetCustomFieldsPathParams {
	return &GetCustomFieldsPathParams{}
}

type GetCustomFieldsPathParams struct {
}

func (p *GetCustomFieldsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCustomFieldsRequest) PathParams() *GetCustomFieldsPathParams {
	return r.pathParams
}

func (r *GetCustomFieldsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomFieldsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCustomFieldsRequestBody() GetCustomFieldsRequestBody {
	return GetCustomFieldsRequestBody{}
}

type GetCustomFieldsRequestBody struct {
}

func (r *GetCustomFieldsRequest) RequestBody() *GetCustomFieldsRequestBody {
	return &r.requestBody
}

func (r *GetCustomFieldsRequest) SetRequestBody(body GetCustomFieldsRequestBody) {
	r.requestBody = body
}

func (r *GetCustomFieldsRequest) NewResponseBody() *GetCustomFieldsResponseBody {
	return &GetCustomFieldsResponseBody{}
}

type GetCustomFieldsResponseBody struct {
	Success bool       `json:"success"`
	Count   int        `json:"count"`
	Total   int        `json:"total"`
	Data    []struct{} `json:"data"`
}

func (r *GetCustomFieldsRequest) URL() url.URL {
	return r.client.GetEndpointURL("getCustomFields", r.PathParams())
}

func (r *GetCustomFieldsRequest) Do() (GetCustomFieldsResponseBody, error) {
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
