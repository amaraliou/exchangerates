package controllers

import "net/http"

import "curve-exchangerates/client"

import "curve-exchangerates/responses"

import "curve-exchangerates/utils"

// GetExchangeRate ...
func GetExchangeRate(writer http.ResponseWriter, request *http.Request) {

	base := request.URL.Query().Get("base")
	latestRate, err := client.GetLatestRate(base)
	if err != nil {
		responses.ERROR(writer, http.StatusUnprocessableEntity, err)
		return
	}

	shouldBuy, err := utils.ShouldExchange(base)
	if err != nil {
		responses.ERROR(writer, http.StatusUnprocessableEntity, err)
		return
	}

	latestRate.ShouldBuy = shouldBuy

	responses.JSON(writer, http.StatusOK, latestRate)
}
