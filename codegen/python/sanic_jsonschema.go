package python

import (
	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/python/jsonschema"
	"github.com/Jumpscale/go-raml/codegen/types"
	"github.com/Jumpscale/go-raml/raml"
)

const (
	schemaDir = "schema"
)

func (s SanicServer) generateJSONSchema(dir string) error {
	sDir := filepath.Join(dir, schemaDir)
	commons.CheckCreateDir(sDir)
	for name, t := range types.AllTypes(s.APIDef, "") {
		switch tip := t.Type.(type) {
		case string:
		case types.TypeInBody:
			if tip.ReqResp == types.HTTPRequest {
				methodName := setServerMethodName(tip.Endpoint.Method.DisplayName, tip.Endpoint.Verb, tip.Endpoint.Resource)
				js := raml.NewJSONSchemaFromBodies(tip.Body(), setReqBodyName(methodName))
				if err := jsonschema.Generate(js, sDir); err != nil {
					return err
				}

			}
		case raml.Type:
			js := raml.NewJSONSchema(tip, name)
			if err := jsonschema.Generate(js, sDir); err != nil {
				return err
			}
		}
	}
	return nil
}
