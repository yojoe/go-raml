package codegen

import (
	"github.com/Jumpscale/go-raml/codegen/python"
	"github.com/Jumpscale/go-raml/raml"
)

func GeneratePythonCapnp(apiDef *raml.APIDefinition, dir string) error {
	return python.GeneratePythonCapnpClasses(apiDef, dir)
}
