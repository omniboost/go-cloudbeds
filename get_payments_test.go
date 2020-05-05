package cloudbeds_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-cloudbeds"
)

func TestGetPayments(t *testing.T) {
	client := client()
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	req := client.NewGetPaymentsRequest()
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	// yesterday := today.AddDate(0, 0, -1)
	// req.QueryParams().CreatedFrom = cloudbeds.DateTime{yesterday}
	lastYear := today.AddDate(-1, 0, 0)
	req.QueryParams().CreatedFrom = cloudbeds.DateTime{lastYear}
	req.QueryParams().CreatedTo = cloudbeds.DateTime{today}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
