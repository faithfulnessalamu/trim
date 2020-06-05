package main

import "testing"

// Initialize a memdatabase
var md = &memDatabase{data: make(map[string]string)}

var testKey = "key"
var testValue = "value"

func TestMemDatabaseSave(t *testing.T) {
	if err := md.save(testKey, testValue); err != nil {
		t.Errorf("Expected key to not exist, it does")
	}
}

func TestMemDatabaseSaveDuplicate(t *testing.T) {
	if err := md.save(testKey, testValue); err == nil {
		t.Errorf("Expected error due to key existing, there was no error")
	}
}

func TestMemDatabaseRetrieve(t *testing.T) {
	got, err := md.retrieve(testKey)
	if err != nil {
		t.Error(err)
		return
	}

	if got != testValue {
		t.Errorf("Expected %s, got %s", testValue, got)
	}
}

func TestMemDatabaseRetrieveNonExistent(t *testing.T) {
	randKey := "Nonexistent"
	if _, err := md.retrieve(randKey); err == nil {
		t.Errorf("Expected error due to non existent key, there was no error")
	}
}

func TestDataIsSaved(t *testing.T) {
	key := "Key"
	value := "Value"

	md.save(key, value)

	got, ok := md.data[key]
	if !ok || got != value {
		t.Errorf("Data was not saved correctly")
	}
}
