.PHONY: compile
compile:
	sqlc compile

.PHONY: generate
generate: sqlc.yaml
	sqlc generate

.PHONY: clean
clean:
	rm -rf ./_examples/gen

sqlc.yaml: dist/sqlc-gen-python-orm.wasm.sha256 _sqlc.yaml
	cat _sqlc.yaml | WASM_SHA256=$$(cat $<) envsubst > $@

dist/sqlc-gen-python-orm.wasm.sha256: dist/sqlc-gen-python-orm.wasm
	openssl sha256 $< | awk '{print $$2}' > $@

dist/sqlc-gen-python-orm.wasm: internal/*
	GOOS=wasip1 GOARCH=wasm go build -o ./dist/sqlc-gen-python-orm.wasm ./plugin/main.go