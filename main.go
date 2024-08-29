package main

import (
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gofakeit.Seed(time.Now().UnixNano())

	r.GET("/news.rss", func(c *gin.Context) {
		rss := generateRssFeed()

		c.XML(http.StatusOK, rss)
	})

	r.Run()
}
