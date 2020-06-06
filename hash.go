package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

// Defines the length of a hash
var truncatedLength = 8

//hash returns a 8 digit hash for a url
func hash(actualURL string) string {
	trimHash := sha1.New()
	io.WriteString(trimHash, actualURL)
	sum := trimHash.Sum(nil)
	truncated := fmt.Sprintf("%x", sum)[:truncatedLength]
	return truncated
}
