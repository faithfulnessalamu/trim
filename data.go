package main

//database allows swapping storage for this app and for testing
type database interface {
	saver
	retriever
}

//saver is used by any service that can persist data
type saver interface {
	save(key string, value string) error
}

//retriever is used by any service that can get data by key
type retriever interface {
	retrieve(key string) (string, error)
}
