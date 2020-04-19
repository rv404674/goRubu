package daos

import (
	"context"
	"log"
	"os"

	dbconnection "goRubu/database"
	"goRubu/models"
	model "goRubu/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var db_name string
var collection1_name string
var collection2_name string
var client *mongo.Client
var db *mongo.Database

func init() {
	if err := godotenv.Load("variables.env"); err != nil {
		log.Fatal("Error: No .env file found")
	}

	db_name = os.Getenv("DB_NAME")
	collection1_name = os.Getenv("COLLECTION1_NAME")
	collection2_name = os.Getenv("COLLECTION2_NAME")
	log.Println(collection2_name)

	client = dbconnection.CreateCon()

	// create Index on Id field
	indexOptions := options.Index().SetUnique(true)
	indexKeys := bsonx.MDoc{
		"uniqueid": bsonx.Int64(-1),
	}

	indexModel := mongo.IndexModel{
		Options: indexOptions,
		Keys:    indexKeys,
	}

	db := client.Database(db_name)
	collection1 := db.Collection(collection1_name)

	// we want index on 'shortened_url' which is collection1
	indexName, err := collection1.Indexes().CreateOne(context.Background(), indexModel)

	if err != nil {
		log.Fatal("Error while creating index", err)
	}

	log.Println("IndexName", indexName)
}

// this one inserts in "shortened_url" connections
// this is using call by value i.e creating copy of urlModel everytime,
// if json is large, try to prevent this
func InsertInShortenedUrl(urlModel model.UrlModel) {

	collection := client.Database(db_name).Collection(collection1_name)
	//collection := db.Collection(collection1_name)
	insertResult, err := collection.InsertOne(context.Background(), urlModel)

	if err != nil {
		log.Fatal("Error while writing to shortened_url collection", err)
	}

	log.Println("InsertedId", insertResult.InsertedID)
}

// has indexing on uniqueid field
// perform find on 'shortened_url' collections
func GetUrl(inputUniqueId int) models.UrlModel {
	collection := client.Database(db_name).Collection(collection1_name)

	filter := bson.D{{"uniqueid", inputUniqueId}}
	var result models.UrlModel

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("**ERROR while fetching using Id", err)
	}

	return result
}

// update counter field in second collections - incrementer
// also there will be an already existing value in db (i.e counter will start from)- 10000
// { "_id" : ObjectId("5e9b7c0e7b3a8740a2f828c4"), "uniqueid" : "counter", "value" : 10000 }
func GetCounterValue() int {
	// as there will be one row only
	collection := client.Database(db_name).Collection(collection2_name)
	filter := bson.D{{"uniqueid", "counter"}}
	var result models.IncrementerModel

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal("**ERROR while fetching counter value", err)
	}

	return result.Value
}

func UpdateCounter() {
	collection := client.Database(db_name).Collection(collection2_name)
	filter := bson.D{{"uniqueid", "counter"}}

	// $inc will increase value of counter by 1
	update := bson.D{
		{"$inc", bson.D{
			{"value", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("**ERROR while updating counter value", err)
	}

	log.Println("counter updated", updateResult.ModifiedCount)

}
