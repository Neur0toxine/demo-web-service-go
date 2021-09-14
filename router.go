package main

import "github.com/gin-gonic/gin"

func router() (r *gin.Engine) {
	r = gin.New()
	r.GET("/authors", AuthorsHandler)
	return
}
