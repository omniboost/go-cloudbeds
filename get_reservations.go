package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetReservationsRequest() GetReservationsRequest {
	return GetReservationsRequest{
		client:      c,
		queryParams: c.NewGetReservationsQueryParams(),
		pathParams:  c.NewGetReservationsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetReservationsRequestBody(),
	}
}

type GetReservationsRequest struct {
	client      *Client
	queryParams *GetReservationsQueryParams
	pathParams  *GetReservationsPathParams
	method      string
	headers     http.Header
	requestBody GetReservationsRequestBody
}

func (c *Client) NewGetReservationsQueryParams() *GetReservationsQueryParams {
	return &GetReservationsQueryParams{}
}

type GetReservationsQueryParams struct {
	// ID for the properties to be queried (comma-separated).  It can be omitted
	// if the API key is single-property, or to get results from all properties
	// on an association.
	PropertyID string `schema:"propertyID,omitempty"`
	// Filter by current reservation status
	// Toegestane waarden: "not_confirmed", "confirmed", "canceled",
	// "checked_in", "checked_out", "no_show"
	Status Status `schema:"status,omitempty"`
	// 	Inferior limit datetime, used to filter reservations, based on booking
	// 	date
	ResultsFrom DateTime `schema:"resultsFrom,omitempty"`
	// Superior limit datetime, used to filter reservations, based on booking date
	ResultsTo DateTime `schema:"resultsTo,omitempty"`
	// Inferior limit datetime, used to filter reservations, based on booking modification date
	ModifiedFrom DateTime `schema:"modifiedFrom,omitempty"`
	// Superior limit datetime, used to filter reservations, based on booking modification date
	ModifiedTo DateTime `schema:"modifiedTo,omitempty"`
	// Filters reservations result to return only reservations with check-in
	// date range starting on this date
	CheckinFrom DateTime `schema:"checkinFrom,omitempty"`
	// Filters reservations result to return only reservations with check-in
	// date range ending on this date
	CheckinTo DateTime `schema:"checkinTo,omitempty"`
	// Filters reservations result to return only reservations with check-out
	// date range starting on this date
	CheckoutFrom DateTime `schema:"checkoutFrom,omitempty"`
	// Filters reservations result to return only reservations with check-out
	// date range ending on this date
	CheckoutTo DateTime `schema:"checkoutTo,omitempty"`
	// Filters reservation with the supplied room ID. CheckIn/checkOut dates OR
	// status are required. If dates are provided and span more than one day,
	// more than one reservation can be returned. If roomID supplied, roomName
	// is ignored.
	RoomID string `schema:"roomID,omitempty"`
	// Filters reservation with the supplied room name (customizable by each
	// property). CheckIn/checkOut dates OR status are required. If dates are
	// provided and span more than one day, more than one reservation can be
	// returned.
	RoomName string `schema:"roomName,omitempty"`
	// If guests details should be included or not
	// Standaard waarde: false
	IncludeGuestsDetails Bool `schema:"includeGuestsDetails,omitempty"`
	// Sort response results by most recent action
	SortByRecent bool `schema:"sortByRecent,omitempty"`
	// Results page number
	// Standaard waarde: 1
	PageNumber int `schema:"pageNumber,omitempty"`
	// Results page size. Max = 100
	// Standaard waarde: 100
	PageSize int `schema:"pageSize,omitempty"`
}

func (p GetReservationsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetReservationsRequest) QueryParams() *GetReservationsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetReservationsPathParams() *GetReservationsPathParams {
	return &GetReservationsPathParams{}
}

type GetReservationsPathParams struct {
}

func (p *GetReservationsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetReservationsRequest) PathParams() *GetReservationsPathParams {
	return r.pathParams
}

func (r *GetReservationsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetReservationsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetReservationsRequestBody() GetReservationsRequestBody {
	return GetReservationsRequestBody{}
}

type GetReservationsRequestBody struct {
}

func (r *GetReservationsRequest) RequestBody() *GetReservationsRequestBody {
	return &r.requestBody
}

func (r *GetReservationsRequest) SetRequestBody(body GetReservationsRequestBody) {
	r.requestBody = body
}

func (r *GetReservationsRequest) NewResponseBody() *GetReservationsResponseBody {
	return &GetReservationsResponseBody{}
}

type GetReservationsResponseBody struct {
	Success bool `json:"success"`
	Count   int  `json:"count"`
	Total   int  `json:"total"`
	Data    []struct {
		PropertyID           string   `json:"propertyID"`    // Properties identifier
		ReservationID        string   `json:"reservationID"` // Reservation's unique identifier
		DateCreated          DateTime `json:"dateCreated"`
		DateModified         DateTime `json:"dateModified"`
		Status               Status   `json:"status"`
		GuestID              Int      `json:"guestID"`
		GuestName            string   `json:"guestName"`
		StartDate            Date     `json:"startDate"`
		EndDate              Date     `json:"endDate"`
		Adults               Int      `json:"adults"`
		Children             Int      `json:"children"`
		Balance              float64  `json:"balance"`
		SourceName           string   `json:"sourceName"` // Source of reservation
		ThirdPartyIdentifier string   `json:"thirdPartyIdentifier"`
		GuestList            map[string]struct {
			GuestID string `json:"guestID"` // ID of the guest

			GuestName                   string `json:"guestName"`
			GuestFirstName              string `json:"guestFirstName"`
			GuestLastName               string `json:"guestLastName"`
			Guestgender                 string `json:"guestGender"` // Toegestane waarden: "M", "F", "N/A"
			GuestEmail                  string `json:"guestEmail"`
			GuestPhone                  string `json:"guestPhone"`
			GuestCellPhone              string `json:"guestCellPhone"`
			GuestAddress                string `json:"guestAddress"`
			GuestAddress2               string `json:"guestAddress2"`
			GuestCity                   string `json:"guestCity"`
			GuestState                  string `json:"guestState"`
			GuestCountry                string `json:"guestCountry"`
			GuestZip                    string `json:"guestZip"`
			GuestBirthdate              Date   `json:"guestBirthdate"`
			GuestDocumentType           string `json:"guestDocumentType"`
			GuestDocumentNumber         string `json:"guestDocumentNumber"`
			GuestDocumentIssueDate      Date   `json:"guestDocumentIssueDate"`
			GuestDocumentIsseingCountry string `json:"guestDocumentIssuingCountry"`
			GuestDocumentExpirationDate Date   `json:"guestDocumentExpirationDate"`
			TaxID                       string `json:"taxID"`        //  Guest's tax ID
			CompanyTaxID                string `json:"companyTaxID"` // Guest's company tax ID
			CompanyName                 string `json:"companyName"`  // Guest's company name
			SubReservationID            string `json:"subReservationID"`
			StartDate                   Date   `json:"startDate"`
			EndDate                     Date   `json:"endDate"`
			AssignedRoom                bool   `json:"assignedRoom"` // Returns true if guest has roomed assigned, false if not
			RoomID                      string `json:"roomID"`       // Room ID where guest is assigned
			RoomName                    string `json:"roomName"`     // Room Name where guest is assigned
			RoomTypeName                string `json:"roomTypeName"` // Room Name where guest is assigned
			IsMainGuest                 bool   `json:"isMainGuest"`
			Rooms                       []struct {
				RoomID       string `json:"roomID"`       // Room ID where guest is assigned
				RoomName     string `json:"roomName"`     // Room Name where guest is assigned
				RoomTypeName string `json:"roomTypeName"` // Room Type Name where guest is assigned
			} // List of all rooms that guest is assigned to
			CustomFields []struct {
				CustomFieldName  string `json:"customFieldName"`  // Custom Field Name
				CustomFieldValue string `json:"customFieldValue"` // Custom Field Value
			} // `json:"customFields
			IsAnonymized bool `json:"isAnonymized"` //  Flag indicating the guest data was removed upon request
		} `json:"guestList"`
	} `json:"data"`
}

func (r *GetReservationsRequest) URL() url.URL {
	return r.client.GetEndpointURL("getReservations", r.PathParams())
}

func (r *GetReservationsRequest) Do() (GetReservationsResponseBody, error) {
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
