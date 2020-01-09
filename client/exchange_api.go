package client

import (
	"curve-exchangerates/models"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

// RatesBaseURL ...
const RatesBaseURL = "https://api.exchangeratesapi.io/latest?base=%s&symbols=EUR"

// HistoricRatesBaseURL ...
const HistoricRatesBaseURL = "https://api.exchangeratesapi.io/history?start_at=%s&end_at=%s&base=%s&symbols=EUR"

// APIClient ...
var APIClient = &http.Client{
	Timeout: 15 * time.Second,
}

// InternalClient to later build a mock for testing
type InternalClient interface {
	GetLatestRate(base string) (models.Rate, error)
	GetHistoricRates(base string, start, end time.Time) (models.HistoricRates, error)
}

// GetLatestRate ...
func GetLatestRate(base string) (models.Rate, error) {

	latestRate := models.Rate{}

	if base != "USD" && base != "GBP" && base != "" {
		return models.Rate{}, errors.New("Base currency given is not allowed. Allowed currencies are: GBP, USD")
	}

	if base == "" {
		// GBP as default, in the remaining cases it's either GBP or USD
		base = "GBP"
	}

	requestURL := fmt.Sprintf(RatesBaseURL, base)
	response, err := APIClient.Get(requestURL)
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return models.Rate{}, errors.New("Request to https://api.exchangerates.io timed out. Retry later")
	}

	if err != nil {
		return models.Rate{}, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&latestRate)
	if err != nil {
		return models.Rate{}, err
	}

	return latestRate, err
}

// GetHistoricRates ...
func GetHistoricRates(base string, start, end time.Time) (models.HistoricRates, error) {

	historicRates := models.HistoricRates{}

	if base != "USD" && base != "GBP" && base != "" {
		return models.HistoricRates{}, errors.New("Base currency given is not allowed. Allowed currencies are: GBP, USD")
	}

	if base == "" {
		// GBP as default, in the remaining cases it's either GBP or USD
		base = "GBP"
	}

	if end.Sub(start).Hours() < 0 {
		return models.HistoricRates{}, errors.New("Start date is after end date")
	}

	startDate := fmt.Sprintf("%d-%d-%d", start.Year(), start.Month(), start.Day())
	endDate := fmt.Sprintf("%d-%d-%d", end.Year(), end.Month(), end.Day())

	requestURL := fmt.Sprintf(HistoricRatesBaseURL, startDate, endDate, base)
	response, err := APIClient.Get(requestURL)
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return models.HistoricRates{}, errors.New("Request to https://api.exchangerates.io timed out. Retry later")
	}

	if err != nil {
		return models.HistoricRates{}, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&historicRates)
	if err != nil {
		return models.HistoricRates{}, err
	}

	return historicRates, nil
}
