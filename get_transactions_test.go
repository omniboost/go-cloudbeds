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
			And:
		},
	})

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

// func TestGetTransactionsAll(t *testing.T) {
// 	client := client()
// 	req := client.NewGetTransactionsRequest()

// 	resp, err := req.All()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	b, _ := json.MarshalIndent(resp, "", "  ")
// 	log.Println(string(b))
// }
