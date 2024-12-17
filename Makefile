.PHONY: run

# Version information
LAST_COMMIT := $(or $(shell git rev-parse --short HEAD 2> /dev/null),"unknown")
VERSION := $(or $(shell git describe --tags --abbrev=0 2> /dev/null),"v0.0.0")
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%S%z")
BIN := echo-server

# Build flags
LD_FLAGS := -s -w \
	-X 'main.version=${VERSION}' \
	-X 'main.commit=${LAST_COMMIT}' \
	-X 'main.date=${BUILD_DATE}'

GO ?= $(shell which go)

run:
	CGO_ENABLED=0 $(GO) run -ldflags="${LD_FLAGS}" cmd/*.go

build:
	CGO_ENABLED=0 $(GO) build -o ${BIN} -ldflags="${LD_FLAGS}" cmd/*.go
