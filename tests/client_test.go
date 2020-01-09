package tests

import (
	"curve-exchangerates/client"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetLatestRate(t *testing.T) {

	samples := []struct {
		base                 string
		expectedBase         string
		expectedConvCurrency string
		errorMessage         string
	}{
		{
			base:                 "GBP",
			expectedBase:         "GBP",
			expectedConvCurrency: "EUR",
			errorMessage:         "",
		},
		{
			base:                 "",
			expectedBase:         "GBP",
			expectedConvCurrency: "EUR",
			errorMessage:         "",
		},
		{
			base:                 "USD",
			expectedBase:         "USD",
			expectedConvCurrency: "EUR",
			errorMessage:         "",
		},
		{
			base:                 "HKD",
			expectedBase:         "",
			expectedConvCurrency: "",
			errorMessage:         "Base currency given is not allowed. Allowed currencies are: GBP, USD",
		},
	}

	for _, v := range samples {

		latestRate, err := client.GetLatestRate(v.base)
		if err != nil {
			assert.Equal(t, err.Error(), v.errorMessage)
		} else {
			assert.Equal(t, latestRate.Base, v.expectedBase)
			_, found := latestRate.Rates["EUR"]
			assert.Equal(t, found, true)
		}

	}
}
