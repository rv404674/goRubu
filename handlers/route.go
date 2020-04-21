package handlers

import (
	"encoding/json"
	"fmt"
	"goRubu/middlewares"
	service "goRubu/services"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Response struct {
	Message      string
	ShortenedUrl string
}

type RedirectionResp struct {
	Message     string
	OriginalUrl string
}

type inputUrl struct {
	Url string `json:"Url"`
}

//To test a basic handler
//open this endpoint through browser simultaneously
func Hellohandler(w http.ResponseWriter, r *http.Request) {
	// see the imported package above
	fmt.Fprintf(w, "Hello. Go is an Awesome Language")
	time.Sleep(5 * time.Second)
}

func Basichandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Basic Handler")
	//errorResponse is a struct defined above
	json.NewEncoder(w).Encode(Response{Message: "Used to test Basic Middleware", ShortenedUrl: ""})
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
		fmt.Fprintf(w, "Enter a url to be shortened")
		//json.NewEncoder(w).Encode(Response{Message: "Enter a url", ShortenedUrl: ""})
	} else {
		shortenedUrl := service.CreateShortenedUrl(input.Url)
		fmt.Fprintf(w, shortenedUrl)
		//json.NewEncoder(w).Encode(Response{Message: "Success", ShortenedUrl: shortenedUrl})
	}
}

func RedirectionHandler(w http.ResponseWriter, r *http.Request) {
	var input inputUrl
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("Enter the redirection url")
	}

	json.Unmarshal(reqBody, &input)
	if input.Url == "" {
		fmt.Fprintf(w, "Enter some url")
		//	json.NewEncoder(w).Encode(RedirectionResp{Message: "Enter a url", OriginalUrl: ""})
	} else {
		orgUrl := service.UrlRedirection(input.Url)
		if orgUrl == "" {
			orgUrl = "This shortened Url doesn't exist in DB"
		}
		fmt.Fprintf(w, orgUrl)
		// 	originalUrl := service.UrlRedirection(input.U
		// 	json.NewEncoder(w).Encode(RedirectionResp{Message: "Success", OriginalUrl: originalUr
	}

}

func New() http.Handler {
	//gorilla mux, supports addition of a middleware to a route.
	route := mux.NewRouter()
	route.HandleFunc("/check", Hellohandler)
	route.HandleFunc("/shorten_url", CreateUrlHandler).Methods("POST")
	route.HandleFunc("/redirect", RedirectionHandler).Methods("POST")

	//special route to use middleware
	// when we want to access endpoints having basicMiddleware (or can be a simple Auth),
	// we need to add "/basic" before the route name.
	basicRoute := route.PathPrefix("/middleware").Subrouter()
	basicRoute.Use(middlewares.BasicMiddleware)
	basicRoute.HandleFunc("/check", Basichandler)

	return route
}
