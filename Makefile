# I am a -*-Makefile-*-
#
# --------------------------------------------------
#

GO_FLAGS =

.PHONY: all
all: build test

.PHONY: build
build: voyeur/voyeur

.PHONY: voyeur/voyeur  # force a rebuild always
voyeur/voyeur:
	( cd voyeur; go build ${GO_FLAGS} )

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:

.PHONY: package
package:
	echo "package as (possibly) multile distros including deb"

.PHONY: clean
clean:
	go clean -cache
	go clean ./...
