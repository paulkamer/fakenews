package main

import (
	"encoding/xml"
	"fakenews/feeds"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./templates", ".tmpl")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	valid := app.Group("/valid")
	{
		valid.Get("/rss", func(c *fiber.Ctx) error {
			return c.XML(feeds.GenerateRssFeed())
		})

		valid.Get("/atom", func(c *fiber.Ctx) error {
			return c.XML(feeds.GenerateAtomFeed())
		})
	}

	semiInvalid := app.Group("/semi-invalid")
	{
		semiInvalid.Get("/rss/as-atom", func(c *fiber.Ctx) error {
			c.Set("Content-Type", "application/atom+xml")

			rssFeed := feeds.GenerateRssFeed()
			rssFeedBytes, err := xml.Marshal(rssFeed)
			if err != nil {
				return c.Status(http.StatusInternalServerError).SendString("Failed to marshal RSS feed")
			}
			return c.Send(rssFeedBytes)
		})

		semiInvalid.Get("/rss/as-html", func(c *fiber.Ctx) error {
			c.Set("Content-Type", "text/html")

			rssFeed := feeds.GenerateRssFeed()
			rssFeedBytes, err := xml.Marshal(rssFeed)
			if err != nil {
				return c.Status(http.StatusInternalServerError).SendString("Failed to marshal RSS feed")
			}
			return c.Send(rssFeedBytes)
		})

		semiInvalid.Get("/atom/as-rss", func(c *fiber.Ctx) error {
			c.Set("Content-Type", "application/rss+xml")
			atomFeed := feeds.GenerateAtomFeed()
			atomFeedBytes, err := xml.Marshal(atomFeed)
			if err != nil {
				return c.Status(http.StatusInternalServerError).SendString("Failed to marshal Atom feed")
			}
			return c.Send(atomFeedBytes)
		})
	}

	invalid := app.Group("/invalid")
	{
		invalid.Get("/rss/returns-html", func(c *fiber.Ctx) error {
			c.Set("Content-Type", "application/rss+xml")

			return c.Render("index", fiber.Map{})
		})

		invalid.Get("/rss/invalid-syntax", func(c *fiber.Ctx) error {
			content, err := os.ReadFile("templates/invalid_rss.xml")
			if err != nil {
				return c.Status(http.StatusInternalServerError).SendString("Error reading file")
			}

			c.Set("Content-Type", "application/rss+xml")
			return c.SendString(string(content))
		})
	}

	redirects := app.Group("/redirects")
	{
		redirects.Get("/rss/valid", func(c *fiber.Ctx) error {
			time.Sleep(1 * time.Second)

			return c.Redirect("/valid/rss", http.StatusTemporaryRedirect)
		})

		redirects.Get("/multiple", func(c *fiber.Ctx) error {
			time.Sleep(1 * time.Second)

			return c.Redirect("/redirects/rss/valid", http.StatusTemporaryRedirect)
		})

		redirects.Get("/https/to/http", func(c *fiber.Ctx) error {
			time.Sleep(1 * time.Second)

			return c.Redirect("http://127.0.0.1:8080/valid/rss", http.StatusTemporaryRedirect)
		})

		redirects.Get("/http/to/https", func(c *fiber.Ctx) error {
			time.Sleep(1 * time.Second)

			return c.Redirect("https://127.0.0.1:8443/valid/rss", http.StatusTemporaryRedirect)
		})
	}

	// // Random endpoint
	// This endpoint will randomly select one of the defined endpoints and exectutes it (without redirecting)
	app.Get("/random", func(c *fiber.Ctx) error {
		var endpoints []string
		for _, routes := range app.Stack() {
			for _, route := range routes {
				endpoints = append(endpoints, route.Path)
			}
		}

		randomEndpoint := endpoints[time.Now().UnixNano()%int64(len(endpoints))]
		c.Path(randomEndpoint)
		return c.Next()
	})

	go func() {
		err_http := app.Listen(":8080")
		if err_http != nil {
			log.Fatal("Web server (HTTP): ", err_http)
		}
	}()

	err := app.ListenTLS(":8443", "cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
