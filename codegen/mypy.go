package codegen

import (
	"github.com/Jumpscale/go-raml/codegen/mypy"
	"github.com/Jumpscale/go-raml/raml"
)

func GenerateMyPy(apiDef *raml.APIDefinition, dir string) error {
	return mypy.GenerateMyPy(apiDef, dir)
}
