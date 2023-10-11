# additional scripts are kept in the scripts directory
TARGETS := $(shell ls scripts)

# one make target per script
.PHONY: $(TARGETS)
$(TARGETS):
	@echo "Running $@"
	@bash -c "scripts/$@"
# default to running the ci script
# this overrides the default script 
# that would otherwise be run
default: ci
