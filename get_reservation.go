package cloudbeds

import (
	"context"
	"net/http"
	"net/url"
)

func (c *Client) NewGetReservationRequest() GetReservationRequest {
	return GetReservationRequest{
		client:      c,
		queryParams: c.NewGetReservationQueryParams(),
		pathParams:  c.NewGetReservationPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetReservationRequestBody(),
	}
}

type GetReservationRequest struct {
	client      *Client
	queryParams *GetReservationQueryParams
	pathParams  *GetReservationPathParams
	method      string
	headers     http.Header
	requestBody GetReservationRequestBody
}

func (c *Client) NewGetReservationQueryParams() *GetReservationQueryParams {
	return &GetReservationQueryParams{}
}

type GetReservationQueryParams struct {
	// ID for the properties to be queried (comma-separated).  It can be omitted
	// if the API key is single-property, or to get results from all properties
	// on an association.
	PropertyID int `schema:"propertyID,omitempty"`
	// Reservation Unique Identifier. Obtained from one of the "Reservations" group methods
	ReservationID string `schema:"reservationID,omitempty"`
	// Includes guest requirements data in the response.
	// Defaults to false
	IncludeGuestRequirements bool `schema:"includeGuestRequirements,omitempty"`
}

func (p GetReservationQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetReservationRequest) QueryParams() *GetReservationQueryParams {
	return r.queryParams
}

func (c *Client) NewGetReservationPathParams() *GetReservationPathParams {
	return &GetReservationPathParams{}
}

type GetReservationPathParams struct {
}

func (p *GetReservationPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetReservationRequest) PathParams() *GetReservationPathParams {
	return r.pathParams
}

func (r *GetReservationRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetReservationRequest) Method() string {
	return r.method
}

func (s *Client) NewGetReservationRequestBody() GetReservationRequestBody {
	return GetReservationRequestBody{}
}

type GetReservationRequestBody struct {
}

func (r *GetReservationRequest) RequestBody() *GetReservationRequestBody {
	return &r.requestBody
}

func (r *GetReservationRequest) SetRequestBody(body GetReservationRequestBody) {
	r.requestBody = body
}

func (r *GetReservationRequest) NewResponseBody() *GetReservationResponseBody {
	return &GetReservationResponseBody{}
}

type GetReservationResponseBody struct {
	Success bool        `json:"success"`
	Count   int         `json:"count"`
	Total   int         `json:"total"`
	Message string      `json:"message"`
	Data    Reservation `json:"data"`
}

func (r *GetReservationRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/v1.3/getReservation", r.PathParams())
}

func (r *GetReservationRequest) Do(ctx context.Context) (GetReservationResponseBody, error) {
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
