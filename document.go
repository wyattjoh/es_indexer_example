package esindexer

import "time"

type Timestamp struct {
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type Document struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Timestamp Timestamp `json:"timestamp"`
}
