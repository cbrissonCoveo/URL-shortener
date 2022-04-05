package main

import (
	"github.com/cbrissonCoveo/URL-shortener/handler"
	"github.com/cbrissonCoveo/URL-shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r  := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener API!",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()
	err := r.Run(":9998")
	if err != nil {
		panic(err)
	}
}