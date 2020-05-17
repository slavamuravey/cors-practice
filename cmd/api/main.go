package main

import (
	"log"
	"net/http"
	"slavamuravey/cors/pkg/handler"
)

func main() {
	http.HandleFunc("/api", handler.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}