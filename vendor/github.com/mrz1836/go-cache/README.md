# go-cache
**go-cache** is a simple redis cache dependency system on-top of the famous [redigo](https://github.com/gomodule/redigo) package

| | | | | | | |
|-|-|-|-|-|-|-|
| ![License](https://img.shields.io/github/license/mrz1836/go-cache.svg?style=flat) | [![Report](https://goreportcard.com/badge/github.com/mrz1836/go-cache?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-cache)  | [![Codacy Badge](https://api.codacy.com/project/badge/Grade/b6c2832dee5442c7a79b482114100814)](https://www.codacy.com/app/mrz1818/go-cache?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-cache&amp;utm_campaign=Badge_Grade) |  [![Build Status](https://travis-ci.com/mrz1836/go-cache.svg?branch=master)](https://travis-ci.com/mrz1836/go-cache)   |  [![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme) | [![Release](https://img.shields.io/github/release-pre/mrz1836/go-cache.svg?style=flat)](https://github.com/mrz1836/go-cache/releases) | [![GoDoc](https://godoc.org/github.com/mrz1836/go-cache?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-cache) |

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Installation

**go-cache** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy) and [dep](https://github.com/golang/dep).
```bash
$ go get -u github.com/mrz1836/go-cache
```

Updating dependencies in **go-cache**:
```bash
$ cd ../go-cache
$ dep ensure -update -v
```

### Package Dependencies
- Gary Burd's [Redigo](https://github.com/gomodule/redigo)
- Knowledge of [Redis](https://redis.io/download)

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-cache).

### Features
- Cache Dependencies Between Keys
- Connect via URL
- Better Pool Management & Creation
- Register Scripts
- Helper Methods (Get, Set, HashGet, etc)

## Examples & Tests
All unit tests and [examples](examples/examples.go) run via [Travis CI](https://travis-ci.com/mrz1836/go-cache) and uses [Go version 1.12.x](https://golang.org/doc/go1.12). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```bash
$ cd ../go-cache
$ go test ./... -v
```

Run tests (excluding integration tests)
```bash
$ cd ../go-cache
$ go test ./... -v -test.short
```

View and run the examples:
```bash
$ cd ../go-cache/examples
$ go run examples.go
```

## Benchmarks
Run the Go benchmarks:
```bash
$ cd ../go-cache
$ go test -bench . -benchmem
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

## Usage
View the [examples](examples/examples.go)

Basic implementation:
```golang
package main

import (
	"log"

	"github.com/mrz1836/go-cache"
)

func main() {

	// Create the pool and first connection
	_ = cache.Connect("redis://localhost:6379", 0, 10, 0, 240)

	// Set a key
	_ = cache.Set("key-name", "the-value", "dependent-key-1", "dependent-key-2")

	// Get a key
	value, _ := cache.Get("key-name")
	log.Println("Got value:", value)
	//Output: Got Value: the-value

	// Kill keys by dependency
	keys, _ := cache.KillByDependency("dependent-key-1")
	log.Println("Keys Removed:", keys)
	//Output: Keys Removed: 2
}
```

## Maintainers

[@MrZ1836](https://github.com/mrz1836) | [@kayleg](https://github.com/kayleg)

## Contributing

This project uses Gary Burd's [Redigo](https://github.com/gomodule/redigo) package.

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-cache)

## License

![License](https://img.shields.io/github/license/mrz1836/go-cache.svg?style=flat)
