package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"groupie-tracker/internal/delivery"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		fmt.Print("Enter port :")
		ch := make(chan int)

		go func() {
			fmt.Fscan(os.Stdin, &port)
			ch <- 1
		}()

		select {
		case <-ch:
			server := delivery.New()
			fmt.Printf("Starting server at port :%v \n", port)
			log.Fatal(http.ListenAndServe(":"+port, server.Router()))
			os.Exit(0)
		case <-time.After(3 * time.Second):
			log.Fatal("Time out")
		}
	}
}
