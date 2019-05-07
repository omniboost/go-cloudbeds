package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewUserinfoRequest() UserinfoRequest {
	return UserinfoRequest{
		client:      c,
		queryParams: c.NewUserinfoQueryParams(),
		pathParams:  c.NewUserinfoPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewUserinfoRequestBody(),
	}
}

type UserinfoRequest struct {
	client      *Client
	queryParams *UserinfoQueryParams
	pathParams  *UserinfoPathParams
	method      string
	headers     http.Header
	requestBody UserinfoRequestBody
}

func (c *Client) NewUserinfoQueryParams() *UserinfoQueryParams {
	return &UserinfoQueryParams{}
}

type UserinfoQueryParams struct {
}

func (p UserinfoQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *UserinfoRequest) QueryParams() *UserinfoQueryParams {
	return r.queryParams
}

func (c *Client) NewUserinfoPathParams() *UserinfoPathParams {
	return &UserinfoPathParams{}
}

type UserinfoPathParams struct {
}

func (p *UserinfoPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *UserinfoRequest) PathParams() *UserinfoPathParams {
	return r.pathParams
}

func (r *UserinfoRequest) SetMethod(method string) {
	r.method = method
}

func (r *UserinfoRequest) Method() string {
	return r.method
}

func (s *Client) NewUserinfoRequestBody() UserinfoRequestBody {
	return UserinfoRequestBody{}
}

type UserinfoRequestBody struct {
}

func (r *UserinfoRequest) RequestBody() *UserinfoRequestBody {
	return &r.requestBody
}

func (r *UserinfoRequest) SetRequestBody(body UserinfoRequestBody) {
	r.requestBody = body
}

func (r *UserinfoRequest) NewResponseBody() *UserinfoResponseBody {
	return &UserinfoResponseBody{}
}

type UserinfoResponseBody struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (r *UserinfoRequest) URL() url.URL {
	return r.client.GetEndpointURL("userinfo", r.PathParams())
}

func (r *UserinfoRequest) Do() (UserinfoResponseBody, error) {
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
