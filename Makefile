GO ?= go

BINDIR := $(CURDIR)/bin

PROTOC := protoc
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

.PHONY: generate-proto
generate-proto:
	$(PROTOC) \
		--gogo_out=./ \
		--proto_path=./proto:./proto/internal:./vendor \
		./proto/event.proto

.PHONY: bootstrap
bootstrap:
	GOBIN=$(BINDIR) $(GO) install github.com/mailru/easyjson/...
