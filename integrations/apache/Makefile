INTEGRATION  := $(shell basename $(shell pwd))
BINARY_NAME   = nr-$(INTEGRATION)
GO_FILES     := $(shell find . -type f -name "*.go")
VALIDATE_DEPS = github.com/golang/lint/golint
TEST_DEPS     = github.com/axw/gocov/gocov github.com/AlekSi/gocov-xml

all: build

build: clean validate compile test

clean:
	@echo "=== $(INTEGRATION) === [ clean ]: removing binaries and coverage file..."
	@rm -rfv bin coverage.xml

validate-deps:
	@echo "=== $(INTEGRATION) === [ validate-deps ]: installing validation dependencies..."
	@go get -v $(VALIDATE_DEPS)

validate-only:
ifeq ($(strip $(GO_FILES)),)
	@echo "=== $(INTEGRATION) === [ validate ]: no Go files found. Skipping validation."
else
	@printf "=== $(INTEGRATION) === [ validate ]: running gofmt... "
	@OUTPUT="$(shell gofmt -l $(GO_FILES))" ;\
	if [ -z "$$OUTPUT" ]; then \
		echo "passed." ;\
	else \
		echo "failed. Incorrect syntax in the following files:" ;\
		echo "$$OUTPUT" ;\
		exit 1 ;\
	fi
	@printf "=== $(INTEGRATION) === [ validate ]: running golint... "
	@OUTPUT="$(shell golint ./...)" ;\
	if [ -z "$$OUTPUT" ]; then \
		echo "passed." ;\
	else \
		echo "failed. Issues found:" ;\
		echo "$$OUTPUT" ;\
		exit 1 ;\
	fi
	@printf "=== $(INTEGRATION) === [ validate ]: running go vet... "
	@OUTPUT="$(shell go vet ./...)" ;\
	if [ -z "$$OUTPUT" ]; then \
		echo "passed." ;\
	else \
		echo "failed. Issues found:" ;\
		echo "$$OUTPUT" ;\
		exit 1;\
	fi
endif

validate: validate-deps validate-only

compile:
	@echo "=== $(INTEGRATION) === [ compile ]: building $(BINARY_NAME)..."
	@go build -v -o bin/$(BINARY_NAME) $(GO_FILES)

test-deps:
	@echo "=== $(INTEGRATION) === [ test-deps ]: installing testing dependencies..."
	@go get -v $(TEST_DEPS)

test-only:
	@echo "=== $(INTEGRATION) === [ test ]: running unit tests..."
	@gocov test ./... | gocov-xml > coverage.xml


test: test-deps test-only

.PHONY: all build clean validate-deps validate-only validate compile test-deps test-only test
