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
	// Serve static files
	statics := r.PathPrefix("/static/")
	statics.Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":8080", r)
}
