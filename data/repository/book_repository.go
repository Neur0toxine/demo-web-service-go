package repository

import (
	"math/rand"
	"time"

	"github.com/Neur0toxine/demo-web-service-go/data/model"
	"github.com/brianvoe/gofakeit/v6"
)

func Books() (books []model.Book) {
	books = make([]model.Book, 10)
	for i := 0; i < 10; i++ {
		books[i] = model.Book{
			ID:              uint64(i + 1),
			Name:            gofakeit.Phrase(),
			PublicationDate: time.Now(),
			Rate:            uint8(rand.Intn(5)),
		}
	}
	return
}
