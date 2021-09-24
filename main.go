package main

import (
	"log"
	"net/http"

	"github.com/rbixby/journal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
