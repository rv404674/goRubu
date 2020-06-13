// used to establish db connection, by reading data from a variables.env file

package database

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// normal "variables.env" was working for the application, but
	// go test was unable to find it.

	// when doing "go test ./tests -v", I am getting "pwd" as "/Users/home/goRubu/tests"
	// when doing make execute or go run main.go, I am getting "pwd" as
	// "Users/home/goRubu"
	dir, _ := os.Getwd()
	envFile := "variables.env"
	if strings.Contains(dir, "test") {
		envFile = "../variables.env"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Error: No .env file found, dbCon.go ", err)
	}
}

// CreateCon - create db connection
func CreateCon() *mongo.Client {
	var dbDomain = os.Getenv("DB_DOMAIN")

	clientOptions := options.Client().ApplyURI(dbDomain)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Connection Failed", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Mongo!")
	return client
}
