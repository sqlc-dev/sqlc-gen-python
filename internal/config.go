package python

// TODO: Where are these properly documented?
type Config struct {
	EmitAsync                   bool     `json:"emit_async"` // Emits async code instead of sync
	EmitExactTableNames         bool     `json:"emit_exact_table_names"`
	EmitGenerators              bool     `json:"emit_generators"` // Will we use generators or lists, defaults to false
	EmitModule                  bool     `json:"emit_module"`     // If true emits functions in module, else wraps in a class.
	EmitPydanticModels          bool     `json:"emit_pydantic_models"`
	EmitSyncQuerier             bool     `json:"emit_sync_querier"`  // DEPRECATED ALIAS FOR: emit_type = 'class', emit_generators = True
	EmitAsyncQuerier            bool     `json:"emit_async_querier"` // DEPRECATED ALIAS FOR: emit_type = 'class', emit_generators = True
	InflectionExcludeTableNames []string `json:"inflection_exclude_table_names"`
	Out                         string   `json:"out"`
	OutputModelsFileName        *string  `json:"output_models_file_name,omitempty"` // Can be string or null to exclude generating models file
	OutputQuerierFile           bool     `json:"output_querier_file,omitempty"`     // Skips outputting queries
	Package                     string   `json:"package"`
	QueryParameterLimit         *int32   `json:"query_parameter_limit"`
}
