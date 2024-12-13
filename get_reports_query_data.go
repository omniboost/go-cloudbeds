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
	PropertyIds []int `json:"property_ids,omitempty"`
	DatasetID   int   `json:"dataset_id,omitempty"`

	Columns []struct {
		Cdf struct {
			Type   string `json:"type,omitempty"`
			Column string `json:"column,omitempty"`
		} `json:"cdf,omitempty"`

		Metrics []string `json:"metrics,omitempty"`
	} `json:"columns,omitempty"`

	GroupRows    string `json:"group_rows,omitempty"`
	GroupColumns string `json:"group_columns,omitempty"`
	CustomCdfs   string `json:"custom_cdfs,omitempty"`

	Filters struct {
		And []struct {
			Cdf struct {
				Type   string `json:"type,omitempty"`
				Column string `json:"column,omitempty"`
			} `json:"cdf,omitempty"`

			Operator string `json:"operator,omitempty"`
			Value    string `json:"value,omitempty"`

			Or []struct {
				Cdf struct {
					Type   string `json:"type,omitempty"`
					Column string `json:"column,omitempty"`
				} `json:"cdf,omitempty"`

				Operator string `json:"operator,omitempty"`
				Value    string `json:"value,omitempty"`
			} `json:"or,omitempty"`
		} `json:"and,omitempty"`
	} `json:"filters,omitempty"`

	Sort string `json:"sort,omitempty"`

	Settings struct {
		Details   bool `json:"details,omitempty"`
		Totals    bool `json:"totals,omitempty"`
		Transpose bool `json:"transpose,omitempty"`
	} `json:"settings,omitempty"`

	Periods string `json:"periods,omitempty"`

	Formats struct {
		Date string `json:"date,omitempty"`
		Link bool   `json:"link,omitempty"`
	} `json:"formats,omitempty"`

	Comparisons string `json:"comparisons,omitempty"`

	CustomFieldCdfs []struct {
		Name string `json:"name,omitempty"`

		Properties []struct {
			InternalName string `json:"internal_name,omitempty"`
			PropertyID   int    `json:"property_id,omitempty"`
		} `json:"properties,omitempty"`
	} `json:"custom_field_cdfs,omitempty"`
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
		GroupRows       string `json:"group_rows"`
		GroupColumns    string `json:"group_columns"`
		Periods         string `json:"periods"`
		Records         string `json:"records"`
		Subtotals       string `json:"subtotals"`
		Totals          string `json:"totals"`
		Type            string `json:"type"`
		Comparisons     string `json:"comparisons"`
		AggregatedCount int    `json:"aggregated_count"`
	}
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
