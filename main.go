package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()

    c.OnHTML("table.wikitable > tbody", func(e *colly.HTMLElement) {
        e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.ChildText("td:nth-child(1)"))
		})
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })

    c.Visit("https://en.wikipedia.org/wiki/National_Basketball_Association#Teams")
}