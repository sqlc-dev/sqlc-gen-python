all: sqlc-gen-python sqlc-gen-python.wasm

sqlc-gen-python:
	cd cmd/sqlc-gen-python && go build -o ~/bin/sqlc-gen-python ./main.go

sqlc-gen-python.wasm:
	cd cmd/sqlc-gen-python && tinygo build -o sqlc-gen-python.wasm -gc=leaking -scheduler=none -wasm-abi=generic -target=wasi main.go
	openssl sha256 cmd/sqlc-gen-python/sqlc-gen-python.wasm

