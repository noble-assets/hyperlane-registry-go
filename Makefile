.PHONY: license

###############################################################################
###                                 Tooling                                 ###
###############################################################################

license_cmd=github.com/palantir/go-license

FILES := $(shell find . -name "*.go")
license:
	@go run $(license_cmd) --config .github/license.yaml $(FILES)
