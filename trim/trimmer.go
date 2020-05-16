package trim

import "regexp"

/*ErrorBadURL is a custom error for handling invalid urls*/
type ErrorBadURL struct {
	msg string
}

func (e *ErrorBadURL) Error() string {
	return e.msg
}

/*GetTrimmed generates and returns a shorter link from longURL
It returns an error if trimming fails for some reason*/
func GetTrimmed(longURL string) (string, error) {

}

/*isValidURL checks if a url is valid or not.
Returns true if it is, else false*/
func isValidURL(url string) bool {
	matcher := regexp.MustCompile(`[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	return matcher.MatchString(url)
}
