GOROOT := $(shell go env GOROOT)

.PHONY=cp_wasm_exec
cp_wasm_exec:
	cp $(GOROOT)/misc/wasm/wasm_exec.js build/web/

.PHONY=web
web: cp_wasm_exec
	GOOS=js GOARCH=wasm go build -o build/web/sszb.wasm github.com/cat-in-the-dark/sszb

.PHONY=build
build:
	go build -o build/sszb.exe
