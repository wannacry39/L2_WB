package main

import (
	"httpserver/internal/API"
	"log"
)

func main() {
	s := API.NewServer()
	s.SetupHandlers()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
