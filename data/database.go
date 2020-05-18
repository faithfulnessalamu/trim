package data

import "errors"

var errorKeyExists = errors.New("Key exists")
var errorInvalidKey = errors.New("Key does not exist")

var storage = make(map[string]string)

/*SavePair persists the url pairs for retrieval later
Returns error if key exists in storage already*/
func SavePair(trimmedURL, longURL string) error {
	_, exists := storage[trimmedURL]
	if exists {
		return errorKeyExists
	}
	storage[trimmedURL] = longURL
	return nil
}

/*GetLongURL returns the long url of a trimmed url if it exists in storage*/
func GetLongURL(trimmedURL string) (string, error) {
	longURL, exists := storage[trimmedURL]
	if !exists {
		return "", errorInvalidKey
	}
	return longURL, nil
}
