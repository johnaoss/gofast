# Required: The name of your application
APP=GoFast
# To sign your app, set IDENTITY:
#   IDENTITY=Developer ID Application: Hello World LLC (AP2AFA8XAX)
IDENTITY=
# To set your CFBundleIdentifier, set IDENTIFIER:
# e.g.  IDENTIFIER=whyawake.caseymrm.github.com
IDENTIFIER=gofast.johnaoss.github.com

# Include the menuet specific Makefile
# TODO: Fix this for signing.
include hack/menuet.mk

.PHONY: test
test:
	@go test ./...

.PHONY: reset-preferences
reset-preferences:
	defaults delete ${IDENTIFIER}