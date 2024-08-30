package feeds

import (
	"encoding/xml"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Atom struct {
	XMLName xml.Name `xml:"feed"`
	Xmlns   string   `xml:"xmlns,attr"`
	Title   string   `xml:"title"`
	Link    Link     `xml:"link"`
	Id      string   `xml:"id"`
	Updated string   `xml:"updated"`
	Author  Author   `xml:"author"`
	Entry   []Entry  `xml:"entry"`
}

type Link struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}

type Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type Entry struct {
	Title   string `xml:"title"`
	Link    Link   `xml:"link"`
	Id      string `xml:"id"`
	Updated string `xml:"updated"`
	Summary string `xml:"summary"`
}

func GenerateAtomFeed() Atom {
	numberOfEntries := gofakeit.Number(10, 25)

	entries := []Entry{}
	for i := 0; i < numberOfEntries; i++ {
		randomTitleLength := gofakeit.Number(5, 10)
		randomSummaryLength := gofakeit.Number(3, 10)

		entry := Entry{
			Title:   gofakeit.Sentence(randomTitleLength),
			Link:    Link{Href: gofakeit.URL(), Rel: "alternate"},
			Id:      gofakeit.UUID(),
			Updated: gofakeit.Date().Format(time.RFC3339),
			Summary: gofakeit.Paragraph(randomSummaryLength, 5, 10, " "),
		}
		entries = append(entries, entry)
	}

	atom := Atom{
		Xmlns:   "http://www.w3.org/2005/Atom",
		Title:   gofakeit.Sentence(8),
		Link:    Link{Href: gofakeit.URL(), Rel: "self"},
		Id:      gofakeit.UUID(),
		Updated: gofakeit.Date().Format(time.RFC3339),
		Entry:   entries,
	}

	return atom
}
