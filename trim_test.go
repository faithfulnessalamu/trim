package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidTrimURL(t *testing.T) {
	goodTrimURL := rootPath + "/h67j82er"
	badTrimURL := rootPath + "/ht87j"

	require.True(t, isValidTrimURL(goodTrimURL), "isValidTrimURL rejects a trim URL: "+goodTrimURL)
	require.False(t, isValidTrimURL(badTrimURL), "isValidTrimURL allows a non-trim URL "+badTrimURL)
}
