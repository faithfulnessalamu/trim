package data

import "errors"

var errorKeyExists = errors.New("Key Exists")
var storage = make(map[string]string)

/*SavePair persists the url pairs for retrieval later*/
func SavePair(trimmedURL, longURL string) error {
	_, exists := storage[trimmedURL]
	if exists {
		return errorKeyExists
	}
	storage[trimmedURL] = longURL
	return nil
}
