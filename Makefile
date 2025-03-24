# Common makefile commands & variables between projects
include .make/common.mk

# Common Golang makefile commands & variables between projects
include .make/go.mk

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

.PHONY: all
all: ## Runs multiple commands
	@$(MAKE) test-short

.PHONY: clean
clean: ## Remove previous builds and any test cache data
	@go clean -cache -testcache -i -r
	@test $(DISTRIBUTIONS_DIR)
	@if [ -d $(DISTRIBUTIONS_DIR) ]; then rm -r $(DISTRIBUTIONS_DIR); fi

.PHONY: db
db: ## Creates a fresh database
	@mysql -u root < ./database/reset/reset_api_database.sql && goose -dir "./database/sql" mysql "$(DB_USER):$(DB_PASSWORD)@/$(DB_NAME)?parseTime=true" up

.PHONY: env
env: ## Creates a fresh database
	@ . scripts/set_env.sh

.PHONY: flush-redis
flush-redis: ## Wipe out all data in redis (requires redli)
	@redli --raw FLUSHALL

#.PHONY: install
#install: ## Run the Custom installation
#	@go get -u github.com/pressly/goose/cmd/goose
#	@go get -u -t github.com/volatiletech/sqlboiler/v4
#	@go get -u github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql
#	@$(MAKE) env
#	@$(MAKE) db

.PHONY: schema
schema: ## Run the Model/schema generation
	@sqlboiler mysql --wipe --no-tests #&& sed -i "" 's/fmt.Fprintf(tmp, "ssl-mode/\/\/fmt.Fprintf(tmp, "ssl-mode/' models/schema/mysql_main_test.go

.PHONY: release
release:: ## Runs common.release then runs godocs
	@$(MAKE) godocs

.PHONY: run
run: ## Runs the application
	@go run cmd/application/main.go

.PHONY: run-examples
run-examples: ## Runs all the examples
	@go run examples/examples.go