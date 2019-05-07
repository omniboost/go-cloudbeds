package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetHotelDetailsRequest() GetHotelDetailsRequest {
	return GetHotelDetailsRequest{
		client:      c,
		queryParams: c.NewGetHotelDetailsQueryParams(),
		pathParams:  c.NewGetHotelDetailsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetHotelDetailsRequestBody(),
	}
}

type GetHotelDetailsRequest struct {
	client      *Client
	queryParams *GetHotelDetailsQueryParams
	pathParams  *GetHotelDetailsPathParams
	method      string
	headers     http.Header
	requestBody GetHotelDetailsRequestBody
}

func (c *Client) NewGetHotelDetailsQueryParams() *GetHotelDetailsQueryParams {
	return &GetHotelDetailsQueryParams{}
}

type GetHotelDetailsQueryParams struct {
}

func (p GetHotelDetailsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetHotelDetailsRequest) QueryParams() *GetHotelDetailsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetHotelDetailsPathParams() *GetHotelDetailsPathParams {
	return &GetHotelDetailsPathParams{}
}

type GetHotelDetailsPathParams struct {
}

func (p *GetHotelDetailsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetHotelDetailsRequest) PathParams() *GetHotelDetailsPathParams {
	return r.pathParams
}

func (r *GetHotelDetailsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetHotelDetailsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetHotelDetailsRequestBody() GetHotelDetailsRequestBody {
	return GetHotelDetailsRequestBody{}
}

type GetHotelDetailsRequestBody struct {
}

func (r *GetHotelDetailsRequest) RequestBody() *GetHotelDetailsRequestBody {
	return &r.requestBody
}

func (r *GetHotelDetailsRequest) SetRequestBody(body GetHotelDetailsRequestBody) {
	r.requestBody = body
}

func (r *GetHotelDetailsRequest) NewResponseBody() *GetHotelDetailsResponseBody {
	return &GetHotelDetailsResponseBody{}
}

type GetHotelDetailsResponseBody struct {
	Success bool `json:"success"`
	Count   int  `json:"count"`
	Total   int  `json:"total"`
	Data    struct {
		PropertyID    string `json:"propertyID"`
		PropertyName  string `json:"propertyName"`
		PropertyImage []struct {
			Thumb string `json:"thumb"`
			Image string `json:"image"`
		} `json:"propertyImage"`
		PropertyDescription string `json:"propertyDescription"`
		PropertyCurrency    struct {
			CurrencyCode     string `json:"currencyCode"`
			CurrencySymbol   string `json:"currencySymbol"`
			CurrencyPosition string `json:"currencyPosition"`
		} `json:"propertyCurrency"`
		PropertyAdditionalPhotos []struct {
			Thumb string `json:"thumb"`
			Image string `json:"image"`
		} `json:"propertyAdditionalPhotos"`
		PropertyPhone   string `json:"propertyPhone"`
		PropertyEmail   string `json:"propertyEmail"`
		PropertyAddress struct {
			PropertyAddress1  string `json:"propertyAddress1"`
			PropertyAddress2  string `json:"propertyAddress2"`
			PropertyCity      string `json:"propertyCity"`
			PropertyState     string `json:"propertyState"`
			PropertyZip       string `json:"propertyZip"`
			PropertyCountry   string `json:"propertyCountry"`
			PropertyLatitude  string `json:"propertyLatitude"`
			PropertyLongitude string `json:"propertyLongitude"`
		} `json:"propertyAddress"`
		PropertyPolicy struct {
			PropertyCheckInTime         string `json:"propertyCheckInTime"`
			PropertyCheckOutTime        string `json:"propertyCheckOutTime"`
			PropertyLateCheckOutAllowed bool   `json:"propertyLateCheckOutAllowed"`
			PropertyLateCheckOutType    string `json:"propertyLateCheckOutType"`
			PropertyLateCheckOutValue   string `json:"propertyLateCheckOutValue"`
			PropertyTermsAndConditions  string `json:"propertyTermsAndConditions"`
		} `json:"propertyPolicy"`
		PropertyAmenities []string `json:"propertyAmenities"`
	} `json:"data"`
}

func (r *GetHotelDetailsRequest) URL() url.URL {
	return r.client.GetEndpointURL("getHotelDetails", r.PathParams())
}

func (r *GetHotelDetailsRequest) Do() (GetHotelDetailsResponseBody, error) {
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
