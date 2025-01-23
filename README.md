## Usage

```yaml
version: "2"
plugins:
  - name: py
    wasm:
      url: https://downloads.sqlc.dev/plugin/sqlc-gen-python_1.3.0.wasm
      sha256: fbedae96b5ecae2380a70fb5b925fd4bff58a6cfb1f3140375d098fbab7b3a3c
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

### Emit Pydantic Models instead of `dataclasses`

Option: `emit_pydantic_models`

By default, `sqlc-gen-python` will emit `dataclasses` for the models. If you prefer to use [`pydantic`](https://docs.pydantic.dev/latest/) models, you can enable this option.

with `emit_pydantic_models`

```py
from pydantic import BaseModel

class Author(pydantic.BaseModel):
    id: int
    name: str
```

without `emit_pydantic_models`

```py
import dataclasses

@dataclasses.dataclass()
class Author:
    id: int
    name: str
```

### Use `enum.StrEnum` for Enums

Option: `emit_str_enum`

`enum.StrEnum` was introduce in Python 3.11.

`enum.StrEnum` is a subclass of `str` that is also a subclass of `Enum`. This allows for the use of `Enum` values as strings, compared to strings, or compared to other `enum.StrEnum` types.

This is convenient for type checking and validation, as well as for serialization and deserialization.

By default, `sqlc-gen-python` will emit `(str, enum.Enum)` for the enum classes. If you prefer to use `enum.StrEnum`, you can enable this option.

with `emit_str_enum`

```py
class Status(enum.StrEnum):
    """Venues can be either open or closed"""
    OPEN = "op!en"
    CLOSED = "clo@sed"
```

without `emit_str_enum` (current behavior)

```py
class Status(str, enum.Enum):
    """Venues can be either open or closed"""
    OPEN = "op!en"
    CLOSED = "clo@sed"
```
