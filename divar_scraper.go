package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gen2brain/beeep"
)

func main() {
	BASE_URL := "https://divar.ir/s/iran/laptop-notebook-macbook/apple?goods-business-type=all&price=39000000-40000000&cities=1%2C8"
	targetTitle := "مک"

	for {
		response, err := http.Get(BASE_URL)

		// nil == undefined
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			log.Fatalf("Req F: %d", response.StatusCode)
		}

		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".kt-post-card").Each(func(i int, s *goquery.Selection) {
			title := s.Find(".kt-post-card__title").Text()
			description := s.Find(".kt-post-card__description").Text()
			imageURL, _ := s.Find(".kt-post-card__image").Attr("src")

			if strings.Contains(title, targetTitle) {
				err := beeep.Alert("پیدا شد!", "Title: "+title+"\nDescription: "+description+"\nImage URL: "+imageURL, "")
				if err != nil {
					log.Fatal(err)
				}
			}
		})

		time.Sleep(5 * time.Minute)
	}
}
