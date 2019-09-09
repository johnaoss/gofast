BUILD_PATH="GoFast.App/Contents/MacOS/gofast"

# TODO: Should build into the MacOS directory. May be worth making a separate package to create this
# instead of a Makefile, but the Makefile seems easiest for now.
# I thought mholt had a package for that, but I could be wrong.
.PHONY: test
test:
	@go test ./...

.PHONY: build
build: 
	@go build -o ${BUILD_PATH}
	@echo "Binary located at ${BUILD_PATH}"

.PHONY: release
release: test build