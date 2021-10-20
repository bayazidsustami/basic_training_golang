package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type Article struct {
	Title    string
	URL      string
	Category string
}

func main() {
	res, err := http.Get("first-app.test")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error : %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([]Article, 0)

	doc.Find(".my-3").Children().Each(func(i int, selection *goquery.Selection) {
		row := new(Article)
		row.Title = selection.Find(".post-card-title").Text()
		row.URL, _ = selection.Find(".post-card-content-link").Attr("href")
		row.Category = selection.Find(".post-card-tags").Text()
		rows = append(rows, *row)
	})

	bts, err := json.MarshalIndent(rows, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(string(bts))
}
