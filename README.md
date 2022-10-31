# go-api
> Example API with support for common implementations

[![Release](https://img.shields.io/github/release-pre/mrz1836/go-api.svg?logo=github&style=flat&v=1)](https://github.com/mrz1836/go-api/releases)
[![Build Status](https://img.shields.io/github/workflow/status/mrz1836/go-api/run-go-tests?logo=github&v=3)](https://github.com/mrz1836/go-api/actions)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-api?style=flat&v=1)](https://goreportcard.com/report/github.com/mrz1836/go-api)
[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-api)](https://golang.org/)
[![Sponsor](https://img.shields.io/badge/sponsor-MrZ-181717.svg?logo=github&style=flat&v=3)](https://github.com/sponsors/mrz1836)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat)](https://mrz1818.com/?tab=tips&af=go-api)

<br/>

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

<br/>

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

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-api)

[![GoDoc](https://godoc.org/github.com/mrz1836/go-api?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-api)

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
<br/>

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
<br/>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                  Runs multiple commands
clean                Remove previous builds and any test cache data
clean-mods           Remove all the Go mod cache
coverage             Shows the test coverage
db                   Creates a fresh database
diff                 Show the git diff
env                  Creates a fresh database
flush-redis          Wipe out all data in redis (requires redli)
generate             Runs the go generate command in the base of the repo
godocs               Sync the latest tag with GoDocs
help                 Show this help message
install              Install the application
install              Run the Custom installation
install-go           Install the application (Using Native Go)
lint                 Run the golangci-lint application (install if not found)
release              Full production release (creates release in Github)
release              Runs common.release then runs godocs
release-snap         Test the full release (build binaries)
release-test         Full production test release (everything except deploy)
replace-version      Replaces the version in HTML/JS (pre-deploy)
run                  Runs the application
run-examples         Runs all the examples
schema               Run the Model/schema generation
tag                  Generate a new tag and push (tag version=0.0.0)
tag-remove           Remove a tag if found (tag-remove version=0.0.0)
tag-update           Update an existing tag to current commit (tag-update version=0.0.0)
test                 Runs lint and ALL tests
test-ci              Runs all tests via CI (exports coverage)
test-ci-no-race      Runs all tests via CI (no race) (exports coverage)
test-ci-short        Runs unit tests via CI (exports coverage)
test-no-lint         Runs just tests
test-short           Runs vet, lint and tests (excludes integration tests)
test-unit            Runs tests and outputs coverage
uninstall            Uninstall the application (and remove files)
update-linter        Update the golangci-lint package (macOS only)
vet                  Run the Go vet application
```
</details>

<details>
<summary><strong><code>Model Generation</code></strong></summary>
<br/>

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
<br/>

All environment variables are referenced in the [config](config/config.go).

Edit the [`scripts/set_env.sh`](scripts/set_env.sh) file and modify the environment variables - IE:
```shell script
export API_SERVER_PORT=3000
```
</details>

<br/>

## Examples & Tests
All unit tests and examples run via [Github Actions](https://github.com/tonicpow/go-paymail/actions) and
uses [Go version 1.17.x](https://golang.org/doc/go1.17). View the [configuration file](.github/workflows/run-tests.yml).

Run all tests (including integration tests)
```shell script
make test
```

Run tests (excluding integration tests)
```shell script
make test-short
```

<br/>

## Benchmarks
Run the Go benchmarks:
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## Usage
(Coming soon: Examples!)

<br/>

## Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:------------------------------------------------------------------------------------------------:|
|                                [MrZ](https://github.com/mrz1836)                                 |

<br/>

## Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and please follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:! 
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:. 
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap: 
or by making a [**bitcoin donation**](https://mrz1818.com/?tab=tips&af=go-sanitize) to ensure this journey continues indefinitely! :rocket:

<br/>

## License

[![License](https://img.shields.io/github/license/mrz1836/go-api.svg?style=flat&v=1)](LICENSE)
