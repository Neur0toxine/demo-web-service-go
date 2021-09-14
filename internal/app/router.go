package app

import (
	"github.com/gin-gonic/gin"
)

// Router controls the routes
func Router() (r *gin.Engine) {
	r = gin.New()
	r.GET("/authors", AuthorsHandler)
	return
}
