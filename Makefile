.PHONY: build clean tidy

GOCERTCENTER_SOURCES := $(shell find ./*.go ./cmd ./internal -type f -iname '*.go' ! -iname '*_test.go')

all: build

build: gocertcenter

tidy:
	go mod tidy

docs: api.html api.md

gocertcenter: $(GOCERTCENTER_SOURCES) Makefile
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o gocertcenter ./cmd/gocertcenter
	chmod 700 ./gocertcenter

openapi.json: $(GOCERTCENTER_SOURCES) Makefile
	mkdir -p ./tmp
	curl http://localhost:8080/documentation/json -o ./tmp/openapi.json
	mv -f ./tmp/openapi.json openapi.json

# API docs as HTML from OpenAPI specification
api.html: openapi.json Makefile
	mkdir -p ./tmp
	swagger-codegen generate -i openapi.json -l html -o ./tmp
	mv -f ./tmp/index.html api.html

# Markdown version for the API docs
api.md: api.html Makefile
	mkdir -p ./tmp
	pandoc ./api.html -f html -t markdown -o ./tmp/api.md
	mv -f ./tmp/api.md api.md

test: Makefile
	go test -v ./...

clean:
	rm -f gocertcenter

clean-docs:
	rm -f ./api.html ./api.md ./openapi.html ./openapi.json ./tmp/.swagger-codegen-ignore ./tmp/.swagger-codegen/VERSION
	test -e ./tmp/.swagger-codegen && rmdir ./tmp/.swagger-codegen || true
	test -e ./tmp && rmdir ./tmp || true
