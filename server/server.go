package server

import (
	"curve-exchangerates/controllers"
	"os"
)

var server = controllers.Server{}

// Run ...
func Run() {

	// This is where I would initialize a database/cache layer like Redis to, for example, keep the latest
	// responses from https://api.exchangerates.io and minimize requests. In such case, the Server struct would
	// have additional fields like DB (*gorm.DB) and/or Redis
	server.Initialize()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(port)
}
