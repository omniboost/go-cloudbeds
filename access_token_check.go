package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewAccessTokenCheckRequest() AccessTokenCheckRequest {
	return AccessTokenCheckRequest{
		client:      c,
		queryParams: c.NewAccessTokenCheckQueryParams(),
		pathParams:  c.NewAccessTokenCheckPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewAccessTokenCheckRequestBody(),
	}
}

type AccessTokenCheckRequest struct {
	client      *Client
	queryParams *AccessTokenCheckQueryParams
	pathParams  *AccessTokenCheckPathParams
	method      string
	headers     http.Header
	requestBody AccessTokenCheckRequestBody
}

func (c *Client) NewAccessTokenCheckQueryParams() *AccessTokenCheckQueryParams {
	return &AccessTokenCheckQueryParams{}
}

type AccessTokenCheckQueryParams struct {
}

func (p AccessTokenCheckQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccessTokenCheckRequest) QueryParams() *AccessTokenCheckQueryParams {
	return r.queryParams
}

func (c *Client) NewAccessTokenCheckPathParams() *AccessTokenCheckPathParams {
	return &AccessTokenCheckPathParams{}
}

type AccessTokenCheckPathParams struct {
}

func (p *AccessTokenCheckPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccessTokenCheckRequest) PathParams() *AccessTokenCheckPathParams {
	return r.pathParams
}

func (r *AccessTokenCheckRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccessTokenCheckRequest) Method() string {
	return r.method
}

func (s *Client) NewAccessTokenCheckRequestBody() AccessTokenCheckRequestBody {
	return AccessTokenCheckRequestBody{}
}

type AccessTokenCheckRequestBody struct {
}

func (r *AccessTokenCheckRequest) RequestBody() *AccessTokenCheckRequestBody {
	return &r.requestBody
}

func (r *AccessTokenCheckRequest) SetRequestBody(body AccessTokenCheckRequestBody) {
	r.requestBody = body
}

func (r *AccessTokenCheckRequest) NewResponseBody() *AccessTokenCheckResponseBody {
	return &AccessTokenCheckResponseBody{}
}

type AccessTokenCheckResponseBody struct {
	Success bool `json:"success"`
}

func (r *AccessTokenCheckRequest) URL() url.URL {
	return r.client.GetEndpointURL("access_token_check", r.PathParams())
}

func (r *AccessTokenCheckRequest) Do() (AccessTokenCheckResponseBody, error) {
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
