include .env


#variables
APP_NAME:=library
CMD_DIR:=./cmd
BUILD_DIR:=./bin
EXECUTABLE_PATH:="${BUILD_DIR}/cmd.exe"
# ==================================================================================== #
# HELPERS
# ==================================================================================== #
## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# BUILD
# ==================================================================================== #
## build/api: build the cmd/api application
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	@go build -o ${BUILD_DIR} ${CMD_DIR}

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
## run: run the cmd/api application
.PHONY: run
run:
	@CompileDaemon -exclude "./bin" -directory=./ -build="go build -o ${BUILD_DIR} ${CMD_DIR}" -command="${EXECUTABLE_PATH}"

## migrate/new name=$1: create a new database migration
migrate/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## migrate/up: apply all up database migrations
migrate/up:confirm
	migrate -path ./migrations -database ${DB_DSN} up

## migrate/down: revert all down database migrations
migrate/down:confirm
	migrate -path ./migrations -database ${DB_DSN} down

## migrate/fix version=$1: revert all down database migrations to a specific version
migrate/fix:confirm
	migrate -path ./migrations -database ${DB_DSN} force ${version}
# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:vendor format
	@echo 'Vetting code...'
	go vet ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## format: format code
.PHONY: format
format:
	@echo 'Formatting code...'
	go fmt ./...
	staticcheck ./...

## vendor: tidy and vendor dependencies
.PHONY:vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor
