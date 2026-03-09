package cloudbeds_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetReservation(t *testing.T) {
	client := client()
	req := client.NewGetReservationRequest()
	req.QueryParams().PropertyID = client.PropertyID()
	req.QueryParams().ReservationID = "2397774183890"
	req.QueryParams().IncludeGuestRequirements = true

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
