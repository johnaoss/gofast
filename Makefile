# Required: The name of your application
APP=GoFast
# Required: The application's CFBundleIdentifier
# e.g.  IDENTIFIER=whyawake.caseymrm.github.com
IDENTIFIER=gofast.johnaoss.github.com
# To sign your app, set IDENTITY:
#   IDENTITY=Developer ID Application: Hello World LLC (AP2AFA8XAX)
# To get a list of identities, run the following:
#	security find-identity -v -p codesigning
# To generate a new identity if you have none, launch XCode and navigate
# 	to Preferences>Accounts and create one there.
# Currently, I've stored my identity as an env var.
IDENTITY=$(APPLE_CERT)

# Include the menuet specific Makefile
include hack/menuet.mk

# Test just runs all of the tests for this function.
.PHONY: test
test:
	@go test ./...

# reset-preferences resets all user preferences that may have been set by this
# application. 
# TODO: https://superuser.com/questions/907798/deleting-user-defaults-under-mac-os-x-10-10-3/907899#907899
.PHONY: reset-preferences
reset-preferences:
	defaults delete ${IDENTIFIER}