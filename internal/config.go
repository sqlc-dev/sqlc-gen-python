package python

type Config struct {
	EmitModule                  bool     `json:"emit_module"`     // If true emits functions in module, else wraps in a class.
	EmitGenerators              bool     `json:"emit_generators"` // Will we use generators or lists, defaults to true
	EmitAsync                   bool     `json:"emit_async"`      // Emits async code instead of sync
	EmitExactTableNames         bool     `json:"emit_exact_table_names"`
	EmitSyncQuerier             bool     `json:"emit_sync_querier"`  // DEPRECATED ALIAS FOR: emit_type = 'class', emit_generators = True
	EmitAsyncQuerier            bool     `json:"emit_async_querier"` // DEPRECATED ALIAS FOR: emit_type = 'class', emit_generators = True
	Package                     string   `json:"package"`
	Out                         string   `json:"out"`
	EmitPydanticModels          bool     `json:"emit_pydantic_models"`
	QueryParameterLimit         *int32   `json:"query_parameter_limit"`
	InflectionExcludeTableNames []string `json:"inflection_exclude_table_names"`
}
