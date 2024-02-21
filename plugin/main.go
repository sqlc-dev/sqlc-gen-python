package main

import (
	"github.com/sqlc-dev/plugin-sdk-go/codegen"

	python "github.com/sqlc-dev/sqlc-gen-python/internal"
)

func main() {
	codegen.Run(python.Generate)
}
