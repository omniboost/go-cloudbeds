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
		method:      http.MethodPost,
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
	Format string `schema:"format,omitempty"`
	Mode   string `schema:"mode"`
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
	PropertyIds  []string          `json:"property_ids"`
	DatasetID    int               `json:"dataset_id"`
	Columns      []QueryDataColumn `json:"columns"`
	GroupRows    interface{}       `json:"group_rows"`
	GroupColumns interface{}       `json:"group_columns"`
	CustomCdfs   interface{}       `json:"custom_cdfs"`
	Filters      QueryDataFilters  `json:"filters"`
	Sort         interface{}       `json:"sort"`
	Settings     struct {
		Details   bool `json:"details"`
		Totals    bool `json:"totals"`
		Transpose bool `json:"transpose"`
	} `json:"settings"`
	Periods         interface{}   `json:"periods"`
	Comparisons     interface{}   `json:"comparisons"`
	Formats         interface{}   `json:"formats"`
	CustomFieldCdfs []interface{} `json:"custom_field_cdfs"`
}

type QueryDataCDF struct {
	Type         string `json:"type"`
	Column       string `json:"column"`
	MultiLevelID int    `json:"multi_level_id,omitempty"`
}

type QueryDataColumn struct {
	Cdf     QueryDataCDF `json:"cdf,omitempty"`
	Metrics []string     `json:"metrics,omitempty"`
}

type QueryDataFilters struct {
	And []QueryDataAnd `json:"and,omitempty"`
	Or  []QueryDataAnd `json:"or,omitempty"`
}

type QueryDataAnd struct {
	CDF      QueryDataCDF `json:"cdf"`
	Operator string       `json:"operator"`
	Value    string       `json:"value"`
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
	Headers         []string         `json:"headers"`
	Index           [][]string       `json:"index"`
	GroupRows       []string         `json:"group_rows"`
	GroupColumns    []string         `json:"group_columns"`
	Periods         []string         `json:"periods"`
	Records         map[string][]any `json:"records"`
	Subtotals       any              `json:"subtotals"`
	Totals          any              `json:"totals"`
	Type            string           `json:"type"`
	Comparisons     []string         `json:"comparisons"`
	AggregatedCount int              `json:"aggregated_count"`
}

func (r *GetReportsQueryDataRequest) URL() url.URL {
	return r.client.GetEndpointURL("datainsights/v1.1/reports/query/data", r.PathParams())
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
