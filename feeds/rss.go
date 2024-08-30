package feeds

import (
	"encoding/xml"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Guid        string    `xml:"guid"`
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	PubDate     time.Time `xml:"pubDate"`
}

func GenerateRssFeed() RSS {
	numberOfItems := gofakeit.Number(10, 25)

	items := []Item{}
	for i := 0; i < numberOfItems; i++ {
		randomTitleLength := gofakeit.Number(5, 10)
		randomParagraphLength := gofakeit.Number(3, 10)

		item := Item{
			Guid:        gofakeit.UUID(),
			Title:       gofakeit.Sentence(randomTitleLength),
			Link:        gofakeit.URL(),
			Description: gofakeit.Paragraph(randomParagraphLength, 5, 10, " "),
			PubDate:     gofakeit.Date(),
		}
		items = append(items, item)
	}

	rss := RSS{
		Version: "2.0",
		Channel: Channel{
			Title:       gofakeit.Sentence(8),
			Link:        gofakeit.URL(),
			Description: gofakeit.Paragraph(3, 5, 10, " "),
			Items:       items,
		},
	}
	return rss
}
