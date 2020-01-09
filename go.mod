module curve-exchangerates

go 1.13

replace curve-exchangerates/controllers v0.0.0 => ./controllers

require (
	github.com/gorilla/mux v1.7.3
	gopkg.in/go-playground/assert.v1 v1.2.1
)
