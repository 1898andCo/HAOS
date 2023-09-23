TARGETS := $(shell ls scripts)
default:
	@bash -c "scripts/ci"


.DEFAULT_GOAL := default

.PHONY: $(TARGETS)
