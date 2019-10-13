# go-api
**go-api** is a simple example API with support for common implementations

[![Build Status](https://travis-ci.com/mrz1836/go-api.svg?branch=master)](https://travis-ci.com/mrz1836/go-api)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-api?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-api)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b6c2832dee5442c7a79b482114100814)](https://www.codacy.com/app/mrz1818/go-api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-api&amp;utm_campaign=Badge_Grade)
[![Release](https://img.shields.io/github/release-pre/mrz1836/go-api.svg?style=flat)](https://github.com/mrz1836/go-api/releases)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](https://godoc.org/github.com/mrz1836/go-api?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-api)

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

### Prerequisite Applications:
- [MySQL](https://www.mysql.com/) or [MariaDB](https://mariadb.org/) with no password set
- [Redis](https://redis.io/)

1) **go-api** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy) and [dep](https://github.com/golang/dep).
```bash
$ go get -u github.com/mrz1836/go-api
$ go get -u github.com/pressly/goose/cmd/goose
$ go get -u -t github.com/volatiletech/sqlboiler
$ go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
```

2) Set your environment variables (or add to your bash profile):
```bash
$ cd ../go-api
$ . scripts/set_env.sh
```

3) Setup a fresh database (if you don't have one already)
```bash
$ . scripts/setup_db.sh
```

4) Run the API!
```bash
$ go run cmd/application/main.go

    starting Go API server...
```

Test your connection to **go-api**
```bash
$ curl -X GET 'http://localhost:3000'

  Welcome to the Go API!
```

### Managing Dependencies

Updating dependencies in **go-api**:
```bash
$ cd ../go-api
$ dep ensure -update -v
```

### Managing Environment Variables
All environment variables are referenced in the [config](config/config.go).

Edit the [`scripts/set_env.sh`](scripts/set_env.sh) file and modify the environment variables - IE:
```bash
export API_SERVER_PORT=3000
```

### Managing Model Generation

Update the `reset_api_database.sql` if you have issues running the model tests
```sql
GRANT ALL ON `dynamic-database-name-generated-from-sql-boiler`.* to 'apiDbTestUser'@'%';
```

Rebuilding the generated models/schema from the database schema:
```bash
$ cd ../go-api
$ . scripts/rebuild_models.sh
```

### Package Dependencies
- [go-api-router](https://github.com/mrz1836/go-api-router) - Fast and lightweight router
- [go-cache](https://github.com/mrz1836/go-cache) - Redis caching made easy
- [go-logger](https://github.com/mrz1836/go-logger) - Local or remote logging
- [go-sanitize](https://github.com/mrz1836/go-sanitize) - Clean data effortlessly
- [goose](https://github.com/pressly/goose) - Database migration
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Extensible data validation
- [viper](https://github.com/spf13/viper) - Go configuration with fangs
- [SQLBoiler](https://github.com/volatiletech/sqlboiler) - Powerful database ORM & model generation
- [cron](github.com/robfig/cron) - Scheduler to run cron jobs and task

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-api).

### Features
- Combination of powerful Go packages all-in-one API solution
- The fastest router: Julien Schmidt's [httprouter](https://github.com/julienschmidt/httprouter)
- The best redis cache package: Gary Burd's [Redigo](https://github.com/gomodule/redigo)
- Powerful database ORM: VolatileTech's [SQLBoiler](https://github.com/volatiletech/sqlboiler)
- Database migration: Pressly's [Goose](https://github.com/pressly/goose)
- Ready for development or production use
- Cache dependency management via [go-cache](https://github.com/mrz1836/go-cache)
- Supports different incoming load balancer setups (/health)
- Logging each request and whenever you need logs (remote via [LogEntries](https://logentries.com/))
- Flexible environment & configuration management using [viper](https://github.com/spf13/viper)
- Built-in scheduler for any cron jobs or delayed tasks

## Examples & Tests
All unit tests and [examples](examples/examples.go) run via [Travis CI](https://travis-ci.com/mrz1836/go-api) and uses [Go version 1.13.x](https://golang.org/doc/go1.13). View the [deployment configuration file](.travis.yml).

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

## Maintainers

[@MrZ](https://github.com/mrz1836)

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-api)

## License

![License](https://img.shields.io/github/license/mrz1836/go-api.svg?style=flat)
