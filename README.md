# go-api
**go-api** is a simple example API with support for various implementations

| | | | | | | |
|-|-|-|-|-|-|-|
| ![License](https://img.shields.io/github/license/mrz1836/go-api.svg?style=flat) | [![Report](https://goreportcard.com/badge/github.com/mrz1836/go-api?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-api)  | [![Codacy Badge](https://api.codacy.com/project/badge/Grade/b6c2832dee5442c7a79b482114100814)](https://www.codacy.com/app/mrz1818/go-api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-api&amp;utm_campaign=Badge_Grade) |  [![Build Status](https://travis-ci.com/mrz1836/go-api.svg?branch=master)](https://travis-ci.com/mrz1836/go-api)   |  [![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme) | [![Release](https://img.shields.io/github/release-pre/mrz1836/go-api.svg?style=flat)](https://github.com/mrz1836/go-api/releases) | [![GoDoc](https://godoc.org/github.com/mrz1836/go-api?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-api) |

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

**go-api** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy) and [dep](https://github.com/golang/dep).
```bash
$ go get -u github.com/mrz1836/go-api
```

Updating dependencies in **go-api**:
```bash
$ cd ../go-api
$ dep ensure -update -v
```

### Package Dependencies
- todo: @mrz

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-api).

### Features
- todo: @mrz

## Examples & Tests
All unit tests and [examples](examples/examples.go) run via [Travis CI](https://travis-ci.com/mrz1836/go-api) and uses [Go version 1.12.x](https://golang.org/doc/go1.12). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```bash
$ cd ../go-api
$ go test ./... -v
```

Run tests (excluding integration tests)
```bash
$ cd ../go-api
$ go test ./... -v -test.short
```

View and run the examples:
```bash
$ cd ../go-api/examples
$ go run examples.go
```

## Benchmarks
Run the Go benchmarks:
```bash
$ cd ../go-api
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
	"github.com/mrz1836/go-api"
)

func main() {

}
```

## Maintainers

[@MrZ1836](https://github.com/mrz1836)

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-api)

## License

![License](https://img.shields.io/github/license/mrz1836/go-api.svg?style=flat)
