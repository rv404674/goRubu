package models

import "time"

// so basically there are two collections
// one stores the url info, and second a counter

//UrlModel - model for collection containing (counter, orig_url, created_at) fields
type UrlModel struct {
	UniqueId  int       `json: "uniqueid"`
	Url       string    `json: "url"`
	CreatedAt time.Time `json: "created_at"`
}

// IncrementerModel - single row table
// value 10000
type IncrementerModel struct {
	Value int `json: "value"`
}
