package grip

import (
	"net/http"

	"golang.org/x/net/html"
)

// Tag ...
type Tag struct {
	Type              string
	Attributes        []Attribute
	ScrapeTextContent bool
}

// Attribute ...
type Attribute struct {
	Name  string
	Value string
}

// Scraper ...
type Scraper struct {
	tags    []Tag
	Results []ScrapeResult
}

// NewScraper ...
func NewScraper() *Scraper {
	return &Scraper{tags: make([]Tag, 1)}
}

// LookForAll ...
func (s *Scraper) LookForAll(tags []Tag) {
	for _, t := range tags {
		s.LookFor(t)
	}
}

// LookFor ...
func (s *Scraper) LookFor(t Tag) {
	s.tags = append(s.tags, t)
}

// Scrape ...
func (s *Scraper) Scrape(url string) (results []ScrapeResult, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

	results = make([]ScrapeResult, 0)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			for _, tag := range s.tags {
				if t.Data == tag.Type {
					// Verify all the attributes match
					attrs := verifyAttributeMatches(&tag.Attributes, &t.Attr)

					// If everything matched, add a new result
					if attrs != nil {
						text := ""
						if tag.ScrapeTextContent {
							text = string(z.Text())
						}
						results = append(results, ScrapeResult{Tag: tag.Type, Attributes: attrs, Text: text})
					}
				}
			}
		}
	}
}

func verifyAttributeMatches(attr *[]Attribute, tattr *[]html.Attribute) map[string]string {
	result := make(map[string]string)
	for _, ta := range *tattr {
		for _, a := range *attr {
			if ta.Key == a.Name {
				match := a.Value == ta.Val
				if a.Value == "" || match {
					result[ta.Key] = ta.Val
				} else {
					return nil
				}
			}
		}
	}
	return result
}
