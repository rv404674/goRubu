package services

//contains function to create a shortened url
import (
	"encoding/base64"
	dao "goRubu/daos"
	model "goRubu/models"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Race Condition - Undesirable condition where o/p of a program depends on the seq of
// execution of go routines

// To prevent this use Mutex - a locking mechanism, to ensure only one Go routine
// is running in the CS at any point of time

func CreateShortenedUrl(inputUrl string) string {
	var m sync.Mutex

	m.Lock()
	counterVal := dao.GetCounterValue()
	byteNumber := []byte(strconv.Itoa(counterVal))
	tempUrl := base64.StdEncoding.EncodeToString(byteNumber)

	new_url := "https://goRubu/" + tempUrl
	inputModel := model.UrlModel{UniqueId: counterVal, Url: inputUrl, Created_at: time.Now()}
	dao.InsertInShortenedUrl(inputModel)

	dao.UpdateCounter()
	m.Unlock()

	return new_url
}

//Use caching here.
// if
func UrlRedirection(inputUrl string) string {
	// https://goRubu/MTAwMDE=
	i := strings.Index(inputUrl, "Rubu/")
	encodedForm := inputUrl[i+5:]

	byteNumber, _ := base64.StdEncoding.DecodeString(encodedForm)
	UniqueId, _ := strconv.Atoi(string(byteNumber))

	url := dao.GetUrl(UniqueId)
	return url.Url

}
