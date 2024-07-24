package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func downloadPage(pageURL string, visited map[string]bool, baseURL string, currentDepth, maxDepth int) error {
	if currentDepth > maxDepth {
		return nil
	}
	if visited[pageURL] {
		return nil
	}
	visited[pageURL] = true

	resp, err := http.Get(pageURL)
	if err != nil {
		return fmt.Errorf("failed to GET page: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("invalid base URL: %v", err)
	}
	page, err := url.Parse(pageURL)
	if err != nil {
		return fmt.Errorf("invalid page URL: %v", err)
	}
	pageURL = base.ResolveReference(page).String()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %v", err)
	}

	doc.Find("a, img, link, script").Each(func(i int, s *goquery.Selection) {
		if href, exists := s.Attr("href"); exists {
			absURL := resolveURL(href, baseURL)
			if absURL != "" {
				downloadPage(absURL, visited, baseURL, currentDepth+1, maxDepth)
			}
		}
		if src, exists := s.Attr("src"); exists {
			absURL := resolveURL(src, baseURL)
			if absURL != "" {
				downloadResource(absURL, baseURL)
			}
		}
	})

	return savePage(pageURL, doc)
}
