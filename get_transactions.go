package cloudbeds

import (
	"net/http"
	"net/url"
)

// Payment - getTransactions
// Get a list of transactions for a property, or list of properties, for the
// date range specified. If no date range or reservation is specified, it will
// return the transactions for the last 7 days, unless stated otherwise.

func (c *Client) NewGetTransactionsRequest() GetTransactionsRequest {
	return GetTransactionsRequest{
		client:      c,
		queryParams: c.NewGetTransactionsQueryParams(),
		pathParams:  c.NewGetTransactionsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetTransactionsRequestBody(),
	}
}

type GetTransactionsRequest struct {
	client      *Client
	queryParams *GetTransactionsQueryParams
	pathParams  *GetTransactionsPathParams
	method      string
	headers     http.Header
	requestBody GetTransactionsRequestBody
}

func (c *Client) NewGetTransactionsQueryParams() *GetTransactionsQueryParams {
	return &GetTransactionsQueryParams{}
}

type GetTransactionsQueryParams struct {
	// ID for the properties to be queried (comma-separated).
	// It can be omitted if the API key is single-property, or to get results from
	// all properties on an association.
	PropertyID string `json:"propertyID,omitempty"`
	// If the response should include debit transactions
	// Standaard waarde: true
	IncludeDebit bool `json:"includeDebit,omitempty"`
	// If the response should include credit transactions
	// Standaard waarde: true
	IncludeCredit bool `json:"includeCredit,omitempty"`
	// If the response should include deleted transactions
	// Standaard waarde: false
	IncludeDeleted bool `json:"includeDeleted,omitempty"`
	// Reservation Unique Identifier, used to filter transactions result
	// If reservationID is informed, and dates are not, all transactions with the
	// reservationID will be returned.
	ReservationID string `json:"reservationID,omitempty"`
	// Sub Reservation Identifier, used to filter transactions result
	SubReservationID string `json:"subReservationID,omitempty"`
	// Room ID, used to filter transactions result
	RoomID string `json:"roomID,omitempty"`
	// Guest ID, used to filter transactions result
	GuestID int `json:"guestID,omitempty"`
	// House Account ID, used to filter transactions result
	HouseAccountID int `json:"houseAccountID,omitempty"`
	// Inferior limit date, used to filter transactions result (posted transaction date)
	ResultsFrom Date `json:"resultsFrom,omitempty"`
	// Superior limit date, used to filter transactions result (posted transaction date)
	ResultsTo Date `json:"resultTo,omitempty"`
	// Inferior limit date, used to filter transactions result
	ModifiedFrom Date `json:"modifiedFrom,omitempty"`
	// Superior limit date, used to filter transactions result
	ModifiedTo Date `json:"modifiedTo,omitempty"`
	// Inferior limit datetime, used to filter transactions result (creation date of the transaction).
	// If informed, all other dates are ignored (except createdTo).
	// If createdFrom is informed, but createdTo is not, the call will return all results since this datetime.
	// Necessary only if createdTo is sent.
	// If time portion not given, assumes 00:00:00.
	CreatedFrom DateTime `json:"createdFrom,omitempty"`
	// Superior limit datetime, used to filter transactions result (creation date of the transaction).
	// If informed (together with createdFrom), all other dates are ignored.
	// If time portion not given, assumes 23:59:59.
	CreatedTo DateTime `json:"createdTo,omitempty"`
	// transactionFilter optioneel	String
	// Transaction filter is used to filter transactions result
	// Standaard waarde: simple_transactions,adjustments,adjustments_voids,voids,refunds
	TransactionFilter TransactionFilter `json:"transactionFilter,omitempty"`
	// Results page number
	// Standaard waarde: 1
	PageNumber int `json:"pageNumber,omitempty"`
	// Results page size. Max = 100
	// Standaard waarde: 100
	PageSize int `json:"pageSize,omitempty"`
	// Sort response results by field
	// Toegestane waarden: transactionDateTime, transactionModifiedDateTime, guestCheckIn, guestCheckOut
	SortBy string `json:"sortBy,omitempty"`
	// Order response in DESCending or ASCending order, used together with sortBy
	// Standaard waarde: desc
	// Toegestane waarden: desc, asc
	OrderBy string `json:"orderBy,omitempty"`
}

func (p GetTransactionsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTransactionsRequest) QueryParams() *GetTransactionsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetTransactionsPathParams() *GetTransactionsPathParams {
	return &GetTransactionsPathParams{}
}

type GetTransactionsPathParams struct {
}

func (p *GetTransactionsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTransactionsRequest) PathParams() *GetTransactionsPathParams {
	return r.pathParams
}

func (r *GetTransactionsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTransactionsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetTransactionsRequestBody() GetTransactionsRequestBody {
	return GetTransactionsRequestBody{}
}

type GetTransactionsRequestBody struct {
}

func (r *GetTransactionsRequest) RequestBody() *GetTransactionsRequestBody {
	return &r.requestBody
}

func (r *GetTransactionsRequest) SetRequestBody(body GetTransactionsRequestBody) {
	r.requestBody = body
}

func (r *GetTransactionsRequest) NewResponseBody() *GetTransactionsResponseBody {
	return &GetTransactionsResponseBody{}
}

type GetTransactionsResponseBody struct {
	Success bool   `json:"success"`
	Count   Int    `json:"count"`
	Total   Int    `json:"total"`
	Message string `json:"message"`
	Data    []struct {
		PropertyID                     Int      `json:"propertyID"`                     // Property ID
		ReservationID                  string   `json:"reservationID"`                  // Reservation ID
		SubReservationID               string   `json:"subReservationID"`               // Sub Reservation ID
		GuestID                        string   `json:"guestID"`                        // Guest ID
		PropertyName                   string   `json:"propertyName"`                   // Property Name
		TransactionDateTime            DateTime `json:"transactionDateTime"`            // DateTime that the transaction was stored
		TransactionDateTimeUTC         DateTime `json:"transactionDateTimeUTC"`         // DateTime that the transaction was stored, in UTC timezone
		TransactionModifiedDateTime    DateTime `json:"transactionModifiedDateTime"`    // DateTime that the transaction was last modified
		TransactionModifiedDateTimeUTC DateTime `json:"transactionModifiedDateTimeUTC"` // DateTime that the transaction was slast modified, in UTC timezone
		GuestCheckin                   Date     `json:"guestCheckin"`                   // Reservation Check-in date
		GuestCheckout                  Date     `json:"guestCheckout"`                  // Reservation Check-out date
		RoomTypeID                     string   `json:"roomTypeID"`                     // ID of the room type
		RoomTypeName                   string   `json:"roomTypeName"`                   // Name of the room type
		RoomName                       string   `json:"roomName"`                       // Name of the specific room. N/A means not applicable, and it is used if the transaction is not linked to a room.
		GuestName                      string   `json:"guestName"`                      // Name of the first guest of the reservation
		Description                    string   `json:"description"`                    // Description of the transaction
		Category                       string   `json:"category"`                       // Category of the transaction
		TransactionCode                string   `json:"transactionCode"`                // Transaction identifier that can be used, or left blank
		Notes                          string   `json:"notes"`                          // If any special information needs to be added to the transaction, it will be in this field
		Quantity                       Int      `json:"quantity"`                       // Consolidated amount on the transaction (Credit - Debit)
		Currency                       string   `json:"currency"`                       // Currency of the transaction
		Username                       string   `json:"userName"`                       // User responsible for creating the transaction
		TransactionType                string   `json:"transactionType"`                // Consolidated transaction type. Toegestane waarden: debit, credit
		TransactionCategory            string   `json:"transactionCategory"`            // Transaction category. Toegestane waarden: adjustment, addon, custom_item, fee, payment, product, rate, room_revenue, refund, tax, void
		TransactionID                  string   `json:"transactionID"`                  // Transaction identifier
		// Parent transaction identifier. Parent transaction is a transaction to which
		// this current transaction is strongly related to or derived from.
		// Example: Parent transaction to a room rate tax is a room rate.
		// This parent transaction ID will mostly be present on transactions that are
		// taxes, fees and voids. It will not be present on room rates, items, payments
		// and refunds.
		ParentTransactionID string  `json:"parentTransactionID"`
		CardType            string  `json:"cardType"` // Abbreviated name of credit card type
		IsDeleted           bool    `json:"isDeleted"`
		Amount              float64 `json:"amount"`
	} `json:"data"`
}

func (r *GetTransactionsRequest) URL() url.URL {
	return r.client.GetEndpointURL("getTransactions", r.PathParams())
}

func (r *GetTransactionsRequest) Do() (GetTransactionsResponseBody, error) {
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
