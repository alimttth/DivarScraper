package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gen2brain/beeep"
)

type Interface struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	// config filters
	BASE_URL := "https://divar.ir/s/iran/laptop-notebook-macbook/apple?goods-business-type=all&price=39000000-40000000&cities=1%2C8"
	TARGET_TITLE := "مک"

	for {
		response, err := http.Get(BASE_URL)
		// nil == undefined
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()
		// Promise == defer

		if response.StatusCode != 200 {
			log.Fatalf("Req F: %d", response.StatusCode)
		}

		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".kt-post-card").Each(func(i int, post *goquery.Selection) {
			title := post.Find(".kt-post-card__title").Text()
			description := post.Find(".kt-post-card__description").Text()
			if strings.Contains(title, TARGET_TITLE) {
				err := beeep.Alert("پیدا شد!", "Title: "+title+"\nDescription: "+description+"\nImage URL: ", "")
				if err != nil {
					log.Fatal(err)
				}
			}
		})

		time.Sleep(5 * time.Minute)
	}
}
