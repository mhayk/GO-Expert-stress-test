MAIN_PACKAGE_PATH := ./cmd/stress-test
BINARY_NAME := stress-test

# ============================================================================ #
# HELPERS
# ============================================================================ #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ============================================================================ #
# Development
# ============================================================================ #

## test: run all tests
.PHONY: test
test:
		go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
		go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
		go tool cover -html=/tmp/coverage.out

## build: build the application
.PHONY: build
build:
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the  application
.PHONY: run
run: build
	/tmp/bin/$(BINARY_NAME) --url=https://google.com --requests=100 --concurrency=10