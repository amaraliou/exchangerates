package controllers

import (
	"curve-exchangerates/responses"
	"net/http"
)

// Home ...
func Home(writer http.ResponseWriter, request *http.Request) {
	responses.JSON(writer, http.StatusOK, "Welcome To This Awesome API")
}
