package main

import (
	"net/http"

	"github.com/rbixby/journal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
}
