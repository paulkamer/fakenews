package main

import (
	"fakenews/feeds"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/news.rss", func(c *gin.Context) {
		c.XML(http.StatusOK, feeds.GenerateRssFeed())
	})

	r.GET("/news.atom", func(c *gin.Context) {
		c.XML(http.StatusOK, feeds.GenerateAtomFeed())
	})

	// 'semi-invalid' responses
	r.GET("/semi-invalid/rss/as-atom", func(c *gin.Context) {
		c.Header("Content-Type", "application/atom+xml")
		c.XML(http.StatusOK, feeds.GenerateRssFeed())
	})

	r.GET("/semi-invalid/rss/as-html", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.XML(http.StatusOK, feeds.GenerateRssFeed())
	})

	r.GET("/semi-invalid/atom/as-rss", func(c *gin.Context) {
		c.Header("Content-Type", "application/rss+xml")
		c.XML(http.StatusOK, feeds.GenerateAtomFeed())
	})

	r.Run()
}
