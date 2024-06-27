package e2e

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"io"
	"net/http"
	"strings"
	"testing"
)

var (
	userJSON = `{
    "email": "newUser21@example.com",
    "password": "password1234!"
}`
	quantity = `{"quantity" : 3}`
)

func Test_AllEndpoints(t *testing.T) {
	client := &http.Client{}
	tokenData := &entity.TokenData{}

	t.Run("Get All Restaurants", func(t *testing.T) {
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
	})

	t.Run("Get All Menu", func(t *testing.T) {
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
	})
	t.Run("Registration", func(t *testing.T) {
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
	})
	t.Run("Login", func(t *testing.T) {
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
		var data entity.CustomResponse
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		_ = json.Unmarshal(body, &data)
		tokenRAW := gjson.Get(string(body), "data.token").String()
		tokenData.Token = "Bearer " + tokenRAW
		t.Logf("Received response: %s", tokenData.Token)
	})
	t.Run("Add_Item", func(t *testing.T) {
		req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/cart/item/12", strings.NewReader(quantity))
		req.Header.Add("Authorization", tokenData.Token)

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
	})
	t.Run("Add_Item", func(t *testing.T) {
		req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/cart/item/43", strings.NewReader(quantity))
		req.Header.Add("Authorization", tokenData.Token)

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
	})
	t.Run("Get cart items", func(t *testing.T) {
		req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/cart", nil)
		req.Header.Add("Authorization", tokenData.Token)
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
	})
	t.Run("Remove items", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "http://localhost:3000/api/v1/cart/item/43", nil)
		req.Header.Add("Authorization", tokenData.Token)
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
	})
	t.Run("Update item quantity", func(t *testing.T) {
		req, err := http.NewRequest("PATCH", "http://localhost:3000/api/v1/cart/item/12", strings.NewReader(quantity))
		req.Header.Add("Authorization", tokenData.Token)
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
	})
	t.Run("Create order", func(t *testing.T) {
		req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/order", nil)
		req.Header.Add("Authorization", tokenData.Token)
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
	})
	t.Run("Ship order", func(t *testing.T) {
		req, err := http.NewRequest("PATCH", "http://localhost:3000/api/v1/order/pTSZIngsJbN0pGmlYFtyt/shipping", nil)
		req.Header.Add("Authorization", tokenData.Token)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusAccepted {
			t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Received response: %s", body)
	})
	t.Run("Delivery Order", func(t *testing.T) {
		req, err := http.NewRequest("PATCH", "http://localhost:3000/api/v1/order/pTSZIngsJbN0pGmlYFtyt/delivery", nil)
		req.Header.Add("Authorization", tokenData.Token)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusAccepted {
			t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Received response: %s", body)
	})
	t.Run("Get Order History", func(t *testing.T) {
		req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/order/history", nil)
		req.Header.Add("Authorization", tokenData.Token)
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
	})
}
