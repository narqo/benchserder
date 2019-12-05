PKG := github.com/narqo/benchserder

GO ?= go
GOIMPORTS ?= goimports
PROTOC ?= protoc

BINDIR := $(CURDIR)/bin
TESTFLAGS :=

SHELL := /bin/bash
PATH := $(PATH):$(BINDIR)

.DEFAULT_GOAL := all
.SUFFIXES:

.PHONY: all
all:

.PHONY: test
test:
	$(GO) test -mod=vendor $(TESTFLAGS) ./...

.PHONY: generate-json
generate-json:
	$(BINDIR)/easyjson -no_std_marshalers event.go
	$(BINDIR)/ffjson event.go \
		&& sed -i'' -e 's/) UnmarshalJSON(/) UnmarshalFFJSON(/' -e ') MarshalJSON(/) MarshalFFJSON(' event_ffjson.go

PROTO_FILES := $(shell find proto -name '*.proto')
PB_GO_FILES := $(patsubst proto/$(PKG)/%.proto,%.pb.go,$(PROTO_FILES))

.PHONY: generate-proto
generate-proto: $(PB_GO_FILES)

%.pb.go:
	$(PROTOC) \
		-I proto:vendor \
		--gogofaster_out=../../../ \
		$(addprefix proto/$(PKG)/,$(@:.pb.go=.proto))
	$(GOIMPORTS) -w $(@)

.PHONY: generate-proto-src
generate-proto-src:
	mkdir -p proto
	GOPATH=$$HOME/Workspace \
		   GO111MODULE=off \
		$(BINDIR)/proteus proto \
			-p $(PKG) \
			-p $(PKG)/internal/callback \
			-p $(PKG)/internal/money \
			-p $(PKG)/internal/nullable \
			-f proto \
			--verbose

.PHONY: bootstrap
bootstrap:
	GO111MODULE=off GOBIN=$(BINDIR) $(GO) install ./vendor/github.com/mailru/easyjson/...
	GO111MODULE=off GOBIN=$(BINDIR) $(GO) install ./vendor/github.com/pquerna/ffjson/...
	GO111MODULE=off GOBIN=$(BINDIR) $(GO) install ./vendor/gopkg.in/src-d/proteus.v1/...
	GO111MODULE=off GOBIN=$(BINDIR) $(GO) install ./vendor/github.com/gogo/protobuf/protoc-gen-gofast/...
	GO111MODULE=off GOBIN=$(BINDIR) $(GO) install ./vendor/github.com/gogo/protobuf/protoc-gen-gogofaster/...

.PHONY: clean
clean:
	$(RM) $(PB_GO_FILES)
	$(RM) -r proto

.PHONY: distclean
distclean: clean
	$(RM) -r $(GOBIN)
	$(GO) clean -cache
