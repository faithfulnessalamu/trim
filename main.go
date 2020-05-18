package main

import (
	"fmt"
	"net/http"

	"github.com/vague369/trim/data"
	"github.com/vague369/trim/trim"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/trim", trimHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Trim")
}

func trimHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	input := query.Get("input")

	trimmedURL, err := trim.GetTrimmed(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
	}
	// input is validated to be a URL
	longURL := input
	// Save url pair to database
	saveErr := data.SavePair(trimmedURL, longURL)
	fmt.Fprintln(w, trimmedURL)
}
