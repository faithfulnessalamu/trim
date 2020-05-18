package data

var storage = make(map[string]string)

/*SavePair persists the url pairs for retrieval later*/
func SavePair(trimmedURL, longURL string) error {
	storage[trimmedURL] = longURL
	return nil
}
