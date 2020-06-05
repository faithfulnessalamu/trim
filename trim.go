package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{hash}", RedirectHandlerFunc).Methods(http.MethodGet)
	// Serve the home page
	r.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", r)
}
