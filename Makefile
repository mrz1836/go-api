## Default database name
ifndef DB_NAME
override DB_NAME=api_example
endif

## Default database user
ifndef DB_USER
override DB_USER=apiDbTestUser
endif

## Default database password
ifndef DB_PASSWORD
override DB_PASSWORD=ThisIsSecureEnough123
endif

## Default Repo Domain
GIT_DOMAIN=github.com

## Automatically detect the repo owner and repo name
REPO_NAME=$(shell basename `git rev-parse --show-toplevel`)
REPO_OWNER=$(shell git config --get remote.origin.url | sed 's/git@$(GIT_DOMAIN)://g' | sed 's/\/$(REPO_NAME).git//g')

## Set the version (for go docs)
VERSION_SHORT=$(shell git describe --tags --always --abbrev=0)

## Set the distribution folder
ifndef DISTRIBUTIONS_DIR
override DISTRIBUTIONS_DIR=./dist
endif

.PHONY: test lint clean release

all: test-short ## Runs lint, test-short and vet

bench:  ## Run all benchmarks in the Go application
	@go test -bench ./... -benchmem -v

clean: ## Remove previous builds and any test cache data
	@go clean -cache -testcache -i -r
	@if [ -d $(DISTRIBUTIONS_DIR) ]; then rm -r $(DISTRIBUTIONS_DIR); fi

clean-mods: ## Remove all the Go mod cache
	@go clean -modcache

coverage: ## Shows the test coverage
	@go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

db: ## Creates a fresh database
	@mysql -u root < ./database/reset/reset_api_database.sql && goose -dir "./database/sql" mysql "$(DB_USER):$(DB_PASSWORD)@/$(DB_NAME)?parseTime=true" up

env: ## Creates a fresh database
	. scripts/set_env.sh

godocs: ## Sync the latest tag with GoDocs
	@curl https://proxy.golang.org/$(GIT_DOMAIN)/$(REPO_OWNER)/$(REPO_NAME)/@v/$(VERSION_SHORT).info

flush-redis: ## Wipe out all data in redis
	@redli --raw FLUSHALL

help: ## Show all make commands available
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Run the Custom installation
	@go get -u github.com/pressly/goose/cmd/goose
	@go get -u -t github.com/volatiletech/sqlboiler
	@go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
	@$(MAKE) env
	@$(MAKE) db

lint: ## Run the Go lint application
	@if [ "$(shell command -v golint)" == "" ]; then go get -u golang.org/x/lint/golint; fi
	@golint

schema: ## Run the Model/schema generation
	@sqlboiler mysql --wipe && sed -i "" 's/fmt.Fprintf(tmp, "ssl-mode/\/\/fmt.Fprintf(tmp, "ssl-mode/' models/schema/mysql_main_test.go

release: ## Full production release (creates release in Github)
	@goreleaser --rm-dist
	@$(MAKE) godocs

release-test: ## Full production test release (everything except deploy)
	@goreleaser --skip-publish --rm-dist

release-snap: ## Test the full release (build binaries)
	@goreleaser --snapshot --skip-publish --rm-dist

run: ## Runs the application
	@go run cmd/application/main.go

run-examples: ## Runs all the examples
	@go run examples/examples.go

tag: ## Generate a new tag and push (IE: tag version=0.0.0)
	@test $(version)
	@git tag -a v$(version) -m "Pending full release..."
	@git push origin v$(version)
	@git fetch --tags -f

tag-remove: ## Remove a tag if found (IE: tag-remove version=0.0.0)
	@test $(version)
	@git tag -d v$(version)
	@git push --delete origin v$(version)
	@git fetch --tags

tag-update: ## Update an existing tag to current commit (IE: tag-update version=0.0.0)
	@test $(version)
	@git push --force origin HEAD:refs/tags/v$(version)
	@git fetch --tags -f

test: ## Runs vet, lint and ALL tests
	@$(MAKE) vet
	@$(MAKE) lint
	@go test ./... -v

test-short: ## Runs vet, lint and tests (excludes integration tests)
	@$(MAKE) vet
	@$(MAKE) lint
	@go test ./... -v -test.short

update:  ## Update all project dependencies
	@go get -u ./...
	@go mod tidy

update-releaser:  ## Update the goreleaser application
	@brew update
	@brew upgrade goreleaser

vet: ## Run the Go vet application
	@go vet -v cmd/service/main.go