# sqlc-gen-python-orm

sqlc-gen-python-orm is a plugin for [sqlc](https://sqlc.dev/) that generates an ORM (now, support SQLAlchemy only) for Python.

This softwaer forked from [sqlc-gen-python](https://github.com/sqlc-dev/sqlc-gen-python) and modified to generate ORM.

## Usage

get sha256 hash of wasm file

```bash
curl -sSL https://github.com/veltiosoft/sqlc-gen-python-orm/releases/download/v0.0.1/sqlc-gen-python-orm.wasm.sha256
```

add plugin to sqlc.yaml
```yaml
version: '2'
plugins:
- name: py
  wasm:
    url: https://github.com/veltiosoft/sqlc-gen-python-orm/releases/download/v0.0.1/sqlc-gen-python-orm.wasm
    sha256: <sha256 hash>
sql:
- schema: "schema.sql"
  queries: "query.sql"
  engine: postgresql
  codegen:
  - out: src/authors
    plugin: py
    options:
      package: .
      emit_sync_querier: true
      emit_async_querier: true
```

## Refs

- [sqlc plugin を書こう - 薄いブログ](https://orisano.hatenablog.com/entry/2023/09/06/010926)
