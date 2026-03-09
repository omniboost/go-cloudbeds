package cloudbeds_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-cloudbeds"
)

func TestGetReservations(t *testing.T) {
	client := client()
	req := client.NewGetReservationsRequest()
	req.QueryParams().IncludeGuestsDetails = true
	req.QueryParams().ResultsFrom = cloudbeds.DateTime{Time: time.Date(2024, 12, 30, 3, 0, 0, 0, time.UTC)}
	req.QueryParams().ResultsTo = cloudbeds.DateTime{Time: time.Date(2026, 12, 31, 3, 0, 0, 0, time.UTC)}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetReservationsAll(t *testing.T) {
	client := client()
	req := client.NewGetReservationsRequest()
	req.QueryParams().IncludeGuestsDetails = true
	// req.QueryParams().Status = cloudbeds.StatusCheckedIn

	resp, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
