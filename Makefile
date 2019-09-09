APP=GoFast
IDENTITY=Developer ID Application: TODO (AP2AEA9WAW)
IDENTIFIER=gofast.johnaoss.github.com

# Include the menuet specific Makefile
include hack/menuet.mk

.PHONY: test
test:
	@go test ./...