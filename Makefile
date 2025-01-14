# I am a -*-Makefile-*-
#
# --------------------------------------------------

.PHONY: all
all: build test

GIT_SUMMARY := $(shell git describe --tags --dirty --always)
BUILD_DATE  := $(shell date -u "+%Y-%m-%dT%H:%M:%SZ")

# --------------------------------------------------
GO_FLAGS = 

TOOLS_BIN_DIR := $(CURDIR)/tools/bin
GORELEASER_VERSION ?= latest
GORELEASER := $(TOOLS_BIN_DIR)/goreleaser

$(GORELEASER):
	GOBIN=$(TOOLS_BIN_DIR) go install github.com/goreleaser/goreleaser/v2@$(GORELEASER_VERSION)

check: $(GORELEASER)
	$(GORELEASER) check

package: $(GORELEASER)
	$(GORELEASER) release --skip=publish --snapshot

.PHONY: build
build: voyeur

.PHONY: voyeur  # force a rebuild always
voyeur:
	go build -ldflags "-X main.GitSummary=$(GIT_SUMMARY) -X main.BuildDate=${BUILD_DATE}" ${GO_FLAGS} -o $@

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:

.PHONY: clean
clean:
	go clean -cache
	go clean ./...
	rm -rf dist
