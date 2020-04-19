package models

import "time"

// so basically there are two collections
// one stores the url info, and second a counter

type UrlModel struct {
	UniqueId   int       `json: "uniqueid"`
	Url        string    `json: "url"`
	Created_at time.Time `json: "created_at"`
}

// single row table
// uniqueid value
// counter 1000
type IncrementerModel struct {
	UniqueId string `json: "uniqueid"`
	Value    int    `json: "value"`
}
