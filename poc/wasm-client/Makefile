
# Path to the Go WebAssembly JS support file
WASM_EXEC_JS := $(shell go env GOROOT)/misc/wasm/wasm_exec.js

# Default target executed when no arguments are given to make.
default: server wasm

# Target for building the Go HTTP server
server: ./cmd/server/main.go wasm
	go build -o server ./cmd/server

# Target for building the WebAssembly module (replace with your actual build command)
wasm: ./go/main.go
	cp $(WASM_EXEC_JS) ./cmd/server/web/
	GOOS=js GOARCH=wasm go build -o ./cmd/server/web/myapp.wasm ./go/main.go

# Target for running the server
run: server
	./server

# Target for cleaning up build artifacts
clean:
	rm -f ./server ./cmd/server/web/myapp.wasm ./cmd/server/web/wasm_exec.js

.PHONY: default wasm run clean
