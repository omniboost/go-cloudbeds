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

func TestAccountingTrialBalanceReportGet(t *testing.T) {
	client := client()
	req := client.NewAccountingTrialBalanceReportGetRequest()
	req.QueryParams().Date = cloudbeds.Date{Time: time.Date(2024, 10, 11, 0, 0, 0, 0, time.Local)}
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
