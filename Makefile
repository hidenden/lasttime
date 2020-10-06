# Meta informations
NAME := lasttime
## VERSION := $(godump show -r)
VERSION := 1
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

export GO111MODULE=on

.PHONY test
test: deps
	go test ./...


