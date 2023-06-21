all: sqlc-gen-python sqlc-gen-python.wasm

sqlc-gen-python:
	cd plugin && go build -o ~/bin/sqlc-gen-python ./main.go

sqlc-gen-python.wasm:
	cd plugin && GOOS=wasip1 GOARCH=wasm /Users/kyle/projects/goroot/bin/go build -o sqlc-gen-python.wasm main.go
	openssl sha256 plugin/sqlc-gen-python.wasm
