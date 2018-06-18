package main

import (
	"net/http"
	"./handlers"
)


func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/agg", handlers.NewsAggHandler)
	http.ListenAndServe(":8000", nil)
}
