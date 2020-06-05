package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//RedirectHandler handles all redirects
func RedirectHandlerFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trimHash := vars["hash"]
	fmt.Fprintf(w, "Redirecting now from %s\n", trimHash)
}
