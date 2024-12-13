package cloudbeds_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetReportsQueryData(t *testing.T) {
	client := client()
	req := client.NewGetReportsQueryDataRequest()
	// req.RequestBody().PropertyIds = []int{316633}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "")
	log.Println(string(b))
}
