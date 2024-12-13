package cloudbeds_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-cloudbeds"
)

func TestGetTransactions(t *testing.T) {
	client := client()
	req := client.NewGetTransactionsRequest()
	req.SetRequestBody(cloudbeds.GetTransactionsRequestBody{
		Filters: cloudbeds.Filters{
			And: []cloudbeds.And{
				{
					Operator: "greater_than_or_equal",
					Value:    "2024-10-01T03:00:00",
					Field:    "TRANSACTION_DATETIME",
				},
				{
					Operator: "less_than_or_equal",
					Value:    "2024-12-12T03:00:00",
					Field:    "TRANSACTION_DATETIME",
				},
			},
		},
	})

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetTransactionsAll(t *testing.T) {
	client := client()
	req := client.NewGetTransactionsRequest()
	req.SetRequestBody(cloudbeds.GetTransactionsRequestBody{
		Filters: cloudbeds.Filters{
			And: []cloudbeds.And{
				{
					Operator: "greater_than_or_equal",
					Value:    "2024-10-01T03:00:00",
					Field:    "TRANSACTION_DATETIME",
				},
				{
					Operator: "less_than_or_equal",
					Value:    "2024-12-12T03:00:00",
					Field:    "TRANSACTION_DATETIME",
				},
			},
		},
		Limit: 10,
	})

	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
