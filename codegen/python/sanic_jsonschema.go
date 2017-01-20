package python

import (
	"io/ioutil"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/chuckpreslar/inflect"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/python/jsonschema"
	"github.com/Jumpscale/go-raml/codegen/resource"
)

func (s SanicServer) generateJSONSchema(dir string) error {
	if err := s.genJSONSchemaFromTypes(dir); err != nil {
		return err
	}
	if err := s.genJSONSchemaFromMethods(dir); err != nil {
		return err
	}
	return nil
}

func (s SanicServer) genJSONSchemaFromTypes(dir string) error {
	for name, t := range s.APIDef.Types {
		js := jsonschema.NewJSONSchema(t, name)
		if err := s.genSchemaFile(js, dir); err != nil {
			return err
		}
	}
	return nil
}

func (s SanicServer) genJSONSchemaFromMethods(dir string) error {
	jsonSchemaFromMethod := func(m serverMethod) error {
		// TODO : merge this code with the flask version
		// to avoid code duplication
		// request body
		if commons.HasJSONBody(&m.Bodies) {
			name := inflect.UpperCamelCase(m.MethodName + "ReqBody")
			js := jsonschema.NewJSONSchemaFromProps(m.Bodies.ApplicationJSON.Properties, "object", name)
			if err := s.genSchemaFile(js, dir); err != nil {
				return err
			}
		}

		// response body
		for _, r := range m.Responses {
			if !commons.HasJSONBody(&r.Bodies) {
				continue
			}
			name := inflect.UpperCamelCase(m.MethodName + "RespBody")
			js := jsonschema.NewJSONSchemaFromProps(r.Bodies.ApplicationJSON.Properties, "object", name)
			if err := s.genSchemaFile(js, dir); err != nil {
				return err
			}
		}
		return nil
	}
	for _, rdi := range s.ResourcesDef {
		pr := newResourceFromDef(rdi.(resource.Resource), s.APIDef, newServerMethodSanic)
		for _, mi := range pr.Methods {
			m := mi.(serverMethod)
			if err := jsonSchemaFromMethod(m); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s SanicServer) genSchemaFile(js jsonschema.JSONSchema, dir string) error {
	fileName := filepath.Join(dir, js.Name+"_schema.json")
	log.Infof("generating file %v", fileName)
	return ioutil.WriteFile(fileName, js.String(), 0755)
}
