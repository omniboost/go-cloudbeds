package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetFiscalDocumentRecipientsRequest() GetFiscalDocumentRecipientsRequest {
	return GetFiscalDocumentRecipientsRequest{
		client:      c,
		queryParams: c.NewGetFiscalDocumentRecipientsQueryParams(),
		pathParams:  c.NewGetFiscalDocumentRecipientsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetFiscalDocumentRecipientsRequestBody(),
	}
}

type GetFiscalDocumentRecipientsRequest struct {
	client      *Client
	queryParams *GetFiscalDocumentRecipientsQueryParams
	pathParams  *GetFiscalDocumentRecipientsPathParams
	method      string
	headers     http.Header
	requestBody GetFiscalDocumentRecipientsRequestBody
}

func (c *Client) NewGetFiscalDocumentRecipientsQueryParams() *GetFiscalDocumentRecipientsQueryParams {
	return &GetFiscalDocumentRecipientsQueryParams{}
}

type GetFiscalDocumentRecipientsQueryParams struct {
	PageToken string `schema:"pageToken,omitempty"`
	Limit     int    `schema:"limit,omitempty"`
	// 1 to 100
	// Defaults to 20
	Sort string `schema:"sort,omitempty"`
	// Supported fields:
	// createdAt, serviceDate, sourceId, recipientDate, internalCode
	// Supported sort modes asc:desc. If not supplied default is asc.

	IncludeLinkedDocumentRecipients bool `schema:"includeLinkedDocumentRecipients,omitempty"`
	// Defaults to false
	// Include recipients from linked documents.

	FolioIDs []int64 `schema:"folioIds,omitempty"`
	// Filter by folio IDs.
}

func (p GetFiscalDocumentRecipientsQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetFiscalDocumentRecipientsRequest) QueryParams() *GetFiscalDocumentRecipientsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetFiscalDocumentRecipientsPathParams() *GetFiscalDocumentRecipientsPathParams {
	return &GetFiscalDocumentRecipientsPathParams{}
}

type GetFiscalDocumentRecipientsPathParams struct {
	ID string `schema:"id"`
}

func (p *GetFiscalDocumentRecipientsPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetFiscalDocumentRecipientsRequest) PathParams() *GetFiscalDocumentRecipientsPathParams {
	return r.pathParams
}

func (r *GetFiscalDocumentRecipientsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFiscalDocumentRecipientsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetFiscalDocumentRecipientsRequestBody() GetFiscalDocumentRecipientsRequestBody {
	return GetFiscalDocumentRecipientsRequestBody{}
}

type GetFiscalDocumentRecipientsRequestBody struct {
}

func (r *GetFiscalDocumentRecipientsRequest) RequestBody() *GetFiscalDocumentRecipientsRequestBody {
	return &r.requestBody
}

func (r *GetFiscalDocumentRecipientsRequest) SetRequestBody(body GetFiscalDocumentRecipientsRequestBody) {
	r.requestBody = body
}

func (r *GetFiscalDocumentRecipientsRequest) NewResponseBody() *GetFiscalDocumentRecipientsResponseBody {
	return &GetFiscalDocumentRecipientsResponseBody{}
}

//	{
//	 "recipients": [
//	   {
//	     "id": "string",
//	     "propertyId": "string",
//	     "sourceId": "string",
//	     "sourceKind": "GROUP_PROFILE",
//	     "recipientDate": "2026-02-18T09:07:47.925Z",
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
type GetFiscalDocumentRecipientsResponseBody []FiscalDocumentRecipient

func (r *GetFiscalDocumentRecipientsRequest) URL() url.URL {
	return r.client.GetEndpointURL("fiscal-document/v1/fiscal-documents/{{.id}}/recipients", r.PathParams())
}

func (r *GetFiscalDocumentRecipientsRequest) Do() (GetFiscalDocumentRecipientsResponseBody, error) {
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
