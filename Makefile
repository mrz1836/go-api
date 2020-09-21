# Common makefile commands & variables between projects
include .make/Makefile.common

# Common Golang makefile commands & variables between projects
include .make/Makefile.go

## Not defined? Use default repo name which is the application
ifeq ($(REPO_NAME),)
	REPO_NAME="go-api"
endif

## Not defined? Use default repo owner
ifeq ($(REPO_OWNER),)
	REPO_OWNER="mrz1836"
endif

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

.PHONY: clean

all: ## Runs multiple commands
	@$(MAKE) test-short

clean: ## Remove previous builds and any test cache data
	@go clean -cache -testcache -i -r
	@test $(DISTRIBUTIONS_DIR)
	@if [ -d $(DISTRIBUTIONS_DIR) ]; then rm -r $(DISTRIBUTIONS_DIR); fi

db: ## Creates a fresh database
	@mysql -u root < ./database/reset/reset_api_database.sql && goose -dir "./database/sql" mysql "$(DB_USER):$(DB_PASSWORD)@/$(DB_NAME)?parseTime=true" up

env: ## Creates a fresh database
	@ . scripts/set_env.sh

flush-redis: ## Wipe out all data in redis (requires redli)
	@redli --raw FLUSHALL

install: ## Run the Custom installation
	@go get -u github.com/pressly/goose/cmd/goose
	@go get -u -t github.com/volatiletech/sqlboiler/v4
	@go get -u github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql
	@$(MAKE) env
	@$(MAKE) db

schema: ## Run the Model/schema generation
	@sqlboiler mysql --wipe && sed -i "" 's/fmt.Fprintf(tmp, "ssl-mode/\/\/fmt.Fprintf(tmp, "ssl-mode/' models/schema/mysql_main_test.go

release:: ## Runs common.release then runs godocs
	@$(MAKE) godocs

run: ## Runs the application
	@go run cmd/application/main.go

run-examples: ## Runs all the examples
	@go run examples/examples.go