######################################################################
# @author      : annika
# @file        : Makefile
# @created     : Monday Aug 22, 2022 18:35:58 CEST
######################################################################

# Force using the vendored dependencies
VENDOR := false

HOSTNAME=ix-api.net
NAMESPACE=ix-api
NAME=ixapi
BINARY=terraform-provider-${NAME}
RELEASE_VERSION=0.1.0

# OS Detection
UNAME=$(shell uname)
ifeq ($(UNAME), Darwin)
  OS_ARCH=darwin_amd64
else
  OS_ARCH=linux_amd64
endif

# Set the build and version
BUILD := $(shell git rev-parse --short HEAD)
VERSION := $(shell git tag --points-at HEAD)
ifeq ($(VERSION),)
  VERSION=$(RELEASE_VERSION)
endif


PROVIDER := main.go

CFLAGS := -buildmode=pie
ifneq ($(VENDOR), false)
  CFLAGS += -mod=vendor
endif

LDFLAGS := -s -w -X main.version=$(VERSION) -X main.commit=$(BUILD)

default: install

docs:
	go generate ./...

build:
	CGO_ENABLED=0 go build $(CFLAGS) -o ./bin/$(BINARY) -ldflags '$(LDFLAGS)' $(PROVIDER)
	
	mkdir -p bin/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH} 
	cp bin/${BINARY} bin/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}


all: build docs

install: all
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp bin/${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test ./internal/...

.PHONY: docs
