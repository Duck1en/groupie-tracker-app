package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"groupie-tracker/internal/delivery"
)

func main() {
	port := os.Getenv("PORT")

	server := delivery.New()

	if port == "" {
		fmt.Print("Enter port :")
		ch := make(chan int)

		go func() {
			fmt.Fscan(os.Stdin, &port)
			ch <- 1
		}()

		select {
		case <-ch:
			if _, err := strconv.Atoi(port); err != nil || port == "" {
				log.Fatal("PORT is NULL or string")
			}

			fmt.Printf("Starting server at port :%v \n", port)
			log.Fatal(http.ListenAndServe(":"+port, server.Router()))
			os.Exit(0)

		case <-time.After(5 * time.Second):
			log.Fatal("Time out")
		}

	} else {
		fmt.Printf("Starting server at port :%v \n", port)
		log.Fatal(http.ListenAndServe(":"+port, server.Router()))
		os.Exit(0)
	}
}
