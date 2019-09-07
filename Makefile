PKG := github.com/narqo/benchserder

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

PROTO_FILES := $(shell find proto -name '*.proto')

.PHONY: generate-proto
generate-proto:
	for proto_file in $(PROTO_FILES); do \
		$(PROTOC) \
			-I proto:vendor \
			--gofast_out=../../../ \
			$$proto_file; \
	done

.PHONY: generate-proto-src
generate-proto-src:
	mkdir -p proto
	GOPATH=$$HOME/Workspace \
		   GO111MODULE=off \
		$(BINDIR)/proteus proto \
			-p $(PKG) \
			-p $(PKG)/internal/callback \
			-p $(PKG)/internal/device \
			-p $(PKG)/internal/engagementtype \
			-p $(PKG)/internal/fraud \
			-p $(PKG)/internal/money \
			-p $(PKG)/internal/nullable \
			-p $(PKG)/internal/replacer \
			-p $(PKG)/internal/tokens \
			-f proto \
			--verbose

.PHONY: bootstrap
bootstrap:
	GO111MODULE=off GOBIN=$(BINDIR) $(GO) install ./vendor/github.com/mailru/easyjson/...
	GO111MODULE=off GOBIN=$(BINDIR) $(GO) install ./vendor/gopkg.in/src-d/proteus.v1/...
	GO111MODULE=off GOBIN=$(BINDIR) $(GO) install ./vendor/github.com/gogo/protobuf/protoc-gen-gofast/...

.PHONY: clean
clean:
	$(RM) -r proto

.PHONY: distclean
distclean: clean
	$(RM) -r $(GOBIN)
	$(GO) clean -cache
