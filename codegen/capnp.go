package codegen

import (
	"github.com/Jumpscale/go-raml/codegen/capnp"
	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

func GenerateCapnp(apiDef *raml.APIDefinition, dir, lang, pkg string) error {
	// create directory if needed
	if err := commons.CheckCreateDir(dir); err != nil {
		return err
	}

	// generate types
	for name, tipe := range apiDef.Types {
		s, err := capnp.NewStruct(tipe, name, lang, pkg)
		if err != nil {
			return err
		}
		if err := s.Generate(dir); err != nil {
			return err
		}
	}
	return nil
}
