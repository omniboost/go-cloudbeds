package cloudbeds_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetFiscalDocumentRecipients(t *testing.T) {
	client := client()
	req := client.NewGetFiscalDocumentRecipientsRequest()
	req.PathParams().ID = "135421797875914"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

