package app

import (
	"net/http"

	"github.com/Neur0toxine/demo-web-service-go/internal/data/repository"
	"github.com/gin-gonic/gin"
)

// AuthorsHandler will return formatted JSOn response with authors
func AuthorsHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repository.Authors())
}
