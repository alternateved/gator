.DEFAULT_GOAL := build

fmt:
		go fmt $(go list ./... | grep -v '/database')

lint: fmt
		golangci-lint run

test:
		go test $(go list ./... | grep -v '/database')

clean:
		go clean

build: lint
		go build .

run: build
		go run .

install: build
		go install .

.PHONY: fmt lint test clean build run
