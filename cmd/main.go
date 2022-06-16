package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	// server := delivery.New()
	fmt.Printf("Starting server at port :%v \n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
