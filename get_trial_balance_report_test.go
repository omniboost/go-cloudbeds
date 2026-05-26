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

func TestGetTrialBalanceReport(t *testing.T) {
	client := client()
	req := client.NewGetTrialBalanceReportRequest()
	d, _ := time.Parse("2006-01-02", "2026-05-20")
	req.QueryParams().Date = cloudbeds.Date{Time: d}
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
