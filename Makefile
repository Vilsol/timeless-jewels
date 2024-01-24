.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -v -o frontend/static/calculator.wasm ./wasm
dev: wasm
	cd frontend; bun run dev