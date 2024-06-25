package integration_test

import (
	"io"
	"net/http"
	"testing"
)

func TestApplicationEndpoints(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/restaurants?price_range=$&zip_code=23834", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Received response: %s", body)
}
