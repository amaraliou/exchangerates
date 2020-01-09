package tests

import (
	"curve-exchangerates/controllers"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

// I could have gone with the approach of have the test files in the same package as the
// entity I am testing (for example this file staying in the controllers package) but I
// wanted to have them all in one place.

// Also could have use table tests to make the tests shorter

func TestGetExchangeRate(t *testing.T) {

	req, err := http.NewRequest("GET", "/rates", nil)
	if err != nil {
		t.Errorf("this is the error: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetExchangeRate)
	handler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	// Would implement a mock so that we know the responses to test
	assert.Equal(t, rr.Code, 200)
	assert.Equal(t, responseMap["base"], "GBP")
}

func TestGetExchangeRateBaseUSD(t *testing.T) {

	req, err := http.NewRequest("GET", "/rates?base=USD", nil)
	if err != nil {
		t.Errorf("this is the error: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetExchangeRate)
	handler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 200)
	assert.Equal(t, responseMap["base"], "USD")
}

func TestGetExchangeRateWrongBase(t *testing.T) {

	req, err := http.NewRequest("GET", "/rates?base=HKD", nil)
	if err != nil {
		t.Errorf("this is the error: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetExchangeRate)
	handler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 422)
	assert.Equal(t, responseMap["error"], "Base currency given is not allowed. Allowed currencies are: GBP, USD")
}
