package main

import (
	"log"
	"net/http"

	"github.com/rbixby/journal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/edit/", handlers.EditHandler)
	http.HandleFunc("/save/", handlers.SaveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
