package handlers

import (
	"net/http"
	"fmt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>First Golang Site..<h1>")
	fmt.Fprintf(w, "<h2>Washington Post News Aggregator<h3>")
}
