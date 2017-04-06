package python

import (
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/python/jsonschema"
	"github.com/Jumpscale/go-raml/codegen/types"
	"github.com/Jumpscale/go-raml/raml"
)

var (
	jsObjects map[string]raml.JSONSchema
)

func init() {
	jsObjects = map[string]raml.JSONSchema{}
}

func (s SanicServer) schemaDir() string {
	return "schema"
}
func (s SanicServer) generateJSONSchema(dir string) error {

	// array of tip that need to be generated in the end of this
	// process. because it needs other object to be registered first
	delayedMI := []string{} // delayed multiple inheritance

	sDir := filepath.Join(dir, s.schemaDir())

	commons.CheckCreateDir(sDir)

	for name, t := range types.AllTypes(s.APIDef, "") {
		switch tip := t.Type.(type) {
		case string:
			if commons.IsMultipleInheritance(tip) {
				delayedMI = append(delayedMI, tip)
			}
		case types.TypeInBody:
			if tip.ReqResp == types.HTTPRequest {
				methodName := setServerMethodName(tip.Endpoint.Method.DisplayName, tip.Endpoint.Verb, tip.Endpoint.Resource)
				js := raml.NewJSONSchemaFromProps(nil, tip.Properties, "object", setReqBodyName(methodName))
				jsObjects[js.Name] = js
			}
		case raml.Type:
			js := raml.NewJSONSchema(tip, name)
			jsObjects[js.Name] = js
		}
	}

	for _, tip := range delayedMI {
		parents, _ := commons.MultipleInheritance(tip)

		name := jsMultipleInheritanceName(parents)
		js := raml.NewJSONSchemaFromProps(nil, map[string]interface{}{}, "object", name)

		js.Inherit(getParentsObjs(parents))
		jsObjects[js.Name] = js
	}

	for _, js := range jsObjects {
		jsHandleAdvancedType(&js)
		if err := jsonschema.Generate(js, sDir); err != nil {
			return err
		}
	}
	return nil
}

// TODO : refactor it
// this func is ugly, it should be part of raml.JSONSchema class
// or we inherit that class
func jsHandleAdvancedType(js *raml.JSONSchema) {
	parents, isMult := commons.MultipleInheritance(js.Type)
	switch {
	case isMult:
		js.Inherit(getParentsObjs(parents))
	}
}

// get JSON schema objects from array of JSON schema name
func getParentsObjs(parents []string) []raml.JSONSchema {
	objs := []raml.JSONSchema{}
	for _, p := range parents {
		if v, ok := jsObjects[p]; ok {
			objs = append(objs, v)
		}
	}
	return objs
}

func jsMultipleInheritanceName(parents []string) string {
	return strings.Join(parents, "")
}
