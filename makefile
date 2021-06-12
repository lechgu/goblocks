.PHONY: wasm

wasm:
	rm -rf cmd/server/html
	mkdir cmd/server/html
	cp assets/index.html cmd/server/html/
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js cmd/server/html/
	GOOS=js GOARCH=wasm go build -o cmd/server/html/main.wasm cmd/desktop/main.go