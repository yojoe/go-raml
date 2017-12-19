package codegen

import (
	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/codegen/python"
)

func GeneratePythonCapnp(apiDef *raml.APIDefinition, dir string) error {
	return python.GeneratePythonCapnpClasses(apiDef, dir)
}
