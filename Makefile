VERSION := $$(make -s show-version)
CURRENT_REVISION := $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS := "-s -w -X main.revision=$(CURRENT_REVISION)"
GOBIN ?= $(shell go env GOPATH)/bin

.PHONY: show-version
show-version: $(GOBIN)/gobump
	@gobump show -r .

$(GOBIN)/gobump:
	@go install github.com/x-motemen/gobump/cmd/gobump@latest

.PHONY: compile
compile:
	sqlc compile

.PHONY: generate
generate: sqlc.yaml
	sqlc generate


.PHONY: release
release: dist/sqlc-gen-ts-d1.wasm dist/sqlc-gen-ts-d1.wasm.sha256
	gh release delete -y --cleanup-tag v0.0.0-a
	gh release create v0.0.0-a dist/sqlc-gen-ts-d1.wasm dist/sqlc-gen-ts-d1.wasm.sha256

.PHONY: clean
clean:
	rm -rf ./_examples/gen

sqlc.yaml: dist/sqlc-gen-python-orm.wasm.sha256 _sqlc.yaml
	cat _sqlc.yaml | WASM_SHA256=$$(cat $<) envsubst > $@

dist/sqlc-gen-python-orm.wasm.sha256: dist/sqlc-gen-python-orm.wasm
	openssl sha256 $< | awk '{print $$2}' > $@

dist/sqlc-gen-python-orm.wasm: internal/*
	GOOS=wasip1 GOARCH=wasm go build -o $@ ./cmd/sqlc-gen-python-orm/main.go
