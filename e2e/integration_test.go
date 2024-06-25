package integration_test

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

var (
	userJSON = `{
    "email": "newUser@example.com",
    "password": "password1234!"
}`
	quantity = `{"quantity" : 3}`
)

func Test_RestaurantsListing(t *testing.T) {
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

func Test_MenuListing(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/menu/restaurant/3", nil)
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

func Test_Registration(t *testing.T) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/auth/signup", strings.NewReader(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Received response: %s", body)
}

func Test_Login(t *testing.T) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/auth/login", strings.NewReader(userJSON))
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

func Test_AddItem(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/cart/item/12", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3ZmZhMmMtODYzMC00YjBkLWFmMzQtNzhlYzhlN2FmMmY3IiwiZXhwIjoxNzE5MzM4MTQyfQ.X6jJ0D7mlsiEh29XnNBAG6YG4QN6FJrkZ_P2r4pEFEo")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Received response: %s", body)
}

func Test_GetCartItem(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/cart", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3ZmZhMmMtODYzMC00YjBkLWFmMzQtNzhlYzhlN2FmMmY3IiwiZXhwIjoxNzE5MzM4MTQyfQ.X6jJ0D7mlsiEh29XnNBAG6YG4QN6FJrkZ_P2r4pEFEo")
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

func Test_RemoveItems(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:3000/api/v1/cart/item/12", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3ZmZhMmMtODYzMC00YjBkLWFmMzQtNzhlYzhlN2FmMmY3IiwiZXhwIjoxNzE5MzM4MTQyfQ.X6jJ0D7mlsiEh29XnNBAG6YG4QN6FJrkZ_P2r4pEFEo")
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

func Test_UpdateItemQuantity(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", "http://localhost:3000/api/v1/cart/item/12", strings.NewReader(quantity))
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3ZmZhMmMtODYzMC00YjBkLWFmMzQtNzhlYzhlN2FmMmY3IiwiZXhwIjoxNzE5MzM4MTQyfQ.X6jJ0D7mlsiEh29XnNBAG6YG4QN6FJrkZ_P2r4pEFEo")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Received response: %s", body)
}

func Test_CreateOrder(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/order", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3ZmZhMmMtODYzMC00YjBkLWFmMzQtNzhlYzhlN2FmMmY3IiwiZXhwIjoxNzE5MzM4MTQyfQ.X6jJ0D7mlsiEh29XnNBAG6YG4QN6FJrkZ_P2r4pEFEo")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Received response: %s", body)
}

func Test_ShipOrder(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", "http://localhost:3000/api/v1/order/b7bJlJLRMGpyoC3LXbaSt/shipping", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3ZmZhMmMtODYzMC00YjBkLWFmMzQtNzhlYzhlN2FmMmY3IiwiZXhwIjoxNzE5MzM4MTQyfQ.X6jJ0D7mlsiEh29XnNBAG6YG4QN6FJrkZ_P2r4pEFEo")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Received response: %s", body)
}

func Test_DeliverOrder(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", "http://localhost:3000/api/v1/order/b7bJlJLRMGpyoC3LXbaSt/delivery", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3ZmZhMmMtODYzMC00YjBkLWFmMzQtNzhlYzhlN2FmMmY3IiwiZXhwIjoxNzE5MzM4MTQyfQ.X6jJ0D7mlsiEh29XnNBAG6YG4QN6FJrkZ_P2r4pEFEo")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Received response: %s", body)
}

func Test_GetOrderHistory(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/order/history", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3ZmZhMmMtODYzMC00YjBkLWFmMzQtNzhlYzhlN2FmMmY3IiwiZXhwIjoxNzE5MzM4MTQyfQ.X6jJ0D7mlsiEh29XnNBAG6YG4QN6FJrkZ_P2r4pEFEo")
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
