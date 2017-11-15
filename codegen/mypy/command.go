package mypy

import (
	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/python"
	"github.com/Jumpscale/go-raml/raml"
)

func GenerateMyPy(apiDef *raml.APIDefinition, dir string) error {
	client := python.Client{
		Name:    commons.NormalizeURI(apiDef.Title),
		APIDef:  apiDef,
		BaseURI: apiDef.BaseURI,
	}

	return client.GenerateClasses(dir)
}
