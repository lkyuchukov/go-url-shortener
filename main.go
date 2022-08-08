package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lkyuchukov/go-url-shortener/handler"
	"github.com/lkyuchukov/go-url-shortener/store"
)

func main() {
	r := gin.Default()

	r.POST("/shorten", func(c *gin.Context) {
		handler.ShortenUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.UrlRedirect(c)
	})

	store.Init()

	err := r.Run(":9000")
	if err != nil {
		log.Fatalf("Failed to start the web server - Error: %v", err)
	}
}
