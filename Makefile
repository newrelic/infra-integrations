# Global variables
INTEGRATIONS     ?= all
WORKDIR          := $(shell pwd)
INTEGRATIONS_DIR := $(WORKDIR)/integrations
TARGET_DIR        = $(WORKDIR)/target

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
