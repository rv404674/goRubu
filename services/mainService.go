package services

//contains function to create a shortened url
import (
	"context"
	"encoding/base64"
	dao "goRubu/daos"
	model "goRubu/models"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rainycape/memcache"
)

// Race Condition - Undesirable condition where o/p of a program depends on the seq of
// execution of go routines

// To prevent this use Mutex - a locking mechanism, to ensure only one Go routine
// is running in the CS at any point of time

var mc *memcache.Client
var CACHE_EXPIRATION int64
var err error
var EXPIRY_TIME int

func init() {
	dir, _ := os.Getwd()
	envFile := "variables.env"
	if strings.Contains(dir, "test") {
		envFile = "../variables.env"
	}

	log.Println("Working dir", dir)

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Unable to load env file from urlCreationService Init", err)
	}

	// in seconds
	CACHE_EXPIRATION, _ = strconv.ParseInt(os.Getenv("CACHE_EXPIRATION"), 0, 64)
	mc, err = memcache.New(os.Getenv("MEMCACHED_DOMAIN"))
	if err != nil {
		log.Fatal("Unable to establish connection with the cache")
	} else {
		log.Println("Connection to Memcached Established")
	}

	// in min
	EXPIRY_TIME, _ = strconv.Atoi(os.Getenv("EXPIRY_TIME"))
}

func CreateShortenedUrl(inputUrl string) string {

	// TODO handle concurrency

	counterVal := dao.GetCounterValue()
	byteNumber := []byte(strconv.Itoa(counterVal))
	tempUrl := base64.StdEncoding.EncodeToString(byteNumber)

	new_url := "https://goRubu/" + tempUrl
	inputModel := model.UrlModel{UniqueId: counterVal, Url: inputUrl, Created_at: time.Now()}

	//first update the cache with (key,val) => (new_url, inputUrl)
	err = mc.Set(&memcache.Item{
		Key:        new_url,
		Value:      []byte(inputUrl),
		Expiration: int32(CACHE_EXPIRATION),
	})

	if err != nil {
		log.Fatal("Error in setting memcached value ", err)
	}
	dao.InsertInShortenedUrl(inputModel)
	dao.UpdateCounter()
	return new_url
}

//Use caching here.
func UrlRedirection(inputUrl string) string {
	// https://goRubu/MTAwMDE=
	i := strings.Index(inputUrl, "Rubu/")
	encodedForm := inputUrl[i+5:]

	byteNumber, _ := base64.StdEncoding.DecodeString(encodedForm)
	UniqueId, _ := strconv.Atoi(string(byteNumber))

	// try hitting the cache first
	// stored as "https://goRubu/MTW" -> "www.google.com"
	url, err := mc.Get(inputUrl)
	if err == nil {
		log.Println("Shortened url found in cache", string(url.Value))
		return string(url.Value)
	} else if err != memcache.ErrCacheMiss {
		log.Fatal("Memcached error ", err)
	}

	// if its a cache miss, fetch the value from db and update the cache.
	urlModel := dao.GetUrl(UniqueId)

	err2 := mc.Set(&memcache.Item{
		Key:        inputUrl,
		Value:      []byte(urlModel.Url),
		Expiration: int32(CACHE_EXPIRATION),
	})

	if err2 != nil {
		log.Fatal("Error in writing Memcached Value ", err2)
	}

	// urlMode.Url will be "", if the given shortened url does't exists in db.
	return urlModel.Url
}

// removed db entries after 5 min
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