package tests

import (
	"os"
	"testing"
)

// Here is where I would define the mock to test the HTTP responses from the
// controllers test properly after implementing GetLatestRate and GetHistoricRates
// as methods of the interface InternalClient (see client/exchange_api.go)

// TestMain ...
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
