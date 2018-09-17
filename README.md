# grip
Web scraping utility in Go

```golang
func main() {
    
    scraper := grip.Scraper{}

    tag := grip.Tag{Type: "a", Attributes: [{Name: "href"}], ScrapeTextContent: false}

    scraper.LookFor(tag)

    scraper.Scrape("http://google.com")

    // or scraper.ScrapeAll([]string{"http://google.com","https://facebook.com"})

    for _, site := range scraper.Results {
        fmt.Printf("Links from '%s'\n", site)
        for key, value := range site.Results() {
            fmt.Printf("    Tag: %s, Attribute Value: %s\n", key, value.Attributes["href"])
        }
    }
}
```

## Tags and Attributes

Tags define what the scraper should look for, including whether it should include the text content of tag.

```golang
type Tag struct {
    Type string
    Location string
    Attributes []Attribute
    ScrapeTextContent bool
    StartDepth uint32
    ExploreDepth uint32
}

type Attribute struct {
    Name string
    Value string
}
```

Tags include a list of attributes. If only the attribute name is set, then all tags with that attribute will be retrieved. If a value is given to the attribute, then only tags that have a matching attribute name/value pair will be retrieved.

## Results

```golang
type ScrapeSite struct {
    Url string
    Results map[string]ScrapeResult
}


type ScrapeResult struct {
    Tag string
    Attributes map[string]string
}
```