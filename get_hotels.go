package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetHotelsRequest() GetHotelsRequest {
	return GetHotelsRequest{
		client:      c,
		queryParams: c.NewGetHotelsQueryParams(),
		pathParams:  c.NewGetHotelsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetHotelsRequestBody(),
	}
}

type GetHotelsRequest struct {
	client      *Client
	queryParams *GetHotelsQueryParams
	pathParams  *GetHotelsPathParams
	method      string
	headers     http.Header
	requestBody GetHotelsRequestBody
}

func (c *Client) NewGetHotelsQueryParams() *GetHotelsQueryParams {
	return &GetHotelsQueryParams{}
}

type GetHotelsQueryParams struct {
}

func (p GetHotelsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetHotelsRequest) QueryParams() *GetHotelsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetHotelsPathParams() *GetHotelsPathParams {
	return &GetHotelsPathParams{}
}

type GetHotelsPathParams struct {
}

func (p *GetHotelsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetHotelsRequest) PathParams() *GetHotelsPathParams {
	return r.pathParams
}

func (r *GetHotelsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetHotelsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetHotelsRequestBody() GetHotelsRequestBody {
	return GetHotelsRequestBody{}
}

type GetHotelsRequestBody struct {
}

func (r *GetHotelsRequest) RequestBody() *GetHotelsRequestBody {
	return &r.requestBody
}

func (r *GetHotelsRequest) SetRequestBody(body GetHotelsRequestBody) {
	r.requestBody = body
}

func (r *GetHotelsRequest) NewResponseBody() *GetHotelsResponseBody {
	return &GetHotelsResponseBody{}
}

type GetHotelsResponseBody struct {
	Success bool `json:"success"`
	Count   int  `json:"count"`
	Total   int  `json:"total"`
	Data    []struct {
		PropertyID          string `json:"propertyID"`
		PropertyName        string `json:"propertyName"`
		PropertyImage       string `json:"propertyImage"`
		PropertyDescription string `json:"propertyDescription"`
		PropertyTimezone    string `json:"propertyTimezone"`
		PropertyCurrency    struct {
			CurrencyCode     string `json:"currencyCode"`
			CurrencySymbol   string `json:"currencySymbol"`
			CurrencyPosition string `json:"currencyPosition"`
		} `json:"propertyCurrency"`
	} `json:"data"`
}

func (r *GetHotelsRequest) URL() url.URL {
	return r.client.GetEndpointURL("getHotels", r.PathParams())
}

func (r *GetHotelsRequest) Do() (GetHotelsResponseBody, error) {
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
