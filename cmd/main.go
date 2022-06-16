package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupie-tracker/internal/delivery"
)

func main() {
	port := os.Getenv("PORT")

	server := delivery.New()
	fmt.Printf("Starting server at port :%v \n", port)
	log.Fatal(http.ListenAndServe(port, server.Router()))
}
