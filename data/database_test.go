package data

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveAndGet(t *testing.T) {
	trimmedURL := "trim.ly/r@nd0mS7"
	longURL := "something.com/randomstring"

	SavePair(trimmedURL, longURL)
	rLongURL, _ := GetLongURL(trimmedURL)

	require.Equal(t, longURL, rLongURL, "Database does not return right URL")
}

func TestSavePair(t *testing.T) {
	trimmedURL := "trim.ly/r@nd0mS6"
	longURL := "something.com/randomstring"

	SavePair(trimmedURL, longURL)
	err := SavePair(trimmedURL, longURL)
	if err == nil {
		t.Fatal("SavePair function fails to check that a key already exists")
	}
}

func TestGetLongURL(t *testing.T) {
	trimmedURL := "trim.ly/r@nd0mS8"

	_, err := GetLongURL(trimmedURL)
	if err == nil {
		t.Fatal("GetLongURL returns data for key that does not exists")
	}
}
