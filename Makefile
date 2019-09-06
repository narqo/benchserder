GO ?= go
PROTOC ?= protoc

BINDIR := $(CURDIR)/bin

EASYJSON := $(BINDIR)/easyjson

SHELL := /bin/bash
PATH := $(PATH):$(BINDIR)

.DEFAULT_GOAL := all
.SUFFIXES:

.PHONY: all
all:

.PHONY: generate-json
generate-json:
	$(EASYJSON) -no_std_marshalers event.go

PROTO_FILES := $(wildcard ./internal/**/*.proto ./*.proto)

.PHONY: generate-proto
generate-proto: $(PROTO_FILES:.proto=.pb.go)

%.pb.go: %.proto
	$(PROTOC) \
		-I ./:vendor/ \
		--gogo_out=. $(<)

.PHONY: generate-proto-src
generate-proto-src:
	GOPATH=$$HOME/Workspace \
		   GO111MODULE=off \
		proteus proto \
			-p github.com/narqo/benchserder \
			-p github.com/narqo/benchserder/internal/callback \
			-p github.com/narqo/benchserder/internal/device \
			-p github.com/narqo/benchserder/internal/engagementtype \
			-p github.com/narqo/benchserder/internal/fraud \
			-p github.com/narqo/benchserder/internal/header \
			-p github.com/narqo/benchserder/internal/nullable \
			-p github.com/narqo/benchserder/internal/money \
			-p github.com/narqo/benchserder/internal/replacer \
			-p github.com/narqo/benchserder/internal/tokens \
			-f ../../../../src/ --verbose

.PHONY: bootstrap
bootstrap:
	GOBIN=$(BINDIR) $(GO) install github.com/mailru/easyjson/...
