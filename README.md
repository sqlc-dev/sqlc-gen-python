## Usage

```yaml
version: '2'
plugins:
- name: py
  wasm:
    url: https://downloads.sqlc.dev/plugin/sqlc-gen-python_1.0.0.wasm
    sha256: "FIXME"
sql:
- schema: "schema.sql"
  queries: "query.sql"
  engine: postgresql
  codegen:
  - out: src/authors
    plugin: py
    options:
      package: authors
      emit_sync_querier: true
      emit_async_querier: true
```
