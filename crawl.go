package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func crawl(filename string) {

	fName := filename
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Write CSV header
	writer.Write([]string{"Text", "URL"})
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML(".athing", func(e *colly.HTMLElement) {

		writer.Write([]string{
			e.ChildText("a"),
			e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
		})

	})

	c.Visit("https://news.ycombinator.com/")

}



