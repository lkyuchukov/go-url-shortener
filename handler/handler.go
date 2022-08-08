package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkyuchukov/go-url-shortener/shortener"
	"github.com/lkyuchukov/go-url-shortener/store"
)

type ShortenUrlRequest struct {
	LongUrl string `json:"longUrl" binding:"required"`
}

const host = "http://localhost:9000/"

func ShortenUrl(c *gin.Context) {
	var req ShortenUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.ShortenUrl(req.LongUrl)
	store.SaveShortUrl(shortUrl, req.LongUrl)

	c.JSON(200, gin.H{
		"shortUrl": host + shortUrl,
	})
}

func UrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	longUrl := store.RetrieveLongUrl(shortUrl)
	c.Redirect(http.StatusFound, longUrl)
}
