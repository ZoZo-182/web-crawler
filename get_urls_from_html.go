package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
	"net/url"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	page, err := html.Parse(htmlReader)
	if err != nil {
		log.Fatal(err)
	}

	linksSlice := []string{}

	base, _ := url.Parse(rawBaseURL)
	if page.Data == "a" {
		for _, attr := range page.Attr {
			rel, _ := url.Parse(attr.Val)
			abs := base.ResolveReference(rel)
			linksSlice = append(linksSlice, abs.String())

		}
	}
	return linksSlice, nil
}
