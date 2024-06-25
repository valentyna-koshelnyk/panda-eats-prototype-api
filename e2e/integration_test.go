package e2e

import (
	v1 "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1"
	"net/http"

	"net/http/httptest"
	"testing"
)

func TestFindAllRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/restaurants", nil)
	if err != nil {
		t.Fatal(err)
	}
	newRecorder := httptest.NewRecorder()
	v1.Routes(HTTPController).ServeHTTP(newRecorder, req)
	statusCode := 200
	if newRecorder.Result().StatusCode != statusCode {
		t.Errorf("TestInfoRequest() test returned an unexpected result: got %v want %v", newRecorder.Result().StatusCode, statusCode)
	}
}
