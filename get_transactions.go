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
		ID                              string      `json:"id"`
		PropertyID                      string      `json:"propertyId"`
		InternalTransactionCode         string      `json:"internalTransactionCode"`
		CustomTransactionCode           interface{} `json:"customTransactionCode"`
		GeneralLedgerCustomCode         interface{} `json:"generalLedgerCustomCode"`
		Amount                          float64     `json:"amount"`
		Currency                        string      `json:"currency"`
		CustomerID                      string      `json:"customerId"`
		RootID                          string      `json:"rootId"`
		ParentID                        interface{} `json:"parentId"`
		SourceID                        string      `json:"sourceId"`
		SubSourceID                     string      `json:"subSourceId"`
		SourceKind                      string      `json:"sourceKind"`
		Account                         interface{} `json:"account"`
		ExternalRelationID              string      `json:"externalRelationId"`
		ExternalRelationKind            string      `json:"externalRelationKind"`
		OriginID                        string      `json:"originId"`
		RoutedFrom                      interface{} `json:"routedFrom"`
		Quantity                        int         `json:"quantity"`
		Description                     string      `json:"description"`
		UserID                          string      `json:"userId"`
		SourceDatetime                  time.Time   `json:"sourceDatetime"`
		TransactionDatetime             time.Time   `json:"transactionDatetime"`
		TransactionDatetimePropertyTime time.Time   `json:"transactionDatetimePropertyTime"`
		ServiceDate                     string      `json:"serviceDate"`
		CreatedAt                       time.Time   `json:"createdAt"`
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
