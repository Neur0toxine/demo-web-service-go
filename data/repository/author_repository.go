package repository

import (
	"github.com/Neur0toxine/demo-web-service-go/data/model"
	"github.com/brianvoe/gofakeit/v6"
)

func Authors() (authors []model.Author) {
	authors = make([]model.Author, 10)
	for i := 0; i < 10; i++ {
		authors[i] = model.Author{
			ID:        uint64(i + 1),
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Books:     Books(),
		}
	}
	return
}
