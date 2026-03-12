package cloudbeds

import (
	"context"
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

	Transactions FiscalDocumentTransactions `json:"transactions"`
}

func (r *GetFiscalDocumentTransactionsRequest) URL() url.URL {
	return r.client.GetEndpointURL("fiscal-document/v1/fiscal-documents/{{.id}}/transactions", r.PathParams())
}

func (r *GetFiscalDocumentTransactionsRequest) Do(ctx context.Context) (GetFiscalDocumentTransactionsResponseBody, error) {
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

func (r *GetFiscalDocumentTransactionsRequest) All(ctx context.Context) (GetFiscalDocumentTransactionsResponseBody, error) {
	resp, err := r.Do(ctx)
	if err != nil {
		return resp, err
	}

	concat := resp

	for resp.NextPageToken != "" {
		r.QueryParams().PageToken = resp.NextPageToken
		resp, err = r.Do(ctx)
		if err != nil {
			return resp, err
		}
		concat.Transactions = append(concat.Transactions, resp.Transactions...)
	}

	return concat, nil
}
