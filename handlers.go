package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//RedirectHandler handles all redirects
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trimHash := vars["hash"]
	fmt.Fprintf(w, "Redirecting now from %s\n", trimHash)
}
