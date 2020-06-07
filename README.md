# Trim - URL Shortener

[![Build Status](https://travis-ci.com/thealamu/trim.svg?branch=master)](https://travis-ci.com/thealamu/trim) [![codecov](https://codecov.io/gh/thealamu/trim/branch/master/graph/badge.svg)](https://codecov.io/gh/thealamu/trim) [![Go Report Card](https://goreportcard.com/badge/github.com/thealamu/trim)](https://goreportcard.com/report/github.com/thealamu/trim)

Trim is a URL shortener built with golang using a redis backing database. An instance is deployed at [trim.ly](https://trimly.herokuapp.com).

## Features
- Fast API response and redirect
- Statistics for your URLs - See how frequently users access your website
- Private deploys, it's free and open source
- Extendability, you can hot swap backing services such as the database

## API :rocket:
- [x] GET /
- [x] POST /trim

## Development

### Requirements
- [Go](https://golang.org)
- [Redis](https://redis.io/)

### Build
Isolated dependencies, it's easy to build from source.
```go
go get github.com/thealamu/trim
```
### Testing
```go
cd trim
go test ./...
```
## Contributing
Start by forking this project, we would love to receive your pull requests.
