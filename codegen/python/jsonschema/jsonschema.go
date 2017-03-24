package jsonschema

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

const (
	schemaVer  = "http://json-schema.org/schema#"
	fileSuffix = "_schema.json"
)

var (
	scalarTypes = map[string]bool{
		"string":  true,
		"number":  true,
		"boolean": true,
		"integer": true,
	}
)

// JSONSchema represents a json-schema json file
type JSONSchema struct {
	Schema     string              `json:"$schema"`
	Name       string              `json:"-"`
	Type       string              `json:"type"`
	Items      *arrayItem          `json:"items,omitempty"`
	Properties map[string]property `json:"properties,omitempty"`
	Required   []string            `json:"required,omitempty"`

	// Array properties
	MinItems    int  `json:"minItems,omitempty"`
	MaxItems    int  `json:"maxItems,omitempty"`
	UniqueItems bool `json:"uniqueItems,omitempty"`
}

// NewJSONSchema creates JSON schema from an raml type
func NewJSONSchema(t raml.Type, name string) JSONSchema {
	typ := fmt.Sprintf("%v", t.Type)
	if typ == "" || t.Type == nil {
		typ = "object"
	}

	return newJSONSchemaFromProps(&t, t.Properties, typ, name)
}

// NewJSONSchemaFromBodies creates JSON schema from raml bodies
func NewJSONSchemaFromBodies(b raml.Bodies, name string) JSONSchema {
	if b.ApplicationJSON.Type != "" {
		var t raml.Type
		if err := json.Unmarshal([]byte(b.ApplicationJSON.TypeString()), &t); err == nil {
			return NewJSONSchema(t, name)
		}
	}
	return newJSONSchemaFromProps(nil, b.ApplicationJSON.Properties, "object", name)
}

// newJSONSchemaFromProps creates json schmema
// from a map of properties
func newJSONSchemaFromProps(t *raml.Type, properties map[string]interface{}, typ, name string) JSONSchema {
	var required []string

	if isTypeArray(typ) {
		return newArraySchema(t, typ, name)
	}

	props := make(map[string]property, len(properties))
	for k, v := range properties {
		rp := raml.ToProperty(k, v)
		if !isPropTypeSupported(rp) {
			continue
		}

		p := newProperty(rp)
		props[p.Name] = p
		if p.Required {
			required = append(required, p.Name)
		}
	}

	// we need it to be sorted for testing purpose
	sort.Strings(required)
	return JSONSchema{
		Schema:     schemaVer,
		Name:       name,
		Type:       typ,
		Properties: props,
		Required:   required,
	}
}

// Supported returns true if the type is supported
func (js JSONSchema) Supported() bool {
	return js.Type == "object" || js.Type == "array"
}
func (js JSONSchema) String() string {
	// for unsupported type, force the type to `object` type
	if !js.Supported() {
		js.Type = "object"
	}

	b, err := json.MarshalIndent(&js, "", "\t")
	if err != nil {
		return "{}"
	}
	return string(b)
}

// Generate generates a json file of this schema
func (js JSONSchema) Generate(dir string) error {
	filename := filepath.Join(dir, js.Name+fileSuffix)
	ctx := map[string]interface{}{
		"Content": js.String(),
	}
	return commons.GenerateFile(ctx, "./templates/json_schema.tmpl", "json_schema", filename, false)
}

type property struct {
	Name     string      `json:"-"`
	Ref      string      `json:"$ref,omitempty"`
	Type     string      `json:"type,omitempty"`
	Required bool        `json:"-"`
	Enum     interface{} `json:"enum,omitempty"`

	// string
	MinLength *int    `json:"minLength,omitempty"`
	MaxLength *int    `json:"maxLength,omitempty"`
	Pattern   *string `json:"pattern,omitempty"`

	// number
	Minimum    *float64 `json:"minimum,omitempty"`
	Maximum    *float64 `json:"maximum,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty"`

	// array
	MinItems    *int       `json:"minItems,omitempty"`
	MaxItems    *int       `json:"maxItems,omitempty"`
	UniqueItems bool       `json:"uniqueItems,omitempty"`
	Items       *arrayItem `json:"items,omitempty"`
}

func newProperty(rp raml.Property) property {
	_, isScalar := scalarTypes[rp.Type]

	// complex type
	if rp.Type != "" && !isScalar && !rp.IsArray() && !rp.IsBidimensiArray() {
		return property{
			Name:     rp.Name,
			Ref:      rp.Type + fileSuffix,
			Required: rp.Required,
		}
	}

	p := property{
		Name:        rp.Name,
		Type:        rp.Type,
		Required:    rp.Required,
		Enum:        rp.Enum,
		MinLength:   rp.MinLength,
		MaxLength:   rp.MaxLength,
		Pattern:     rp.Pattern,
		Minimum:     rp.Minimum,
		Maximum:     rp.Maximum,
		MultipleOf:  rp.MultipleOf,
		MinItems:    rp.MinItems,
		MaxItems:    rp.MaxItems,
		UniqueItems: rp.UniqueItems,
	}

	// array
	if rp.IsArray() && !rp.IsBidimensiArray() {
		p.Type = "array"
		p.Items = newArrayItem(rp.ArrayType())
	}
	return p
}

func isPropTypeSupported(p raml.Property) bool {
	return !p.IsBidimensiArray() && !p.IsUnion()
}
