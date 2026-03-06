package cloudbeds

import (
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
	Success bool   `json:"success"`
	Count   int    `json:"count"`
	Total   int    `json:"total"`
	Message string `json:"message"`
	Data    struct {
		PropertyID           string          `json:"propertyID"`    // Properties identifier
		ReservationID        string          `json:"reservationID"` // Reservation's unique identifier
		DateCreated          DateTime        `json:"dateCreated"`
		DateModified         DateTime        `json:"dateModified"`
		EstimatedArrivalTime HourMinute      `json:"estimatedArrivalTime"`
		Source               string          `json:"source"`
		SourceID             string          `json:"sourceID"`
		Status               Status          `json:"status"`
		TotalRevenue         float64         `json:"totalRevenue"`
		GuestID              Int             `json:"guestID"`
		GuestName            string          `json:"guestName"`
		GuestEmail           string          `json:"guestEmail"`
		StartDate            Date            `json:"startDate"`
		EndDate              Date            `json:"endDate"`
		OrderID              string          `json:"orderID"`
		Origin               string          `json:"origin"`
		Cancellation         string          `json:"cancellation"`
		AllotmentBlockCode   string          `json:"allotmentBlockCode"`
		Adults               Int             `json:"adults"`
		Children             Int             `json:"children"`
		Total                float64         `json:"total"`
		Balance              float64         `json:"balance"`
		BalanceDetailed      BalanceDetailed `json:"balanceDetailed"`
		Assigned             AssignedRooms   `json:"assigned"`
		Unassigned           AssignedRooms   `json:"unassigned"`
		CardsOnFile          CardsOnFile     `json:"cardsOnFile"`
		CustomFields         CustomFields    `json:"customFields"`
		SourceName           string          `json:"sourceName"` // Source of reservation
		ThirdPartyIdentifier string          `json:"thirdPartyIdentifier"`
		GuestList            map[string]struct {
			GuestID string `json:"guestID"` // ID of the guest

			GuestName                   string            `json:"guestName"`
			GuestFirstName              string            `json:"guestFirstName"`
			GuestLastName               string            `json:"guestLastName"`
			Guestgender                 string            `json:"guestGender"` // Toegestane waarden: "M", "F", "N/A"
			GuestEmail                  string            `json:"guestEmail"`
			GuestPhone                  string            `json:"guestPhone"`
			GuestCellPhone              string            `json:"guestCellPhone"`
			GuestAddress                string            `json:"guestAddress"`
			GuestAddress2               string            `json:"guestAddress2"`
			GuestCity                   string            `json:"guestCity"`
			GuestState                  string            `json:"guestState"`
			GuestStatus                 string            `json:"guestStatus"` // Toegestane waarden: "active", "inactive"
			GuestCountry                string            `json:"guestCountry"`
			GuestZip                    string            `json:"guestZip"`
			GuestBirthdate              Date              `json:"guestBirthdate"`
			GuestDocumentType           string            `json:"guestDocumentType"`
			GuestDocumentNumber         string            `json:"guestDocumentNumber"`
			GuestDocumentIssueDate      Date              `json:"guestDocumentIssueDate"`
			GuestDocumentIsseingCountry string            `json:"guestDocumentIssuingCountry"`
			GuestDocumentExpirationDate Date              `json:"guestDocumentExpirationDate"`
			TaxID                       string            `json:"taxID"`        //  Guest's tax ID
			CompanyTaxID                string            `json:"companyTaxID"` // Guest's company tax ID
			CompanyName                 string            `json:"companyName"`  // Guest's company name
			SubReservationID            string            `json:"subReservationID"`
			StartDate                   Date              `json:"startDate"`
			EndDate                     Date              `json:"endDate"`
			AssignedRoom                bool              `json:"assignedRoom"` // Returns true if guest has roomed assigned, false if not
			RoomID                      string            `json:"roomID"`       // Room ID where guest is assigned
			RoomName                    string            `json:"roomName"`     // Room Name where guest is assigned
			RoomTypeName                string            `json:"roomTypeName"` // Room Name where guest is assigned
			IsMainGuest                 bool              `json:"isMainGuest"`
			UnassignedRooms             Rooms             `json:"unassignedRooms"` // List of all rooms that guest is assigned to but not yet checked in
			Rooms                       Rooms             `json:"rooms"`           // List of all rooms that guest is assigned to
			GuestRequirements           GuestRequirements `json:"guestRequirements,omitzero"`
			CustomFields                CustomFields      `json:"customFields"`
			IsAnonymized                bool              `json:"isAnonymized"` //  Flag indicating the guest data was removed upon request
		} `json:"guestList"`
		IsAnonymized bool `json:"isAnonymized"` //  Flag indicating the guest data was removed upon request
	} `json:"data"`
}

func (r *GetReservationRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/v1.3/getReservation", r.PathParams())
}

func (r *GetReservationRequest) Do() (GetReservationResponseBody, error) {
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
