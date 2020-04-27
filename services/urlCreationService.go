package services

//contains function to create a shortened url
import (
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

func init() {
	if err = godotenv.Load("variables.env"); err != nil {
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

	urlModel := dao.GetUrl(UniqueId)
	return urlModel.Url
}
