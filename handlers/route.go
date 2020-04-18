package handlers

import (
	"encoding/json"
	"fmt"
	"goRubu/middlewares"
	service "goRubu/services"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type errorResponse struct {
	Message string
}

type inputUrl struct {
	Url string `json:"Url"`
}

//To test a basic handler
func Hellohandler(w http.ResponseWriter, r *http.Request) {
	// see the imported package above
	fmt.Fprintf(w, "Hello. Go is an Awesome Language")
}

func Basichandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Basic Handler")
	//errorResponse is a struct defined above
	json.NewEncoder(w).Encode(errorResponse{Message: "Used to test Basic Middleware"})
}

//created new shortened Url
func CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	var input inputUrl
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// enter url:"" pair in body
		log.Println("Enter the url you want to shorten")
	}

	// get go object from json
	json.Unmarshal(reqBody, &input)
	if input.Url == "" {
		json.NewEncoder(w).Encode(errorResponse{Message: "Enter a url"})
	} else {
		service.CreateShortenedUrl(input.Url)
	}

}

func New() http.Handler {
	//gorilla mux, supports addition of a middleware to a route.
	route := mux.NewRouter()
	route.HandleFunc("/check", Hellohandler)
	route.HandleFunc("/shorten_url", CreateUrlHandler).Methods("POST")

	//special route to use middleware
	// when we want to access endpoints having basicMiddleware (or can be a simple Auth),
	// we need to add "/basic" before the route name.
	basicRoute := route.PathPrefix("/middleware").Subrouter()
	basicRoute.Use(middlewares.BasicMiddleware)
	basicRoute.HandleFunc("/check", Basichandler)

	return route
}
