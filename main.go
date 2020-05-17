package main

import (
	"fmt"
	"net/http"

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
	longURL := query.Get("longUrl")

	trimmedURL, err := trim.GetTrimmed(longURL)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, trimmedURL)
}
