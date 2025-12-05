package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	page := 1
	resp, err := http.Get(fmt.Sprintf("https://zukan.inazuma.jp/en/chara_param/?page=%d&per_page=1000", page))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("test")
	doc.Find("ul.charaListBox li").Each(func(i int, s *goquery.Selection) {
		name := s.Find("div.nameBox span.name ruby").First().Text()
		if name == "" {
			return
		}
		fmt.Printf("Player Name: %s", name)
	})
}
