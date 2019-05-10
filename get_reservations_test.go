package cloudbeds_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-cloudbeds"
)

func TestGetReservations(t *testing.T) {
	client := client()
	req := client.NewGetReservationsRequest()
	req.QueryParams().IncludeGuestsDetails = true
	req.QueryParams().Status = cloudbeds.StatusCheckedIn

	resp, err := req.Do()
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

	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
