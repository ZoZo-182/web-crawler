package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
	"net/url"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Errorf("getURLSFromHTML(): failed to parse rawBaseURL - %v", err)
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Errorf("getURLSFromHTML(): failed to parse html reader- %v", err)
	}

	var linksSlice []string

	var visitNodes func(*html.Node)
	visitNodes = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					href, _ := url.Parse(a.Val)
					abs := baseURL.ResolveReference(href)
					linksSlice = append(linksSlice, abs.String())
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			visitNodes(child)
		}
	}
	visitNodes(doc)
	return linksSlice, nil
}
