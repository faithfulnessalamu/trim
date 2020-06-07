package main

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type trimRequest struct {
	URL string `json:"url"`
}

//payload is used to return a json payload
type payload struct {
	Msg string `json:"msg"`
}

// Create json from string
func payloadify(s string) *payload {
	p := new(payload)
	p.Msg = s
	return p
}

func writePayload(status int, w http.ResponseWriter, p *payload) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

//TrimHandler returns a function that handles all trims
//A closure is used so database can be mocked for tests
//Matches /trim
func TrimHandler(db database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse json body
		decoder := json.NewDecoder(r.Body)
		var tr trimRequest
		dErr := decoder.Decode(&tr)
		if dErr != nil {
			// Couldn't get post request body
			writePayload(http.StatusBadRequest, w, payloadify("Invalid POST body"))
			infoLogger.Print("Invalid POST body")
			return
		}
		actualURL := tr.URL
		infoLogger.Printf("Trim request received for %s", actualURL)
		// Check if is valid url
		if !isValidURL(actualURL) {
			writePayload(http.StatusBadRequest, w, payloadify("Not a valid URL"))
			infoLogger.Printf("Not a valid URL: %s", actualURL)
			return
		}
		// Get hash for actualURL
		trimHash := hash(actualURL)
		// Try to save
		err := db.save(trimHash, actualURL)
		if err != nil {
			infoLogger.Printf("%s already exists in database", actualURL)
		}
		// Create the return url
		returnURL := filepath.Join(baseURL, trimHash)
		writePayload(http.StatusOK, w, payloadify(returnURL))
		infoLogger.Printf("Trim for %s is %s", actualURL, returnURL)
	}
}

//RedirectHandler returns a function that handles all redirects
//A closure is used here so database can be mocked for tests
func RedirectHandler(db database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		trimHash := vars["hash"]
		infoLogger.Printf("Redirect request received for %s", trimHash)
		// Try to get the actual URL for the hash
		actualURL, err := db.retrieve(trimHash)
		if err != nil {
			infoLogger.Printf("%s key not in database", trimHash)
			http.NotFound(w, r)
			return
		}
		infoLogger.Printf("Redirecting to %s", actualURL)
		http.Redirect(w, r, actualURL, http.StatusTemporaryRedirect)
	}
}
