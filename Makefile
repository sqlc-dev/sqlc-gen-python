all: sqlc-gen-python

sqlc-gen-python:
	cd plugin && go build -o ~/bin/sqlc-gen-python ./main.go

sqlc-gen-python.wasm:
	cd plugin && tinygo build -o sqlc-gen-python.wasm -gc=leaking -scheduler=none -wasm-abi=generic -target=wasi main.go
	openssl sha256 plugin/sqlc-gen-python.wasm

