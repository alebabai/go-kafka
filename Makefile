GO             ?= go
GOLANGCI_LINT  ?= golangci-lint

PACKAGES       ?= ./...
TEST_PACKAGES  ?= $(PACKAGES)
COVER_PROFILE  ?= coverage.out
COVER_REPORT   ?= coverage.html

.SILENT:

.PHONY: build
build:
	$(GO) build -v $(PACKAGES)

.PHONY: test
test: COVER_PACKAGES ?= $(shell echo $(TEST_PACKAGES) | tr " " ",")
test:
	$(GO) test -v -race -coverpkg $(COVER_PACKAGES) -coverprofile $(COVER_PROFILE) $(TEST_PACKAGES)

.PHONY: coverage
coverage: test
	$(GO) tool cover -func $(COVER_PROFILE)
	$(GO) tool cover -html $(COVER_PROFILE) -o $(COVER_REPORT)

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run -v

.PHONY: fmt
fmt:
	$(GO) fmt $(PACKAGES)

.PHONY: mod/tidy
mod/tidy:
	$(GO) mod tidy

.PHONY: prepare
prepare: mod/tidy fmt

.PHONY: clean
clean:
	$(GO) clean
	rm -f $(COVER_PROFILE) $(COVER_REPORT)

.PHONY: all
all: prepare build lint test
