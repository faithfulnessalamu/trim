package main

import (
	"testing"
)

func TestIsValidURL(t *testing.T) {
	goodURLs := []string{"https://www.github.com/github", "http://homeofthefree.org"}
	badURLs := []string{"hpt//isnota.protocol", "http://..."}

	// Check goodURLs
	for _, url := range goodURLs {
		if !isValidURL(url) {
			t.Errorf("%s should be a valid URL", url)
		}
	}
	// Check badURLs
	for _, url := range badURLs {
		if isValidURL(url) {
			t.Errorf("%s should not be a valid URL", url)
		}
	}
}
