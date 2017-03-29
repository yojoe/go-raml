package python

import (
	"path/filepath"

	"github.com/chuckpreslar/inflect"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/python/jsonschema"
	"github.com/Jumpscale/go-raml/raml"
)

const (
	schemaDir = "schema"
)

func (s SanicServer) generateJSONSchema(dir string) error {
	sDir := filepath.Join(dir, schemaDir)
	commons.CheckCreateDir(sDir)
	if err := s.genJSONSchemaFromTypes(sDir); err != nil {
		return err
	}
	if err := s.genJSONSchemaFromMethods(sDir); err != nil {
		return err
	}
	return nil
}

func (s SanicServer) genJSONSchemaFromTypes(dir string) error {
	for name, t := range s.APIDef.Types {
		js := raml.NewJSONSchema(t, name)
		if err := jsonschema.Generate(js, dir); err != nil {
			return err
		}
	}
	return nil
}

func (s SanicServer) genJSONSchemaFromMethods(dir string) error {
	// creates JSON schema from an raml method
	jsonSchemaFromMethod := func(m serverMethod) error {
		if commons.HasJSONBody(&m.Bodies) {
			name := inflect.UpperCamelCase(m.MethodName + "ReqBody")
			js := raml.NewJSONSchemaFromBodies(m.Bodies, name)
			if err := jsonschema.Generate(js, dir); err != nil {
				return err
			}
		}

		// no need to generate schema for response body
		// because we never need to validate it first
		return nil
	}
	for _, pr := range s.ResourcesDef {
		for _, mi := range pr.Methods {
			m := mi.(serverMethod)
			if err := jsonSchemaFromMethod(m); err != nil {
				return err
			}
		}
	}
	return nil
}
