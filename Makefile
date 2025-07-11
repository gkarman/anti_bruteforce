BIN := ./bin/anti_bruteforce
CONFIG := ./configs/anti_bruteforce_config.yaml
BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(BUILD_DATE) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o ./bin/anti_bruteforce -ldflags "$(LDFLAGS)" ./cmd/anti_bruteforce
	go build -v -o ./bin/anti_bruteforce_cli -ldflags "$(LDFLAGS)" ./cmd/anti_bruteforce_cli

migrate:
	go run ./cmd/anti_bruteforce migrate

test:
	go test -v -race -count=100 -timeout=5m ./...

run: build
	$(BIN) -config $(CONFIG)

install-lint-deps:
	@which golangci-lint > /dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.57.2

lint: install-lint-deps
	golangci-lint run

generate:
	go generate ./...

.PHONY: build run migrate test lint install-lint-deps generate
