package capnp

import (
	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

func GenerateCapnp(apiDef *raml.APIDefinition, dir, lang, pkg string) error {
	// create directory if needed
	if err := commons.CheckCreateDir(dir); err != nil {
		return err
	}

	structs := []Struct{}

	// generate types
	for name, tipe := range apiDef.Types {
		s, err := NewStruct(tipe, name, lang, pkg)
		if err != nil {
			return err
		}
		structs = append(structs, s)
	}
	for _, s := range structs {
		registerType(s.Name)
	}

	for _, s := range structs {
		if err := s.Generate(dir); err != nil {
			return err
		}
	}
	return nil
}
