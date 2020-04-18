package handlers

import (
	"encoding/json"
	"fmt"
	"goRubu/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type errorResponse struct {
	Message string
}

//To test a basic handler
func Hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello. Go is an Awesome Language")
}

func Basichandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Basic Handler")
	//errorResponse is a struct defined above
	json.NewEncoder(w).Encode(errorResponse{Message: "Used to test Basic Middleware"})
}

func New() http.Handler {
	//gorilla mux, supports addition of a middleware to a route.
	route := mux.NewRouter()
	route.HandleFunc("/check", Hellohandler)

	//special route to use middleware
	// when we want to access endpoints having basicMiddleware (or can be a simple Auth),
	// we need to add "/basic" before the route name.
	basicRoute := route.PathPrefix("/middleware").Subrouter()
	basicRoute.Use(middlewares.BasicMiddleware)
	basicRoute.HandleFunc("/check", Basichandler)

	return route
}
