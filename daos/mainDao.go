package daos

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	dbconnection "goRubu/database"
	"goRubu/models"
	model "goRubu/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

//DB_NAME - the db name
var DB_NAME string

//COLLECTION1_NAME - first collection name
var COLLECTION1_NAME string

//COLLECTION2_NAME - second collection name
var COLLECTION2_NAME string
var client *mongo.Client

func init() {
	dir, _ := os.Getwd()
	envFile := "variables.env"
	if strings.Contains(dir, "test") {
		envFile = "../variables.env"
		// TODO: Remove .. , it is a security threat if done from root
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Error: No .env file found, mainDao.go ", err)
	}

	DB_NAME = os.Getenv("DB_NAME")
	COLLECTION1_NAME = os.Getenv("COLLECTION1_NAME")
	COLLECTION2_NAME = os.Getenv("COLLECTION2_NAME")

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

	db := client.Database(DB_NAME)
	collection1 := db.Collection(COLLECTION1_NAME)

	// we want index on 'uniqueid field of 'shortened_url' which is collection1
	indexName, err := collection1.Indexes().CreateOne(context.Background(), indexModel)

	if err != nil {
		log.Fatal("Error while creating index ", err)
	}

	log.Println("IndexName", indexName)
}

// InsertInShortenedUrl - this one inserts in "shortened_url" collection.
// this is using call by value i.e creating copy of urlModel everytime,
// if json is large, try to prevent this
func InsertInShortenedUrl(urlModel model.UrlModel) {

	collection := client.Database(DB_NAME).Collection(COLLECTION1_NAME)
	insertResult, err := collection.InsertOne(context.Background(), urlModel)

	if err != nil {
		log.Fatal("Error while writing to shortened_url collection", err)
	}

	log.Println("InsertedId", insertResult.InsertedID)
}

// GetUrl - has indexing on uniqueid field
// perform find on 'shortened_url' collections
func GetUrl(inputUniqueId int) models.UrlModel {
	collection := client.Database(DB_NAME).Collection(COLLECTION1_NAME)

	filter := bson.D{primitive.E{Key: "uniqueid", Value: inputUniqueId}}
	var result models.UrlModel

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
	}

	return result
}

// CleanDb - Clear the expired entries from the main "shortened_url" mongo collection
func CleanDb(uid int) {
	collection := client.Database(DB_NAME).Collection(COLLECTION1_NAME)
	filter := bson.D{primitive.E{Key: "uniqueid", Value: uid}}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error while deleting a doc", err)
	}
	// if you directly do string(deleteResult.DeletedCount) for deleteResult.DeletedCount = 67. you will get C. Hence use strcov
	log.Println("** Deleted " + strconv.FormatInt(deleteResult.DeletedCount, 10) + " documents ")
}

// GetCounterValue - update counter field in second collection - incrementer
// { "_id" : ObjectId("5e9b7c0e7b3a8740a2f828c4"), "uniqueid" : "counter", "value" : 10000 }
func GetCounterValue() int {
	// as there will be one row only
	collection := client.Database(DB_NAME).Collection(COLLECTION2_NAME)
	filter := bson.D{primitive.E{Key: "uniqueid", Value: "counter"}}

	var result models.IncrementerModel
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("**ERROR while fetching counter value", err)
		// its first hit. You need to create a counter value.
		insertResult, _ := collection.InsertOne(context.Background(), models.IncrementerModel{UniqueId: "counter", Value: 10000})
		log.Println("Initialized counter value InsertedId: ", insertResult.InsertedID)
		_ = collection.FindOne(context.TODO(), filter).Decode(&result)
	}

	return result.Value
}

// UpdateCounter - update the counter value by 1
func UpdateCounter() {
	collection := client.Database(DB_NAME).Collection(COLLECTION2_NAME)
	filter := bson.D{primitive.E{Key: "uniqueid", Value: "counter"}}

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

// GetAll - will be used by CronService for db purging
func GetAll() *mongo.Cursor {
	collection := client.Database(DB_NAME).Collection(COLLECTION1_NAME)

	cur, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Fatal("**ERROR in find all operation")
	}

	return cur
}
