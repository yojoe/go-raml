package mypy

import (
	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/python"
	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/codegen/capnp"
)

func GenerateMyPy(apiDef *raml.APIDefinition, dir string) error {
	if err := capnp.GenerateCapnp(apiDef, dir, "", ""); err != nil {
		return err
	}

	client := python.Client{
		Name:    commons.NormalizeURI(apiDef.Title),
		APIDef:  apiDef,
		BaseURI: apiDef.BaseURI,
	}

	return client.GenerateMyPyClasses(dir)
}
