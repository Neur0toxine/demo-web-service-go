package model

import "time"

type Book struct {
	ID              uint64    `json:"id"`
	Name            string    `json:"name,omitempty"`
	PublicationDate time.Time `json:"publicationDate"`
	Rate            uint8     `json:"rate"`
}
