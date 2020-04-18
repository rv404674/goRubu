package models

import "time"

// so basically there are two collections
// one stores the url info, and second a counter

type UrlModel struct {
	Id         int
	Url        string
	Created_at time.Time
}

type IncrementerMode struct {
	Id      string
	Counter int
}
