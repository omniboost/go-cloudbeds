package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetFiscalDocumentTransactionsRequest() GetFiscalDocumentTransactionsRequest {
	return GetFiscalDocumentTransactionsRequest{
		client:      c,
		queryParams: c.NewGetFiscalDocumentTransactionsQueryParams(),
		pathParams:  c.NewGetFiscalDocumentTransactionsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetFiscalDocumentTransactionsRequestBody(),
	}
}

type GetFiscalDocumentTransactionsRequest struct {
	client      *Client
	queryParams *GetFiscalDocumentTransactionsQueryParams
	pathParams  *GetFiscalDocumentTransactionsPathParams
	method      string
	headers     http.Header
	requestBody GetFiscalDocumentTransactionsRequestBody
}

func (c *Client) NewGetFiscalDocumentTransactionsQueryParams() *GetFiscalDocumentTransactionsQueryParams {
	return &GetFiscalDocumentTransactionsQueryParams{}
}

type GetFiscalDocumentTransactionsQueryParams struct {
	PageToken string `schema:"pageToken,omitempty"`
	Limit     int    `schema:"limit,omitempty"`
	// 1 to 100
	// Defaults to 20
	Sort string `schema:"sort,omitempty"`
	// Supported fields:
	// createdAt, serviceDate, sourceId, transactionDate, internalCode
	// Supported sort modes asc:desc. If not supplied default is asc.

	IncludeLinkedDocumentTransactions bool `schema:"includeLinkedDocumentTransactions,omitempty"`
	// Defaults to false
	// Include transactions from linked documents.

	FolioIDs []int64 `schema:"folioIds,omitempty"`
	// Filter by folio IDs.
}

func (p GetFiscalDocumentTransactionsQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetFiscalDocumentTransactionsRequest) QueryParams() *GetFiscalDocumentTransactionsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetFiscalDocumentTransactionsPathParams() *GetFiscalDocumentTransactionsPathParams {
	return &GetFiscalDocumentTransactionsPathParams{}
}

type GetFiscalDocumentTransactionsPathParams struct {
	ID string `schema:"id"`
}

func (p *GetFiscalDocumentTransactionsPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetFiscalDocumentTransactionsRequest) PathParams() *GetFiscalDocumentTransactionsPathParams {
	return r.pathParams
}

func (r *GetFiscalDocumentTransactionsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFiscalDocumentTransactionsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetFiscalDocumentTransactionsRequestBody() GetFiscalDocumentTransactionsRequestBody {
	return GetFiscalDocumentTransactionsRequestBody{}
}

type GetFiscalDocumentTransactionsRequestBody struct {
}

func (r *GetFiscalDocumentTransactionsRequest) RequestBody() *GetFiscalDocumentTransactionsRequestBody {
	return &r.requestBody
}

func (r *GetFiscalDocumentTransactionsRequest) SetRequestBody(body GetFiscalDocumentTransactionsRequestBody) {
	r.requestBody = body
}

func (r *GetFiscalDocumentTransactionsRequest) NewResponseBody() *GetFiscalDocumentTransactionsResponseBody {
	return &GetFiscalDocumentTransactionsResponseBody{}
}

//	{
//	 "transactions": [
//	   {
//	     "id": "string",
//	     "propertyId": "string",
//	     "sourceId": "string",
//	     "sourceKind": "GROUP_PROFILE",
//	     "transactionDate": "2026-02-18T09:07:47.925Z",
//	     "guestName": "string",
//	     "description": "string",
//	     "internalCode": "string",
//	     "amount": 0,
//	     "folioId": "string",
//	     "status": "PENDING",
//	     "paidAmount": 0,
//	     "allocations": [
//	       {
//	         "receiptNumber": "string"
//	       }
//	     ]
//	   }
//	 ],
//	 "nextPageToken": "string"
//	}
type GetFiscalDocumentTransactionsResponseBody struct {
	NextPageToken string `json:"nextPageToken"`

	Transactions []struct {
		ID                       string   `json:"id"`
		PropertyID               string   `json:"propertyId"`
		SourceID                 string   `json:"sourceId"`
		SourceIdentifier         string   `json:"sourceIdentifier"`
		SourceKind               string   `json:"sourceKind"` // Kind of the source entity (GROUP_PROFILE RESERVATION HOUSE_ACCOUNT ACCOUNTS_RECEIVABLE_LEDGER)
		TransactionDate          string   `json:"transactionDate"`
		GuestName                string   `json:"guestName"`
		Description              string   `json:"description"`
		InternalCode             string   `json:"internalCode"`
		Amount                   float64  `json:"amount"`
		AvailableAmount          float64  `json:"availableAmount"`
		DocumentFiscalizedAmount *float64 `json:"documentFiscalizedAmount"`
		FolioID                  string   `json:"folioId"`
		Status                   string   `json:"status"` // Status of the transaction - PENDING for unpaid transactions, POSTED for paid transactions

		PaidAmount float64 `json:"paidAmount"`

		Allocations []struct {
			ReceiptNumber string `json:"receiptNumber"`
		} `json:"allocations"`
	} `json:"transactions"`
}

func (r *GetFiscalDocumentTransactionsRequest) URL() url.URL {
	return r.client.GetEndpointURL("fiscal-document/v1/fiscal-documents/{{.id}}/transactions", r.PathParams())
}

func (r *GetFiscalDocumentTransactionsRequest) Do() (GetFiscalDocumentTransactionsResponseBody, error) {
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

func (r *GetFiscalDocumentTransactionsRequest) All() (GetFiscalDocumentTransactionsResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := resp

	for resp.NextPageToken != "" {
		r.QueryParams().PageToken = resp.NextPageToken
		resp, err = r.Do()
		if err != nil {
			return resp, err
		}
		concat.Transactions = append(concat.Transactions, resp.Transactions...)
	}

	return concat, nil
}
