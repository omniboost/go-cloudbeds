package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetPaymentsRequest() GetPaymentsRequest {
	return GetPaymentsRequest{
		client:      c,
		queryParams: c.NewGetPaymentsQueryParams(),
		pathParams:  c.NewGetPaymentsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetPaymentsRequestBody(),
	}
}

type GetPaymentsRequest struct {
	client      *Client
	queryParams *GetPaymentsQueryParams
	pathParams  *GetPaymentsPathParams
	method      string
	headers     http.Header
	requestBody GetPaymentsRequestBody
}

func (c *Client) NewGetPaymentsQueryParams() *GetPaymentsQueryParams {
	return &GetPaymentsQueryParams{}
}

type GetPaymentsQueryParams struct {
	// ID for the reservation to be queried.
	ReservationID string `schema:"reservationID"`
	// ID for the house account to be queried.
	HouseAccountID int `schema:"houseAccountID"`
	// ID for the guest to be queried.
	GuestID int `schema:"guestID"`
	// Datetime (lower limit) to be queried. If not sent, and reservationID informed, will use reservation date. In other cases, current date -7 days is used.
	CreatedFrom DateTime `schema:"createdFrom,omitempty"`
	// Datetime (upper limit) to be queried. If not sent, and reservationID informed, will use check-out date. In other cases, current date is used.
	CreatedTo DateTime `schema:"createdTo,omitempty"`
	// Adds payment allocation to response, when available.
	// Standaard waarde: false
	IncludePaymentAllocation bool `schema:"includePaymentAllocation,omitempty"`
	// Page number
	// Standaard waarde: 1
	PageNumber int `schema:"pageNumber,omitempty"`
	// Page size
	// Standaard waarde: 100
	PageSize int `schema:"pageSize,omitempty"`
}

func (p GetPaymentsQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetPaymentsRequest) QueryParams() *GetPaymentsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetPaymentsPathParams() *GetPaymentsPathParams {
	return &GetPaymentsPathParams{}
}

type GetPaymentsPathParams struct {
}

func (p *GetPaymentsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetPaymentsRequest) PathParams() *GetPaymentsPathParams {
	return r.pathParams
}

func (r *GetPaymentsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPaymentsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetPaymentsRequestBody() GetPaymentsRequestBody {
	return GetPaymentsRequestBody{}
}

type GetPaymentsRequestBody struct {
}

func (r *GetPaymentsRequest) RequestBody() *GetPaymentsRequestBody {
	return &r.requestBody
}

func (r *GetPaymentsRequest) SetRequestBody(body GetPaymentsRequestBody) {
	r.requestBody = body
}

func (r *GetPaymentsRequest) NewResponseBody() *GetPaymentsResponseBody {
	return &GetPaymentsResponseBody{}
}

type GetPaymentsResponseBody struct {
	Success bool   `json:"success"`
	Count   int    `json:"count"`
	Total   int    `json:"total"`
	Message string `json:"message"`
	Data    []struct {
	} `json:"data"`
}

func (r *GetPaymentsRequest) URL() url.URL {
	return r.client.GetEndpointURL("getPayments", r.PathParams())
}

func (r *GetPaymentsRequest) Do() (GetPaymentsResponseBody, error) {
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
