package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/vague369/trim/data"
	"github.com/vague369/trim/trimmer"
)

const rootPath = "localhost:8080"

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/trim", trimHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fullURL := rootPath + r.URL.String()
	if isValidTrimURL(fullURL) {
		trimmedURL := fullURL
		// Get longURL and redirect
		longURL, err := data.GetLongURL(trimmedURL)
		if err != nil {
			// This trimmed URL is not in the database
			http.NotFound(w, r)
		} else {
			// Redirect the user temporarily
			http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
		}
	} else {
		// Not a valid URL
		http.NotFound(w, r)
	}
}

/* Checks if a URL is a trim url */
func isValidTrimURL(url string) bool {
	matcher := regexp.MustCompile(fmt.Sprintf("^%s/[a-zA-Z0-9]{8,8}", rootPath))
	return matcher.MatchString(url)
}

func trimHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	input := query.Get("input")

	trimmedURL, err := trimmer.GetTrimmed(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
	}
	// input is validated to be a URL
	longURL := input
	// Save url pair to database
	data.SavePair(trimmedURL, longURL)

	fmt.Fprintln(w, trimmedURL)
}
