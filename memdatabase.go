package main

import (
	"errors"
)

//memDatabase implements the database interface.
//It stores data in memory and is useful for testing
//For simplicity, stores strings only
type memDatabase struct {
	data map[string]string
}

func (m *memDatabase) save(key string, value string) error {
	// Check key existence
	if _, ok := m.data[key]; ok {
		return errors.New("Key exists already")
	}
	m.data[key] = value
	return nil
}

func (m *memDatabase) retriever(key string) (string, error) {
	value, ok := m.data[key]
	if !ok {
		return "", errors.New("Key does not exist")
	}
	return value, nil
}
