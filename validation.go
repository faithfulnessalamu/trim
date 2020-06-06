package main

import "regexp"

// Returns true if string is a valid url
func isValidURL(str string) bool {
	matcher := regexp.MustCompile(`[a-zA-Z]*:\/\/[a-z\.@]+\.[a-z]{2,5}[\/]?\S*`)
	return matcher.MatchString(str)
}
