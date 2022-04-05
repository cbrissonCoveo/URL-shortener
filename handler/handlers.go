package handler

import (
	"net/http"

	"github.com/cbrissonCoveo/URL-shortener/shortener"
	"github.com/cbrissonCoveo/URL-shortener/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationReq struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}
func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationReq
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9998/"
	c.JSON(200, gin.H{
		"message" : "short url created successfully",
		"short_url": host +shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}