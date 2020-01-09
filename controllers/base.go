package controllers

import (
	"curve-exchangerates/middlewares"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server ..
type Server struct {
	Router *mux.Router
}

// Initialize ...
func (server *Server) Initialize() {

	server.Router = mux.NewRouter()
	// If we were to extend to multiple routes, I'd create another file routes.go with a function to initialize the routes
	server.Router.HandleFunc("/rates", middlewares.SetMiddlewareJSON(GetExchangeRate)).Methods("GET")
}

// Run ... making my linter happy
func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s\n", addr)
	log.Fatal(http.ListenAndServe(":"+addr, server.Router))
}
