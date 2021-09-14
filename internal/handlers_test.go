package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Neur0toxine/demo-web-service-go/data/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthorsHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/authors", nil)
	require.NoError(t, err)

	r := Router()
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	data, err := io.ReadAll(rec.Body)
	require.NoError(t, err)

	var authors []model.Author
	require.NoError(t, json.Unmarshal(data, &authors))

	assert.Len(t, authors, 10)

	for _, author := range authors {
		assert.NotEmpty(t, author.ID)
		assert.NotEmpty(t, author.FirstName)
		assert.NotEmpty(t, author.LastName)
		assert.Len(t, author.Books, 10)

		for _, book := range author.Books {
			assert.NotEmpty(t, book.ID)
			assert.NotEmpty(t, book.Name)
			assert.NotEmpty(t, book.PublicationDate)
		}
	}
}
