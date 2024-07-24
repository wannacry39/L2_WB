package main

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go-wget <url> <depth>")
		return
	}
	startURL := os.Args[1]
	maxDepth, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid depth value:", os.Args[2])
		return
	}

	visited := make(map[string]bool)
	err = downloadPage(startURL, visited, startURL, 0, maxDepth)
	if err != nil {
		fmt.Printf("Error downloading page: %v\n", err)
	}
}

func resolveURL(rawURL, baseURL string) string {
	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	ref, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return base.ResolveReference(ref).String()
}

func savePage(pageURL string, doc *goquery.Document) error {
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		return fmt.Errorf("invalid page URL: %v", err)
	}
	filePath := path.Join("downloaded", parsedURL.Path)
	if strings.HasSuffix(filePath, "/") {
		filePath = path.Join(filePath, "index.html")
	}
	os.MkdirAll(path.Dir(filePath), os.ModePerm)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	html, err := doc.Html()
	if err != nil {
		return fmt.Errorf("failed to generate HTML: %v", err)
	}

	_, err = file.WriteString(html)
	if err != nil {
		return fmt.Errorf("failed to save page: %v", err)
	}

	fmt.Printf("Page saved: %s\n", pageURL)
	return nil
}
