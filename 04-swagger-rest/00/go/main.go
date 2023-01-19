package main

import (
	"log"
	"net/http"
)

var candies map[string]Candy

func main() {
	log.Printf("Server started")

	candies = fillCandies()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
