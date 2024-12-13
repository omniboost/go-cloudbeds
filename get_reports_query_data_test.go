package cloudbeds_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetReportsQueryData(t *testing.T) {
	client := client()
	req := client.NewGetReportsQueryDataRequest()
	req.QueryParams().Mode = "Preview"

	json.Unmarshal([]byte(`{
    "property_ids": [
        "316633"
    ],
    "dataset_id": 6,
    "columns": [
        {
            "cdf": {
                "type": "default",
                "column": "invoice_number"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "id"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "invoice_generate_datetime"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "bill_to_party"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "bill_to_tax_id"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "bill_to_id"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "bill_to_address_line_1"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "bill_to_city"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "bill_to_state"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "bill_to_postal_code"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "bill_to_country_code"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "reservation_number"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "reservation_checkin_date"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "reservation_checkout_date"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "invoice_type"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "invoice_status"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "payment_due_date"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "invoice_currency_code"
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "total_gross_amount"
            },
            "metrics": [
                "sum"
            ]
        },
        {
            "cdf": {
                "type": "default",
                "column": "taxes_value_amount"
            },
            "metrics": [
                "sum"
            ]
        },
        {
            "cdf": {
                "type": "default",
                "column": "balance_due_amount"
            },
            "metrics": [
                "sum"
            ]
        },
        {
            "cdf": {
                "type": "default",
                "column": "description",
                "multi_level_id": 3
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "quantity",
                "multi_level_id": 3
            },
            "metrics": [
                "sum"
            ]
        },
        {
            "cdf": {
                "type": "default",
                "column": "amount",
                "multi_level_id": 3
            },
            "metrics": [
                "sum"
            ]
        },
        {
            "cdf": {
                "type": "default",
                "column": "tax_total",
                "multi_level_id": 3
            },
            "metrics": [
                "sum"
            ]
        },
        {
            "cdf": {
                "type": "default",
                "column": "tax_type",
                "multi_level_id": 3
            }
        },
        {
            "cdf": {
                "type": "default",
                "column": "tax_total_percentage",
                "multi_level_id": 3
            },
            "metrics": [
                "sum"
            ]
        }
    ],
    "group_rows": null,
    "group_columns": null,
    "custom_cdfs": null,
    "filters": {
        "and": [
            {
                "cdf": {
                    "type": "default",
                    "column": "invoice_generate_datetime"
                },
                "operator": "greater_than_or_equal",
                "value": "2024-10-01T03:00:00"
            },
            {
                "cdf": {
                    "type": "default",
                    "column": "invoice_generate_datetime"
                },
                "operator": "less_than_or_equal",
                "value": "2024-12-12T03:00:00"
            }
        ]
    },
    "sort": null,
    "settings": {
        "details": true,
        "totals": false,
        "transpose": false
    },
    "periods": null,
    "comparisons": null,
    "formats": null,
    "custom_field_cdfs": []
}`), req.RequestBody())

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
