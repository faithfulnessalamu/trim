package trimmer

import (
	"crypto/sha1"
	"fmt"
	"io"
	"path"
	"regexp"
)

const truncatedLength = 8
const rootPath = "localhost:8080/"

/*ErrorBadURL is a custom error for handling invalid urls*/
type ErrorBadURL struct {
	msg string
}

func (e *ErrorBadURL) Error() string {
	return e.msg
}

/*TODO: ADD DOCSTRING*/
func getDigest(longURL string) string {
	hash := sha1.New()
	io.WriteString(hash, longURL)
	sum := hash.Sum(nil)
	truncated := fmt.Sprintf("%x", sum)[:truncatedLength]
	return truncated
}

/*GetTrimmed generates and returns a shorter link from a validated input.
It returns an error if trimming fails for some reason*/
func GetTrimmed(input string) (string, error) {
	if !isValidURL(input) {
		return "", &ErrorBadURL{msg: "Not a valid URL"}
	}
	// The input has been validated to be a URL
	longURL := input
	digest := getDigest(longURL)

	return path.Join(rootPath + digest), nil
}

/*isValidURL checks if a url is valid or not.
Returns true if it is, else false*/
func isValidURL(input string) bool {
	matcher := regexp.MustCompile(`[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	return matcher.MatchString(input)
}

/* IsValidTrimURL checks if a URL is a trim url */
func IsValidTrimURL(url string) bool {
	matcher := regexp.MustCompile(fmt.Sprintf("^%s/[a-zA-Z0-9]{8,8}", rootPath))
	return matcher.MatchString(url)
}
