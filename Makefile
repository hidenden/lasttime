# Meta informations
NAME := lasttime
VERSION := 1
REVISION := $(shell git rev-parse --short HEAD)
GIT_VER := $(shell git describe --tags)
LDFLAGS := "-X main.revision=$(REVISION) -X main.version=$(GIT_VER)"

export GO111MODULE=on

## Install dependencies
.PHONY: deps
deps:
	go get -v -d

## 開発に必要な依存をインストール
.PHONY: devel-deps
devel-deps: deps
	go get github.com/golang/lint/golint github.com/Songmu/make2help/cmd/make2help

## Run tests
.PHONY: test
test: deps
	go test ./...

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint --set-exit_status ./...

## Build binaries ex. make bin/myprj
bin/%: cmd/%/main.go deps
	go build -ldflags  $(LDFLAGS) -o $@ $<

## Build binary
.PHONY: build
build: bin/lasttime

## show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)

