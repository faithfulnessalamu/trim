package main

//memDatabase implements the database interface.
//It stores data in memory and is useful for testing
//For simplicity, stores strings only
type memDatabase struct {
	data map[string]string
}

func (m *memDatabase) save(key string, value string) error {
	// Check key existence
	if _, ok := m.data[key]; ok {
		return errKeyExists
	}
	m.data[key] = value
	return nil
}

func (m *memDatabase) retrieve(key string) (string, error) {
	value, ok := m.data[key]
	if !ok {
		return "", errKeyNotFound
	}
	return value, nil
}
