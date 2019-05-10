package cloudbeds_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-cloudbeds"
)

func TestGetTransactions(t *testing.T) {
	client := client()
	req := client.NewGetTransactionsRequest()
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	yesterday := today.AddDate(0, 0, -1)
	req.QueryParams().ResultsFrom = cloudbeds.Date{yesterday}
	req.QueryParams().ResultsTo = cloudbeds.Date{today}

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
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tomorrow := today.AddDate(0, 0, 1)
	yesterday := today.AddDate(0, 0, -1)
	req.QueryParams().ResultsFrom = cloudbeds.Date{yesterday}
	req.QueryParams().ResultsTo = cloudbeds.Date{tomorrow}

	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
