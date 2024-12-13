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
		method:      http.MethodPost,
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
}

func (p GetTransactionsQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
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
	Filters   Filters `json:"filters,omitempty"`
	PageToken string  `json:"pageToken,omitempty"`
	Limit     int     `json:"limit,omitempty"`
	Sort      []Sort  `json:"sort,omitempty"`
}
type Filters struct {
	And []any `json:"and,omitempty"`
	Or  []any `json:"or,omitempty"`
}
type Sort struct {
	Field     string `json:"field,omitempty"`
	Direction string `json:"direction,omitempty"`
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
		PropertyID                     Int         `json:"propertyID"`                     // Property ID
		ReservationID                  string      `json:"reservationID"`                  // Reservation ID
		SubReservationID               string      `json:"subReservationID"`               // Sub Reservation ID
		HouseAccountID                 interface{} `json:"houseAccountID"`                 // House Account ID
		HouseAccountName               string      `json:"houseAccountName"`               // House Account Name
		GuestID                        string      `json:"guestID"`                        // Guest ID
		PropertyName                   string      `json:"propertyName"`                   // Property Name
		TransactionDateTime            DateTime    `json:"transactionDateTime"`            // DateTime that the transaction was stored
		TransactionDateTimeUTC         DateTime    `json:"transactionDateTimeUTC"`         // DateTime that the transaction was stored, in UTC timezone
		TransactionModifiedDateTime    DateTime    `json:"transactionModifiedDateTime"`    // DateTime that the transaction was last modified
		TransactionModifiedDateTimeUTC DateTime    `json:"transactionModifiedDateTimeUTC"` // DateTime that the transaction was slast modified, in UTC timezone
		GuestCheckin                   Date        `json:"guestCheckin"`                   // Reservation Check-in date
		GuestCheckout                  Date        `json:"guestCheckout"`                  // Reservation Check-out date
		RoomTypeID                     string      `json:"roomTypeID"`                     // ID of the room type
		RoomTypeName                   string      `json:"roomTypeName"`                   // Name of the room type
		RoomName                       string      `json:"roomName"`                       // Name of the specific room. N/A means not applicable, and it is used if the transaction is not linked to a room.
		GuestName                      string      `json:"guestName"`                      // Name of the first guest of the reservation
		Description                    string      `json:"description"`                    // Description of the transaction
		Category                       string      `json:"category"`                       // Category of the transaction
		TransactionCode                string      `json:"transactionCode"`                // Transaction identifier that can be used, or left blank
		Notes                          string      `json:"notes"`                          // If any special information needs to be added to the transaction, it will be in this field
		Quantity                       Int         `json:"quantity"`                       // Consolidated amount on the transaction (Credit - Debit)
		Currency                       string      `json:"currency"`                       // Currency of the transaction
		Username                       string      `json:"userName"`                       // User responsible for creating the transaction
		TransactionType                string      `json:"transactionType"`                // Consolidated transaction type. Toegestane waarden: debit, credit
		TransactionCategory            string      `json:"transactionCategory"`
		ItemCategoryName               string      `json:"itemCategoryName"` // Transaction category. Toegestane waarden: adjustment, addon, custom_item, fee, payment, product, rate, room_revenue, refund, tax, void
		TransactionID                  string      `json:"transactionID"`    // Transaction identifier
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
	return r.client.GetEndpointURL("accounting/v1.0/deposits/transactions", r.PathParams())
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

// func (r *GetTransactionsRequest) All() (GetTransactionsResponseBody, error) {
// 	r.QueryParams().PageNumber = 1
// 	resp, err := r.Do()
// 	if err != nil {
// 		return resp, err
// 	}

// 	concat := GetTransactionsResponseBody{
// 		Count:   resp.Count,
// 		Total:   resp.Total,
// 		Success: true,
// 		Message: "",
// 		Data:    resp.Data,
// 	}

// 	for concat.Count < concat.Total {
// 		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
// 		resp, err := r.Do()
// 		if err != nil {
// 			return resp, err
// 		}

// 		concat.Count = concat.Count + resp.Count
// 		concat.Total = resp.Total
// 		concat.Data = append(concat.Data, resp.Data...)
// 	}

// 	return concat, nil
// }
