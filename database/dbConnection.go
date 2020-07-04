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
	log.Print(dir)
	envFile := "variables.env"
	if strings.Contains(dir, "test") {
		envFile = "../variables.env"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Error: No .env file found, dbCon.go ", err)
	}
}

// CreateCon - create db connection
// support both mongo - one on localhost and other on docker
func CreateCon() *mongo.Client {
	var dbDomain = os.Getenv("DB_DOMAIN_DOCKER")

	clientOptions := options.Client().ApplyURI(dbDomain)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("Connection Failed while connecting to mongo container. Error: ", err)
		clientOptions = options.Client().ApplyURI(os.Getenv("DB_DOMAIN_LOCALHOST"))
		client, err = mongo.Connect(context.TODO(), clientOptions)

		err2 := client.Ping(context.TODO(), nil)

		if err2 != nil {
			log.Fatal("Connection to both Mongo Container and Local Mongo Failed")
		}

	}

	log.Println("Connected to Mongo!")
	return client
}
