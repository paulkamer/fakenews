package main

import (
	"fakenews/feeds"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/news.rss", func(c *gin.Context) {
		c.XML(http.StatusOK, feeds.GenerateRssFeed())
	})

	r.GET("/news.atom", func(c *gin.Context) {
		c.XML(http.StatusOK, feeds.GenerateAtomFeed())
	})

	semiInvalid := r.Group("/semi-invalid")
	{
		semiInvalid.GET("/rss/as-atom", func(c *gin.Context) {
			c.Header("Content-Type", "application/atom+xml")
			c.XML(http.StatusOK, feeds.GenerateRssFeed())
		})

		semiInvalid.GET("/rss/as-html", func(c *gin.Context) {
			c.Header("Content-Type", "text/html")
			c.XML(http.StatusOK, feeds.GenerateRssFeed())
		})

		semiInvalid.GET("/atom/as-rss", func(c *gin.Context) {
			c.Header("Content-Type", "application/rss+xml")
			c.XML(http.StatusOK, feeds.GenerateAtomFeed())
		})
	}

	invalid := r.Group("/invalid")
	{
		invalid.GET("/rss/returns-html", func(c *gin.Context) {
			c.Header("Content-Type", "application/rss+xml")

			c.HTML(http.StatusOK, "index.tmpl", gin.H{})
		})

		invalid.GET("/rss/invalid-syntax", func(c *gin.Context) {
			content, err := os.ReadFile("templates/invalid_rss.xml")
			if err != nil {
				c.String(http.StatusInternalServerError, "Error reading file")
				return
			}

			c.Header("Content-Type", "application/rss+xml")
			c.String(http.StatusOK, string(content))
		})
	}

	r.Run()
}
