package cloudbeds

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetReportsQueryDataRequest() GetReportsQueryDataRequest {
	return GetReportsQueryDataRequest{
		client:      c,
		queryParams: c.NewGetReportsQueryDataQueryParams(),
		pathParams:  c.NewGetReportsQueryDataPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetReportsQueryDataRequestBody(),
	}
}

type GetReportsQueryDataRequest struct {
	client      *Client
	queryParams *GetReportsQueryDataQueryParams
	pathParams  *GetReportsQueryDataPathParams
	method      string
	headers     http.Header
	requestBody GetReportsQueryDataRequestBody
}

func (c *Client) NewGetReportsQueryDataQueryParams() *GetReportsQueryDataQueryParams {
	return &GetReportsQueryDataQueryParams{}
}

type GetReportsQueryDataQueryParams struct {
	PropertyID        string            `schema:"propertyID,omitempty"`
	IncludeDebit      bool              `schema:"includeDebit,omitempty"`
	IncludeCredit     bool              `schema:"includeCredit,omitempty"`
	IncludeDeleted    bool              `schema:"includeDeleted,omitempty"`
	ReservationID     string            `schema:"reservationID,omitempty"`
	SubReservationID  string            `schema:"subReservationID,omitempty"`
	RoomID            string            `schema:"roomID,omitempty"`
	GuestID           int               `schema:"guestID,omitempty"`
	HouseAccountID    int               `schema:"houseAccountID,omitempty"`
	ResultsFrom       Date              `schema:"resultsFrom,omitempty"`
	ResultsTo         Date              `schema:"resultsTo,omitempty"`
	ModifiedFrom      Date              `schema:"modifiedFrom,omitempty"`
	ModifiedTo        Date              `schema:"modifiedTo,omitempty"`
	CreatedFrom       DateTime          `schema:"createdFrom,omitempty"`
	CreatedTo         DateTime          `schema:"createdTo,omitempty"`
	TransactionFilter TransactionFilter `schema:"transactionFilter,omitempty"`
	PageNumber        int               `schema:"pageNumber,omitempty"`
	PageSize          int               `schema:"pageSize,omitempty"`
	SortBy            string            `schema:"sortBy,omitempty"`
	OrderBy           string            `schema:"orderBy,omitempty"`
}

func (p GetReportsQueryDataQueryParams) ToURLValues() (url.Values, error) {
	encoder := NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetReportsQueryDataRequest) QueryParams() *GetReportsQueryDataQueryParams {
	return r.queryParams
}

func (c *Client) NewGetReportsQueryDataPathParams() *GetReportsQueryDataPathParams {
	return &GetReportsQueryDataPathParams{}
}

type GetReportsQueryDataPathParams struct {
}

func (p *GetReportsQueryDataPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetReportsQueryDataRequest) PathParams() *GetReportsQueryDataPathParams {
	return r.pathParams
}

func (r *GetReportsQueryDataRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetReportsQueryDataRequest) Method() string {
	return r.method
}

func (s *Client) NewGetReportsQueryDataRequestBody() GetReportsQueryDataRequestBody {
	return GetReportsQueryDataRequestBody{}
}

type GetReportsQueryDataRequestBody struct {
}

func (r *GetReportsQueryDataRequest) RequestBody() *GetReportsQueryDataRequestBody {
	return &r.requestBody
}

func (r *GetReportsQueryDataRequest) SetRequestBody(body GetReportsQueryDataRequestBody) {
	r.requestBody = body
}

func (r *GetReportsQueryDataRequest) NewResponseBody() *GetReportsQueryDataResponseBody {
	return &GetReportsQueryDataResponseBody{}
}

type GetReportsQueryDataResponseBody struct {
	Success bool   `json:"success"`
	Count   Int    `json:"count"`
	Total   Int    `json:"total"`
	Message string `json:"message"`
	Data    []struct {
		PropertyID                     Int         `json:"propertyID"`
		ReservationID                  string      `json:"reservationID"`
		SubReservationID               string      `json:"subReservationID"`
		HouseAccountID                 interface{} `json:"houseAccountID"`
		HouseAccountName               string      `json:"houseAccountName"`
		GuestID                        string      `json:"guestID"`
		PropertyName                   string      `json:"propertyName"`
		TransactionDateTime            DateTime    `json:"transactionDateTime"`
		TransactionDateTimeUTC         DateTime    `json:"transactionDateTimeUTC"`
		TransactionModifiedDateTime    DateTime    `json:"transactionModifiedDateTime"`
		TransactionModifiedDateTimeUTC DateTime    `json:"transactionModifiedDateTimeUTC"`
		GuestCheckin                   Date        `json:"guestCheckin"`
		GuestCheckout                  Date        `json:"guestCheckout"`
		RoomTypeID                     string      `json:"roomTypeID"`
		RoomTypeName                   string      `json:"roomTypeName"`
		RoomName                       string      `json:"roomName"`
		GuestName                      string      `json:"guestName"`
		Description                    string      `json:"description"`
		Category                       string      `json:"category"`
		TransactionCode                string      `json:"transactionCode"`
		Notes                          string      `json:"notes"`
		Quantity                       Int         `json:"quantity"`
		Currency                       string      `json:"currency"`
		Username                       string      `json:"userName"`
		TransactionType                string      `json:"transactionType"`
		TransactionCategory            string      `json:"transactionCategory"`
		ItemCategoryName               string      `json:"itemCategoryName"`
		TransactionID                  string      `json:"transactionID"`
		ParentTransactionID            string      `json:"parentTransactionID"`
		CardType                       string      `json:"cardType"`
		IsDeleted                      bool        `json:"isDeleted"`
		Amount                         float64     `json:"amount"`
	} `json:"data"`
}

func (r *GetReportsQueryDataRequest) URL() url.URL {
	return r.client.GetEndpointURL("getTransactions", r.PathParams())
}

func (r *GetReportsQueryDataRequest) Do() (GetReportsQueryDataResponseBody, error) {
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

func (r *GetReportsQueryDataRequest) All() (GetReportsQueryDataResponseBody, error) {
	r.QueryParams().PageNumber = 1
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := GetReportsQueryDataResponseBody{
		Count:   resp.Count,
		Total:   resp.Total,
		Success: true,
		Message: "",
		Data:    resp.Data,
	}

	for concat.Count < concat.Total {
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
		resp, err := r.Do()
		if err != nil {
			return resp, err
		}

		concat.Count = concat.Count + resp.Count
		concat.Total = resp.Total
		concat.Data = append(concat.Data, resp.Data...)
	}

	return concat, nil
}
