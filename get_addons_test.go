package cloudbeds_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetAddons(t *testing.T) {
	client := client()
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	req := client.NewGetAddonsRequest()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
