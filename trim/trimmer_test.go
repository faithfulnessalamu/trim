package trim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidURL(t *testing.T) {
	goodURL1 := isValidURL("https://www.github.com/github")
	goodURL2 := isValidURL("http://homeofthefree.org")

	badURL1 := isValidURL("hpt//isnota.protocol")
	badURL2 := isValidURL("http://...")

	require := require.New(t)
	require.True(goodURL1)
	require.True(goodURL2)
	require.False(badURL1)
	require.False(badURL2)
}
