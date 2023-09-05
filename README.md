# sqlc-gen-python-orm

sqlc-gen-python-orm is a plugin for [sqlc](https://sqlc.dev/) that generates an ORM (now, support SQLAlchemy only) for Python.

## Usage

```yaml
version: '2'
plugins:
- name: py
  wasm:
    url: https://downloads.sqlc.dev/plugin/sqlc-gen-python_1.1.0.wasm
    sha256: ef58f143a8c116781091441770c7166caaf361dd645f62b8f05f462e9f95c3b2
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

## Refs

- [sqlc plugin を書こう - 薄いブログ](https://orisano.hatenablog.com/entry/2023/09/06/010926)
