package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetFiscalDocumentsRequest() GetFiscalDocumentsRequest {
	return GetFiscalDocumentsRequest{
		client:      c,
		queryParams: c.NewGetFiscalDocumentsQueryParams(),
		pathParams:  c.NewGetFiscalDocumentsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetFiscalDocumentsRequestBody(),
	}
}

type GetFiscalDocumentsRequest struct {
	client      *Client
	queryParams *GetFiscalDocumentsQueryParams
	pathParams  *GetFiscalDocumentsPathParams
	method      string
	headers     http.Header
	requestBody GetFiscalDocumentsRequestBody
}

func (c *Client) NewGetFiscalDocumentsQueryParams() *GetFiscalDocumentsQueryParams {
	return &GetFiscalDocumentsQueryParams{}
}

const SOURCE_KIND_GROUP_PROFILE = "GROUP_PROFILE"
const SOURCE_KIND_RESERVATION = "RESERVATION"
const SOURCE_KIND_HOUSE_ACCOUNT = "HOUSE_ACCOUNT"
const SOURCE_KIND_ACCOUNTS_RECEIVABLE_LEDGER = "ACCOUNTS_RECEIVABLE_LEDGER"

type GetFiscalDocumentsQueryParams struct {
	PageToken string `schema:"pageToken,omitempty"`
	Limit     int    `schema:"limit,omitempty"`
	// 1 to 100
	// Defaults to 20
	Sort string `schema:"sort,omitempty"`
	// Supported fields:
	// createdAt
	// dueDate
	// invoiceDate
	// kind
	// status
	// Supported sort modes asc:desc. If not supplied default is asc.

	IDs []string `schema:"ids,omitempty"`
	// List of IDs to filter by
	SourceIDs []string `schema:"sourceIDs,omitempty"`
	// List of source IDs to filter by
	SourceIdentifiers []string `schema:"sourceIdentifiers,omitempty"`
	// List of source-specific identifiers
	SourceKind string `schema:"sourceKind,omitempty"`
	// Filter by source kind
	//
	// Allowed:
	// GROUP_PROFILE
	// RESERVATION
	// HOUSE_ACCOUNT
	// ACCOUNTS_RECEIVABLE_LEDGER
	SourceKinds []string `json:"sourceKinds,omitempty"`
	// Filter by source kind
	NumberContains string `schema:"numberContains,omitempty"`
	// Filter by document number partial match
	Statuses []string `schema:"statuses,omitempty"`
	// List of fiscal document statuses to filter by
	Kinds []string `schema:"kinds,omitempty"`
	// List of fiscal document kinds to filter by
	CreatedAtFrom DateTime `schema:"createdAtFrom,omitempty"`
	// Creation date-time range start
	CreatedAtTo DateTime `schema:"createdAtTo,omitempty"`
	// Creation date-time range end
	InvoiceDateFrom Date `schema:"invoiceDateFrom,omitempty"`
	// deprecated
	// Invoice date range start
	InvoiceDateTo Date `schema:"invoiceDateTo,omitempty"`
	// deprecated
	// Invoice date range end
	InvoiceDatePropertyTimezoneFrom Date `schema:"invoiceDatePropertyTimezoneFrom,omitempty"`
	// Invoice date range start
	InvoiceDatePropertyTimezoneTo Date `schema:"invoiceDatePropertyTimezoneTo,omitempty"`
	// Invoice date range end
	DueDateFrom Date `schema:"dueDateFrom,omitempty"`
	// deprecated
	// Due date range start
	DueDateTo Date `schema:"dueDateTo,omitempty"`
	// deprecated
	// Due date range end
	DueDatePropertyTimezoneFrom Date `schema:"dueDatePropertyTimezoneFrom,omitempty"`
	// Due date range start
	DueDatePropertyTimezoneTo Date `schema:"dueDatePropertyTimezoneTo,omitempty"`
	// Due date range end
	AmountFrom float64 `schema:"amountFrom,omitempty"`
	// Minimum document amount
	AmountTo float64 `schema:"amountTo,omitempty"`
	// Maximum document amount
	BalanceFrom float64 `schema:"balanceFrom,omitempty"`
	// Minimum document balance
	BalanceTo float64 `schema:"balanceTo,omitempty"`
	// Maximum document balance
}

func (p GetFiscalDocumentsQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetFiscalDocumentsRequest) QueryParams() *GetFiscalDocumentsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetFiscalDocumentsPathParams() *GetFiscalDocumentsPathParams {
	return &GetFiscalDocumentsPathParams{}
}

type GetFiscalDocumentsPathParams struct {
}

func (p *GetFiscalDocumentsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetFiscalDocumentsRequest) PathParams() *GetFiscalDocumentsPathParams {
	return r.pathParams
}

func (r *GetFiscalDocumentsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFiscalDocumentsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetFiscalDocumentsRequestBody() GetFiscalDocumentsRequestBody {
	return GetFiscalDocumentsRequestBody{}
}

type GetFiscalDocumentsRequestBody struct {
}

func (r *GetFiscalDocumentsRequest) RequestBody() *GetFiscalDocumentsRequestBody {
	return &r.requestBody
}

func (r *GetFiscalDocumentsRequest) SetRequestBody(body GetFiscalDocumentsRequestBody) {
	r.requestBody = body
}

func (r *GetFiscalDocumentsRequest) NewResponseBody() *GetFiscalDocumentsResponseBody {
	return &GetFiscalDocumentsResponseBody{}
}

type GetFiscalDocumentsResponseBody struct {
	NextPageToken string `json:"nextPageToken"`

	FiscalDocuments []struct {
		ID                          string  `json:"id"`                          // Fiscal document unique identifier
		Number                      string  `json:"number"`                      // Fiscal document number
		PropertyID                  string  `json:"propertyId"`                  // Property unique identifier
		UserID                      string  `json:"userId"`                      // User unique identifier who created the fiscal document
		UserFullName                string  `json:"userFullName"`                // User full name who created the fiscal document
		SourceName                  string  `json:"sourceName"`                  // Source of the fiscal document
		SourceID                    string  `json:"sourceId"`                    // Source unique identifier
		SourceKind                  string  `json:"sourceKind"`                  // Source kind (GROUP_PROFILE RESERVATION HOUSE_ACCOUNT ACCOUNTS_RECEIVABLE_LEDGER )
		Kind                        string  `json:"kind"`                        // Kind of fiscal document (INVOICE CREDIT_NOTE RECEIPT RECTIFY_INVOICE PRO_FORMA_INVOICE REFUND_RECEIPT INVOICE_RECEIPT)
		InvoiceDate                 Date    `json:"invoiceDate"`                 // Invoice date (deprecated)
		InvoiceDatePropertyTimezone Date    `json:"invoiceDatePropertyTimezone"` // Invoice date in property timezone
		FileName                    string  `json:"fileName"`                    // Name of the fiscal document file
		Amount                      float64 `json:"amount"`                      // Total amount of the fiscal document
		Balance                     float64 `json:"balance"`                     // Remaining balance of the fiscal document
		DueDate                     Date    `json:"dueDate"`                     // Due date (deprecated)
		DueDatePropertyTimezone     Date    `json:"dueDatePropertyTimezone"`     // Due date in property timezone

		Recipients []struct {
			ID          string `json:"id"`          // Recipient unique identifier
			FirstName   string `json:"firstName"`   // Recipient first name
			LastName    string `json:"lastName"`    // Recipient last name
			Email       string `json:"email"`       // Recipient email address
			Type        string `json:"type"`        // Recipient type (COMPANY PERSON MANUAL)
			CompanyName string `json:"companyName"` // Company name if recipient is a company
		} `json:"recipients"` // List of recipients associated with the fiscal document

		Status     string   `json:"status"`     // Fiscal document status (COMPLETED VOIDED PAID PENDING_INTEGRATION PARTIALLY_PAID COMPLETED_INTEGRATION FAILED_INTEGRATION CORRECTION_NEEDED CANCELED CANCEL_REQUESTED OPEN REQUESTED VOID_REQUESTED FAILED MANUALLY_RECONCILED REJECTED ACCEPTED PENDING_TRANSACTION)
		Origin     string   `json:"origin"`     // Origin of the fiscal document
		ExternalID string   `json:"externalId"` // External identifier related to the fiscal document
		FailReason string   `json:"failReason"` // Reason for failure if the fiscal document generation failed
		Method     string   `json:"method"`     // Method used for the fiscal document (VOID ADJUSTMENT)
		CreatedAt  DateTime `json:"createdAt"`  // Creation date-time of the fiscal document
		ParentID   string   `json:"parentId"`   // Parent fiscal document ID if applicable
		UpdatedAt  DateTime `json:"updatedAt"`  // Last update date-time of the fiscal document

		GovernmentIntegration struct {
			Number string `json:"number"` // Government integration number
			Series string `json:"series"` // Government integration series
			Status string `json:"status"` // Government integration status
			QR     struct {
				URL    string `json:"url"`    // QR code URL
				String string `json:"string"` // QR code string representation
			}
			URL                              string `json:"url"`                              // URL to access the fiscal document on the government portal
			OfficialID                       string `json:"officialId"`                       // Official ID provided by the government
			ExternalID                       string `json:"externalId"`                       // External ID related to the government integration
			RectifyingInvoiceType            string `json:"rectifyingInvoiceType"`            // Type of rectifying invoice if applicable
			CancellationFailedFallbackStatus string `json:"cancellationFailedFallbackStatus"` // Status of the fiscal document (COMPLETED VOIDED PAID PENDING_INTEGRATION PARTIALLY_PAID COMPLETED_INTEGRATION FAILED_INTEGRATION CORRECTION_NEEDED CANCELED CANCEL_REQUESTED OPEN REQUESTED VOID_REQUESTED FAILED MANUALLY_RECONCILED REJECTED ACCEPTED PENDING_TRANSACTION)
			PDFFileBase64                    string `json:"pdfFileBase64"`                    // Base64-encoded PDF file content. Only allowed when status is COMPLETED_INTEGRATION.
			Handwritten                      bool   `json:"handwritten"`                      // Indicates this is a handwritten receipt created during POS unavailability.
		}
		LatestLinkedDocument struct {
			ID        string   `json:"id"`        // ID of the latest linked document
			Number    string   `json:"number"`    // Number of the latest linked document
			CreatedAt DateTime `json:"createdAt"` // Creation date of the latest linked document
			Kind      string   `json:"kind"`      // Kind of fiscal document (INVOICE CREDIT_NOTE RECEIPT RECTIFY_INVOICE PRO_FORMA_INVOICE REFUND_RECEIPT INVOICE_RECEIPT)
			Status    string   `json:"status"`    // Status of the fiscal document (COMPLETED VOIDED PAID PENDING_INTEGRATION PARTIALLY_PAID COMPLETED_INTEGRATION FAILED_INTEGRATION CORRECTION_NEEDED CANCELED CANCEL_REQUESTED OPEN REQUESTED VOID_REQUESTED FAILED MANUALLY_RECONCILED REJECTED ACCEPTED PENDING_TRANSACTION)
		} // Information about the latest document in a rectification chain

		LinkedDocuments []struct {
			ID               string   `json:"id"`               // ID of the linked document
			Number           string   `json:"number"`           // Number of the linked document
			CreatedAt        DateTime `json:"createdAt"`        // Creation date of the linked document
			Kind             string   `json:"kind"`             // Kind of fiscal document (INVOICE CREDIT_NOTE RECEIPT RECTIFY_INVOICE PRO_FORMA_INVOICE REFUND_RECEIPT INVOICE_RECEIPT)
			Status           string   `json:"status"`           // Status of the fiscal document (COMPLETED VOIDED PAID PENDING_INTEGRATION PARTIALLY_PAID COMPLETED_INTEGRATION FAILED_INTEGRATION CORRECTION_NEEDED CANCELED CANCEL_REQUESTED OPEN REQUESTED VOID_REQUESTED FAILED MANUALLY_RECONCILED REJECTED ACCEPTED PENDING_TRANSACTION)
			IsLatest         bool     `json:"isLatest"`         // Whether this is the latest document in the chain
			RelationshipType string   `json:"relationshipType"` // The relationship type - PARENT means this document is linked to the current document, CHILD means the current document is linked to this one (PARENT CHILD)
		} // List of documents linked to this fiscal document (both parent and child relationships)

		Actions []struct {
			Type string `json:"type"` // Action that can be performed on a fiscal document (CANCEL RECTIFY DOWNLOAD CREDIT_NOTE VOID VOID_AND_REFUND ADD_PAYMENT APPLY_TO_INVOICE)
		} // Returns the list of actions available for the transaction

		SourceIdentifier string `json:"sourceIdentifier"` // Reservation Identifier or a group code
		Simplified       bool   `json:"simplified"`       // Applies to invoices only.

	} `json:"fiscalDocuments"`
}

func (r *GetFiscalDocumentsRequest) URL() url.URL {
	return r.client.GetEndpointURL("fiscal-document/v1/fiscal-documents", r.PathParams())
}

func (r *GetFiscalDocumentsRequest) Do() (GetFiscalDocumentsResponseBody, error) {
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

func (r *GetFiscalDocumentsRequest) All() (GetFiscalDocumentsResponseBody, error) {
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
		concat.FiscalDocuments = append(concat.FiscalDocuments, resp.FiscalDocuments...)
	}

	return concat, nil
}
