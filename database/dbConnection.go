// used to establish db connection, by reading data from a variables.env file

package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	if err := godotenv.Load("variables.env"); err != nil {
		log.Fatal("Error: No .env file found")
	}
}

// create db connection
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
