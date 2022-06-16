package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/internal/delivery"
)

type Artist struct {
	id           int
	image        string
	name         string
	members      []string
	creationDate int
	firstAlbum   string
	location     string
	concertDates string
	relations    string
}

func main() {
	server := delivery.New()
	fmt.Printf("Starting server at port :8080\nhttp://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", server.Router()))
}
