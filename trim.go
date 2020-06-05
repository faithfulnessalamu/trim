package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{hash}", RedirectHandlerFunc).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
