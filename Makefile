SHELL = /bin/bash -o pipefail
export PATH := $(shell go env GOPATH)/bin:$(PATH)

ROOT_DIR=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BIN_DIR=$(ROOT_DIR)/build
BIN=$(BIN_DIR)/app

all: clean build run

run:
	@echo "==> Running"
	@${BIN} run

build: deps fmt
	@echo "==> Building for prod"
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=0 go build -o $(BIN) .
	@echo $(BIN)

fmt:
	@echo "==> Running gofmt"
	@gofmt -l -s -w .

deps:
	@echo "==> Installing dependencies"
	@go mod tidy
	@go mod vendor

test: clean build
	@echo "==> Running tests"
	@go test ./... -v -cpu 2 -race -timeout=30s -coverpkg=./... -coverprofile=$(ROOT_DIR)/coverage.out | tee $(ROOT_DIR)/test-report.txt
	@cd / && go install github.com/axw/gocov/gocov@v1.0.0
	@gocov convert $(ROOT_DIR)/coverage.out | gocov report
	@go mod tidy

.PHONY: clean
clean:
ifneq ($(wildcard $(ROOT_DIR)/vendor/.*),)
	@rm -rf ./vendor
	@mkdir ./vendor
	@touch ./vendor/.gitkeep
endif
	@rm -rf coverage.out test-report.{txt,xml}
	@rm -rf $(BIN_DIR)
	@mkdir -p $(BIN_DIR)
	@touch $(BIN_DIR)/.gitkeep
