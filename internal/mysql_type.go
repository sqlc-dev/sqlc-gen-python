package python

import (
	"log"

	"github.com/sqlc-dev/plugin-sdk-go/plugin"
	"github.com/sqlc-dev/plugin-sdk-go/sdk"
)

func mysqlType(req *plugin.GenerateRequest, col *plugin.Column) string {
	columnType := sdk.DataType(col.Type)

	switch columnType {
	case "varchar", "text", "char", "tinytext", "mediumtext", "longtext":
		return "str"
	case "tinyint":
		if col.Length == 1 {
			return "bool"
		} else {
			return "int"
		}
	case "int", "integer", "smallint", "mediumint", "year":
		return "int"
	case "bigint":
		return "int"
	case "blob", "binary", "varbinary", "tinyblob", "mediumblob", "longblob":
		return "Any"
	case "double", "double precision", "real", "float":
		return "float"
	case "decimal", "dec", "fixed":
		return "string"
	case "enum":
		return "string"
	case "date", "timestamp", "datetime", "time":
		return "datetime.date"
	case "boolean", "bool":
		return "bool"
	case "json":
		return "Any"
	case "any":
		return "Any"
	default:
		for _, schema := range req.Catalog.Schemas {
			for _, enum := range schema.Enums {
				if columnType == enum.Name {
					if schema.Name == req.Catalog.DefaultSchema {
						return "models." + modelName(enum.Name, req.Settings)
					}
					return "models." + modelName(schema.Name+"_"+enum.Name, req.Settings)
				}
			}
		}
		log.Printf("unknown MySQL type: %s\n", columnType)
		return "Any"
	}
}
