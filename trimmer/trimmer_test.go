package trimmer

import (
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidURL(t *testing.T) {
	goodURL1 := isValidURL("https://www.github.com/github")
	goodURL2 := isValidURL("http://homeofthefree.org")

	badURL1 := isValidURL("hpt//isnota.protocol")
	badURL2 := isValidURL("http://...")

	require := require.New(t)

	require.True(goodURL1, "https://www.github.com/github should be a valid URL")
	require.True(goodURL2, "http://homeofthefree.org should be a valid URL")

	require.False(badURL1, "hpt//isnota.protocol should not be a valid URL")
	require.False(badURL2, "http://... should not be a valid URL")
}

func TestIsValidTrimURL(t *testing.T) {
	goodTrimURL := path.Join(rootPath, "/h67j82er")
	badTrimURL := path.Join(rootPath, "/ht87j")

	require.True(t, IsValidTrimURL(goodTrimURL), "isValidTrimURL rejects a trim URL: "+goodTrimURL)
	require.False(t, IsValidTrimURL(badTrimURL), "isValidTrimURL allows a non-trim URL "+badTrimURL)
}

func TestGetTrimmed(t *testing.T) {
	// Assert GetTrimmed returns empty string and badurlerror when url is invalid
	badURL := "https://github"
	_, err := GetTrimmed(badURL)

	if err == nil {
		t.Fatal("GetTrimmed allows invalid URLs")
	}

	goodURL := "https://www.github.com"
	trimmed, err := GetTrimmed(goodURL)

	if err != nil {
		t.Fatal("GetTrimmed rejects a good URL")
	}

	if !IsValidTrimURL(trimmed) {
		t.Fatal("GetTrimmed returns an invalid URL: " + trimmed)
	}
}

func TestGetDigest(t *testing.T) {
	longURL := "https://www.github.com"
	digest := getDigest(longURL)

	if len(digest) != truncatedLength {
		t.Fatalf("Digest is not %d characters long", truncatedLength)
	}
}

func TestErrorBadURLString(t *testing.T) {
	errMsg := "This is an error"
	err := &ErrorBadURL{msg: errMsg}

	require.Equal(t, err.Error(), errMsg, "String method of ErrorBadURL does not return saved message")
}
