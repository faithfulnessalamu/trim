package main

import (
	"regexp"
	"testing"
)

func TestHash(t *testing.T) {
	matcher := regexp.MustCompile(`[a-fA-F0-9]{8,8}`)
	toHash := "https://github.com/thealamu/trim"
	trimHash := hash(toHash)
	if !matcher.MatchString(trimHash) {
		t.Errorf("Expected %d digit hex, got %s", truncatedLength, trimHash)
	}
}
