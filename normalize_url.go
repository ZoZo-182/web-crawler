package main

import (
	"fmt"
	"log"
	"net/url"
)

func normalizeURL(inputURL string) (string, error) {
	u, err := url.Parse(inputURL)
	if err != nil {
		log.Fatalf("normalizeURL(): trouble parsing input url (%v)", err)
	}
	commonURL := fmt.Sprintf("%s%s", u.Host, u.Path)
	return commonURL, nil
}
