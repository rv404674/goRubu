package main

import (
	"log"
	"net/http"

	"goRubu/handlers"
	"goRubu/services"

	"github.com/jasonlvhit/gocron"
)

func executeCronJob() {
	log.Println("**Executing Cron Service")
	gocron.Every(5).Minute().Do(services.RemovedExpiredEntries)
	<-gocron.Start()
}

func main() {
	log.Printf("Inside GoRubu. Starting the project")

	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server, Listening at %s", server.Addr)

	// start your cron service in a go routine
	// go executeCronJob()

	// GO thing. You can declare and then use the same variable if.
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%s", err)
	} else {
		log.Printf("Server Closed")
	}

}
