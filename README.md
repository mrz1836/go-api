# go-api
> Example API with support for common implementations

[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-api)](https://golang.org/)
[![Build Status](https://travis-ci.com/mrz1836/go-api.svg?branch=master)](https://travis-ci.com/mrz1836/go-api)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-api?style=flat&v=1)](https://goreportcard.com/report/github.com/mrz1836/go-api)
[![Release](https://img.shields.io/github/release-pre/mrz1836/go-api.svg?style=flat&v=1)](https://github.com/mrz1836/go-api/releases)
[![GoDoc](https://godoc.org/github.com/mrz1836/go-api?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-api)

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

**1)** **go-api** requires [Go](https://golang.org/doc/devel/release.html#policy), [Redis](https://redis.io/) and [MySQL](https://www.mysql.com/) with no password set 
```shell script
go get -u github.com/mrz1836/go-api
make install
```

**2)** Run the API
```shell script
make run

  "starting Go API server..."
```

_Test your connection to the api_
```shell script
curl -X GET 'http://localhost:3000'

  "Welcome to the Go API!"
```

## Documentation
You can view the generated [documentation here](https://pkg.go.dev/github.com/mrz1836/go-api).

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
- Powerful and easy emailing with support for [Postmark](https://postmarkapp.com), [Mandrill](https://mandrillapp.com), [AWS SES](https://aws.amazon.com/ses/) and [SMTP](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol)

<details>
<summary><strong><code>Package Dependencies</code></strong></summary>

- [cron](https://github.com/robfig/cron/v3) - Run cron jobs with ease
- [go-api-router](https://github.com/mrz1836/go-api-router) - Fast and lightweight router
- [go-cache](https://github.com/mrz1836/go-cache) - Redis caching made easy
- [go-logger](https://github.com/mrz1836/go-logger) - Local or remote logging
- [go-mail](https://github.com/mrz1836/go-mail) - Email using multiple providers
- [go-sanitize](https://github.com/mrz1836/go-sanitize) - Clean data effortlessly
- [goose](https://github.com/pressly/goose) - Database migration
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Extensible data validation
- [SQLBoiler](https://github.com/volatiletech/sqlboiler) - Powerful database ORM & model generation
- [viper](https://github.com/spf13/viper) - Go configuration with fangs
</details>

<details>
<summary><strong><code>Library Deployment</code></strong></summary>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
bench                          Run all benchmarks in the Go application
clean                          Remove previous builds and any test cache data
clean-mods                     Remove all the Go mod cache
coverage                       Shows the test coverage
db                             Creates a fresh database
env                            Creates a fresh database
godocs                         Sync the latest tag with GoDocs
flush-redis                    Wipe out all data in redis
help                           Show all make commands available
install                        Run the Custom installation
lint                           Run the Go lint application
schema                         Run the Model/schema generation
release                        Full production release (creates release in Github)
release-test                   Full production test release (everything except deploy)
release-snap                   Test the full release (build binaries)
run-examples                   Runs all the examples
tag                            Generate a new tag and push (IE: make tag version=0.0.0)
tag-remove                     Remove a tag if found (IE: make tag-remove version=0.0.0)
tag-update                     Update an existing tag to current commit (IE: make tag-update version=0.0.0)
test                           Runs vet, lint and ALL tests
test-short                     Runs vet, lint and tests (excludes integration tests)
update                         Update all project dependencies
update-releaser                Update the goreleaser application
vet                            Run the Go vet application
```
</details>

<details>
<summary><strong><code>Model Generation</code></strong></summary>

Update the `reset_api_database.sql` if you have issues running the model tests
```sql
GRANT ALL ON `dynamic-database-name-generated-from-sql-boiler`.* to 'apiDbTestUser'@'%';
```

Rebuilding the generated models/schema from the database schema:
```shell script
make schema
```

Clear local redis and reload the database
```shell script
make db
make flush-redis
```
</details>

<details>
<summary><strong><code>Environment Variables</code></strong></summary>

All environment variables are referenced in the [config](config/config.go).

Edit the [`scripts/set_env.sh`](scripts/set_env.sh) file and modify the environment variables - IE:
```shell script
export API_SERVER_PORT=3000
```
</details>

## Examples & Tests
All unit tests run via [Travis CI](https://travis-ci.com/mrz1836/go-api) and uses [Go version 1.14.x](https://golang.org/doc/go1.14). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```shell script
make test
```

Run tests (excluding integration tests)
```shell script
make test-short
```

## Benchmarks
Run the Go benchmarks:
```shell script
make bench
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

## Usage
(Coming soon: Examples!)

## Maintainers

| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:---:|
| [MrZ](https://github.com/mrz1836) |

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-api)

## License

![License](https://img.shields.io/github/license/mrz1836/go-api.svg?style=flat&v=1)
