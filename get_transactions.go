package cloudbeds

import (
	"net/http"
	"net/url"
	"time"
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
	Transactions []struct {
		// Id of the transaction.
		ID string `json:"id"`
		// Property ID where the transaction was created.
		PropertyID string `json:"propertyId"`
		// Internal code for the transaction, managed by Cloudbeds.
		InternalTransactionCode string `json:"internalTransactionCode"`
		// Custom code for the transaction, managed by Property.
		CustomTransactionCode string `json:"customTransactionCode"`
		// Custom code for general ledger, managed by Property.
		GeneralLedgerCustomCode string `json:"generalLedgerCustomCode"`
		// Amount of the transaction.
		Amount float64 `json:"amount"`
		// Number of decimal places for the currency.
		CurrencyScale int `json:"currencyScale"`
		// Currency (ISO code) applied to the amount of the transaction.
		Currency string `json:"currency"`
		// Id of the user who perform the transaction, also know as guest id.
		CustomerID string `json:"customerId"`
		// Root Id of the transaction, it contains the id of the transaction that is related to it.
		RootID string `json:"rootId"`
		// Id of the transaction that is parent of this one. For example Tax on top of a rate, tax on top of a fee, etc.
		ParentID string `json:"parentId"`
		// Id of the source. It is related with the source_kind, so if sourceKind is RESERVATION, is the reservation id.
		SourceID string `json:"sourceId"`
		// Id of the sub source. At the moment only for reservations that is the booking_room_id.
		SubSourceID string `json:"subSourceId"`
		// Source Kind. At the moment only available:
		// - RESERVATION
		// - GROUP_PROFILE
		// - HOUSE_ACCOUNT
		// - ACCOUNTS_RECEIVABLE.
		SourceKind string  `json:"sourceKind"`
		Account    Account `json:"account,omitzero"`
		// External relation id, for example if the transaction is a payment it will contain payment id.
		ExternalRelationID string `json:"externalRelationId"`
		// Kind of the external relation id, for example if transaction is a payment it will contain PAYMENT.
		// - ROOM
		// - PAYMENT
		// - ITEM
		// - ITEM_POS
		// - ADDON
		// - RESERVATION
		// - ACCOUNTS_RECEIVABLE
		// - ROOM_REVENUE
		// - TAX
		// - FEE
		// - ADJUSTMENT
		// - PAYMENT_FEE
		ExternalRelationKind string `json:"externalRelationKind"`
		// Id of origin of the transaction. For example if the transaction is created based on a rate, is the rate id.
		OriginID string `json:"originId"`
		// Id of the transaction that was routed from. It can be null.
		RoutedFrom interface{} `json:"routedFrom"`
		// Amount of items purchased.
		Quantity int `json:"quantity"`
		// Description of the transaction.
		Description string `json:"description"`
		// ID of the user who created the transaction
		UserID string `json:"userId"`
		// Date time the source was created. (ISO 8601) in UTC
		SourceDatetime time.Time `json:"sourceDatetime"`
		// Date time when the transaction should be created at. (ISO 8601) in UTC
		TransactionDatetime time.Time `json:"transactionDatetime"`
		// Date time when the transaction should be created at base on the property timezone.
		TransactionDatetimePropertyTime time.Time `json:"transactionDatetimePropertyTime"`
		// Date when the posted transaction was created (property time).
		ServiceDate string `json:"serviceDate"`
		// Date time when the transaction was inserted on the database. (ISO 8601) in UTC
		CreatedAt time.Time `json:"createdAt"`
		// if source_kind = RESERVATION, this field will contain a reservation identifier. For a transaction with source_kind = GROUP_PROFILE, this field will contain a group code. For source_king = HOUSE_ACCOUNT it will be null.
		SourceIdentifier string `json:"sourceIdentifier"`
		// identifier of a booking room
		SubSourceIdentifier string `json:"subSourceIdentifier"`
	} `json:"transactions"`
	NextPageToken string `json:"nextPageToken"`
}

func (r *GetTransactionsRequest) URL() url.URL {
	return r.client.GetEndpointURL("accounting/v1.0/transactions", r.PathParams())
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

func (r *GetTransactionsRequest) All() (GetTransactionsResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := resp

	for resp.NextPageToken != "" {
		r.RequestBody().PageToken = resp.NextPageToken
		resp, err = r.Do()
		if err != nil {
			return resp, err
		}
		concat.Transactions = append(concat.Transactions, resp.Transactions...)
	}

	return concat, nil
}
