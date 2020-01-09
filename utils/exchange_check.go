package utils

import "curve-exchangerates/client"

import "curve-exchangerates/models"

import "time"

// ShouldExchange ...
func ShouldExchange(base string) (bool, error) {

	latestRate, err := client.GetLatestRate(base)
	if err != nil {
		return false, err
	}

	historicRates, err := client.GetHistoricRates(base, time.Now().AddDate(0, 0, -7), time.Now())

	avgHistoricRate := getAvgRate(historicRates)

	if latestRate.Rates["EUR"] > avgHistoricRate {
		return true, nil
	}

	return false, nil
}

func getAvgRate(rates models.HistoricRates) float64 {

	var sumOfRates = 0.0
	for _, rate := range rates.Rates {
		sumOfRates += rate["EUR"]
	}
	return sumOfRates / float64(len(rates.Rates))
}
