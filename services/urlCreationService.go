package services

//contains function to create a shortened url
import (
	dao "goRubu/daos"
	model "goRubu/models"
	"time"
)

func CreateShortenedUrl(inputUrl string) {
	urlModel := model.UrlModel{Id: 1000, Url: inputUrl, Created_at: time.Now()}
	dao.InsertInShortenedUrl(urlModel)
}
