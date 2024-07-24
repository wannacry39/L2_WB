package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func downloadResource(resourceURL, baseURL string) error {
	resp, err := http.Get(resourceURL)
	if err != nil {
		return fmt.Errorf("failed to GET resource: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("invalid base URL: %v", err)
	}
	resource, err := url.Parse(resourceURL)
	if err != nil {
		return fmt.Errorf("invalid resource URL: %v", err)
	}
	resourceURL = base.ResolveReference(resource).String()

	filePath := path.Join("downloaded", resource.Path)
	os.MkdirAll(path.Dir(filePath), os.ModePerm)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save resource: %v", err)
	}

	fmt.Printf("Resource saved: %s\n", resourceURL)
	return nil
}
