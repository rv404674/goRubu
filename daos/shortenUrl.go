package daos

import (
	"context"
	"log"
	"os"

	dbconnection "goRubu/database"
	model "goRubu/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var db_name string
var collection1_name string
var collection2_name string
var client *mongo.Client

func init() {
	if err := godotenv.Load("variables.env"); err != nil {
		log.Fatal("Error: No .env file found")
	}

	db_name = os.Getenv("DB_NAME")
	collection1_name = os.Getenv("COLLECTION1_NAME")
	collection2_name = os.Getenv("COLLECTION2_NAME")

	client = dbconnection.CreateCon()
}

// this one inserts in "shortened_url" connections
// this is using call by value i.e creating copy of urlModel everytime,
// if json is large, try to prevent this
func InsertInShortenedUrl(urlModel model.UrlModel) {

	collection := client.Database(db_name).Collection(collection1_name)
	insertResult, err := collection.InsertOne(context.TODO(), urlModel)

	if err != nil {
		log.Println("Error while writing to shortened_url collection", err)
	}

	log.Println("InsertedId", insertResult.InsertedID)

}
