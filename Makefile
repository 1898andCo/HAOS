# additional scripts are kept in the scripts directory
TARGETS := $(shell ls scripts)
# default to running the ci script
default:
	@bash -c "scripts/ci"


.DEFAULT_GOAL := default

# this will create 'phony' targets for each script in the scripts directory
# run them as eg. `make ci`, `make test`, etc.
.PHONY: $(TARGETS)
