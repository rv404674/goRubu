package main

import (
	"log"
	"net/http"

	"goRubu/handlers"
)

func main() {

	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server, Listening at %s", server.Addr)

	// GO thing. You can declare and then use the same variable if.
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%s", err)
	} else {
		log.Printf("Server Closed")
	}

}
