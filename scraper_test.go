package grip

import (
	"fmt"
	"testing"
)

func TestLookFor(t *testing.T) {

	tag := Tag{Type: "a", Attributes: []Attribute{{Name: "href"}}, ScrapeTextContent: false}

	scraper := NewScraper()

	scraper.LookFor(tag)

}

func TestScrape(t *testing.T) {

	tag := Tag{Type: "a", Attributes: []Attribute{{Name: "href"}}, ScrapeTextContent: true}

	scraper := NewScraper()

	scraper.LookFor(tag)

	results, err := scraper.Scrape("http://google.com")
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, sr := range results {
		fmt.Printf("Found <%s> for \"%s\" with content \"%s\"\n", sr.Tag, sr.Attributes["href"], sr.Text)
	}

}
