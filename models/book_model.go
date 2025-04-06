package models

import (
	"time"
	"victorubere/library/lib/structs"
)

type Book struct {
	structs.Model
	Title             string    `json:"title"`
	Author            string    `json:"author"`
	ISBN              string    `json:"isbn"`
	Publisher         string    `json:"publisher"`
	PublicationDate   time.Time `json:"publication_date"`
	Pages             int       `json:"pages"`
	Description       string    `json:"description"`
	QuantityAvailable int       `json:"quantity"`
}
