package main

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main()	{ getLatestUserAgent() }

func getLatestUserAgent() string {
	requestURL := "https://www.whatismybrowser.com/guides/the-latest-user-agent/chrome"
	httpResponse, err := http.Get(requestURL)
	parsedHTML, err := goquery.NewDocumentFromReader(httpResponse.Body)
	if err != nil {
		fmt.Println(err)
	}

	latestUserAgent := parsedHTML.Find(".code").First().Text()
        fmt.Println(latestUserAgent)
	return latestUserAgent
}
