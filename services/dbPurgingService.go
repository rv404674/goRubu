package services

import (
	"context"
	dao "goRubu/daos"
	model "goRubu/models"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var EXPIRY_TIME int

func init() {
	if err := godotenv.Load("variables.env"); err != nil {
		log.Fatal("Error in loading env file from dbPurging Service")
	}

	// in min
	EXPIRY_TIME, _ = strconv.Atoi(os.Getenv("EXPIRY_TIME"))
}

// removed db entries after 60 min
func RemovedExpiredEntries() {

	cur := dao.GetAll()

	for cur.Next(context.TODO()) {
		var input model.UrlModel
		if err := cur.Decode(&input); err != nil {
			log.Fatal("Error while decoding cursor value into model")
		}

		var start time.Time = input.Created_at
		a := time.Now().Sub(start)

		b := a.Minutes()

		if b > float64(EXPIRY_TIME) {
			dao.CleanDb(input.UniqueId)
		}
	}

}
