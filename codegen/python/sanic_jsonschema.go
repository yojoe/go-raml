package python

import (
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/types"
	"github.com/Jumpscale/go-raml/raml"
)

const (
	jsFileSuffix = "_schema.json" // json schema file suffix
)

// just a convenience type
type jsonSchema struct {
	*raml.JSONSchema
}

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
			rt := raml.Type{Type: tip}
			if rt.IsMultipleInheritance() {
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
		rt := raml.Type{Type: tip}
		if parents, isMult := rt.MultipleInheritance(); isMult {
			name := jsMultipleInheritanceName(parents)
			js := raml.NewJSONSchemaFromProps(nil, map[string]interface{}{}, "object", name)

			js.Inherit(getParentsObjs(parents))
			jsObjects[js.Name] = js
		}
	}

	for _, obj := range jsObjects {
		js := jsonSchema{
			JSONSchema: &obj,
		}
		js.HandleAdvancedType()
		if err := js.Generate(sDir); err != nil {
			return err
		}
	}
	return nil
}

// TODO : refactor it
// this func is ugly, it should be part of raml.JSONSchema class
// or we inherit that class
func (js *jsonSchema) HandleAdvancedType() {
	rt := raml.Type{Type: js.Type}
	parent, isSingleInherit := rt.SingleInheritance()
	switch {
	case rt.IsMultipleInheritance():
		parents, _ := rt.MultipleInheritance()
		js.Inherit(getParentsObjs(parents))
	case isSingleInherit:
		js.Inherit(getParentsObjs([]string{parent}))
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

// Generate generates a json file of this schema
func (js jsonSchema) Generate(dir string) error {
	filename := filepath.Join(dir, js.Name+jsFileSuffix)
	ctx := map[string]interface{}{
		"Content": js.String(),
	}
	return commons.GenerateFile(ctx, "./templates/json_schema.tmpl", "json_schema", filename, false)
}
