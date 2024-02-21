.PHONY: build clean tidy

GOCERTCENTER_SOURCES := $(shell find ./cmd ./internal -type f -iname '*.go' ! -iname '*_test.go')

all: build

build: gocertcenter

tidy:
	go mod tidy

gocertcenter: $(GOCERTCENTER_SOURCES)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o gocertcenter ./cmd/gocertcenter
	chmod 700 ./gocertcenter

test:
	go test -v ./...

clean:
	rm -f gocertcenter
