COVER=go tool cover

## Default Repo Domain
GIT_DOMAIN=github.com

## Set the Github Token
#GITHUB_TOKEN=<your_token>

## Default DB user name
DB_NAME=api_example
DB_USER=apiDbTestUser
DB_PASSWORD=ThisIsSecureEnough123

## Default DB user's password

## Automatically detect the repo owner and repo name
REPO_NAME=$(shell basename `git rev-parse --show-toplevel`)
REPO_OWNER=$(shell git config --get remote.origin.url | sed 's/git@$(GIT_DOMAIN)://g' | sed 's/\/$(REPO_NAME).git//g')

## Set the version(s) (injected into binary)
VERSION=$(shell git describe --tags --always --long --dirty)
VERSION_SHORT=$(shell git describe --tags --always --abbrev=0)

.PHONY: test lint clean release

all: lint test-short vet

bench:  ## Run all benchmarks in the Go application
	go test -bench ./... -benchmem -v

clean: ## Remove previous builds and any test cache data
	go clean -cache -testcache -i -r
	if [ -d ${DISTRIBUTIONS_DIR} ]; then rm -r ${DISTRIBUTIONS_DIR}; fi

clean-mods: ## Remove all the Go mod cache
	go clean -modcache

coverage: ## Shows the test coverage
	go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

db: ## Creates a fresh database
	mysql -u root < ./database/reset/reset_api_database.sql && goose -dir "./database/sql" mysql "$(DB_USER):$(DB_PASSWORD)@/$(DB_NAME)?parseTime=true" up

env: ## Creates a fresh database
	. scripts/set_env.sh

godocs: ## Sync the latest tag with GoDocs
	curl https://proxy.golang.org/$(GIT_DOMAIN)/$(REPO_OWNER)/$(REPO_NAME)/@v/$(VERSION_SHORT).info

flush-redis: ## Wipe out all data in redis
	redli --raw FLUSHALL

help: ## Show all make commands available
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Run the Custom installation
	go get -u github.com/pressly/goose/cmd/goose
	go get -u -t github.com/volatiletech/sqlboiler
	go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
	make env
	make db

lint: ## Run the Go lint application
	golint

schema: ## Run the Model/schema generation
	sqlboiler mysql --wipe && sed -i "" 's/fmt.Fprintf(tmp, "ssl-mode/\/\/fmt.Fprintf(tmp, "ssl-mode/' models/schema/mysql_main_test.go

release: ## Full production release (creates release in Github)
	 goreleaser --rm-dist
	 make godocs

release-test: ## Full production test release (everything except deploy)
	 goreleaser --skip-publish --rm-dist

release-snap: ## Test the full release (build binaries)
	goreleaser --snapshot --skip-publish --rm-dist

run: ## Runs the application
	go run cmd/application/main.go

run-examples: ## Runs all the examples
	go run examples/examples.go

tag: ## Generate a new tag and push (IE: make tag version=0.0.0)
	test $(version)
	git tag -a v$(version) -m "Pending full release..."
	git push origin v$(version)
	git fetch --tags -f

tag-remove: ## Remove a tag if found (IE: make tag-remove version=0.0.0)
	test $(version)
	git tag -d v$(version)
	git push --delete origin v$(version)
	git fetch --tags

tag-update: ## Update an existing tag to current commit (IE: make tag-update version=0.0.0)
	test $(version)
	git push --force origin HEAD:refs/tags/v$(version)
	git fetch --tags -f

test: ## Runs vet, lint and ALL tests
	make vet
	make lint
	go test ./... -v

test-short: ## Runs vet, lint and tests (excludes integration tests)
	make vet
	make lint
	go test ./... -v -test.short

update:  ## Update all project dependencies
	go get -u ./...
	go mod tidy

update-releaser:  ## Update the goreleaser application
	brew update
	brew upgrade goreleaser

vet: ## Run the Go vet application
	go vet -v cmd/service/main.go