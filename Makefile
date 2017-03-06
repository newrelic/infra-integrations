# Global variables
INTEGRATIONS       ?= all
WORKDIR            := $(shell pwd)
INTEGRATIONS_DIR   := $(WORKDIR)/integrations
TARGET_DIR          = $(WORKDIR)/target
MAJORMINOR_VERSION ?= 1.0
PATCH              := $(shell date +'%s')
GIT_BRANCH         := $(shell git rev-parse --abbrev-ref HEAD | sed -e 's/origin\///')

# Version generation. If we're building a non-master branch, use an artificially
# old version including the branch name so it will be deployable but not seen as
# newer than master builds.
ifneq ($(USER),jenkins)
    MAJORMINOR_VERSION = 0.0
    PATCH = dev
else ifneq ($(GIT_BRANCH),master)
    MAJORMINOR_VERSION = 0.0.$(subst /,-,$(GIT_BRANCH))
endif
VERSION = $(MAJORMINOR_VERSION).$(PATCH)

# Select integrations to build
ifeq ($(INTEGRATIONS),all)
    INTS=$(shell find $(INTEGRATIONS_DIR) -mindepth 1 -maxdepth 1 -type d | perl -pe 's/[^\/]*\///g')
else
    INTS=$(INTEGRATIONS)
endif

# Default target
all: build

# Common targets
clean:
	@echo "=== Main === [ clean ]: Removing binaries and coverage files..."
	@rm -rfv $(INTEGRATIONS_DIR)/*/bin $(TARGET_DIR)
	@find . -type f -name "coverage.xml" -delete -print

# Include thematic Makefiles
include Makefile-*.mk

.PHONY: all clean
