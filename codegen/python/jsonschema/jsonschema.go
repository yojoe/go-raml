package jsonschema

import (
	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

const (
	fileSuffix = "_schema.json"
)

// Generate generates a json file of this schema
func Generate(js raml.JSONSchema, dir string) error {
	filename := filepath.Join(dir, js.Name+fileSuffix)
	ctx := map[string]interface{}{
		"Content": js.String(),
	}
	return commons.GenerateFile(ctx, "./templates/json_schema.tmpl", "json_schema", filename, false)
}
