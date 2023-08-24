# README
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
      emit_module: false
      emit_generators: true
      emit_async: false
```

## Multiple packages
If you have have a mono-repository setup you may want to split the output of queries and models. You can achieve this by using the `output_models_file_name`
and `output_querier_file` fields. If `output_models_file_name` is set to `null` for it will not output the `models.py` file. Setting `output_querier_file` to false will prevent outputting any query files. Combining these you can set one codegen to only output models while the other codegen outputs only queries. Make sure the `package` configuration is set equally so the query files import correctly the models.

SQLC needs at least one query, so you may need to add a unused query like the following in your schema and reuse the `schema` as `queries`.
```sql
-- name: Skip :one
SELECT 1;
```