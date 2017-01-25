package jsonschema

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

var (
	scalarTypes = map[string]bool{
		"string":  true,
		"number":  true,
		"boolean": true,
		"integer": true,
	}
)

type JSONSchema struct {
	Schema     string              `json:"$schema"`
	Name       string              `json:"-"`
	Type       string              `json:"type"`
	Properties map[string]property `json:"properties"`
	Required   []string            `json:"required,omitempty"`
}

// NewJSONSchema creates JSON schema from an raml type
func NewJSONSchema(t raml.Type, name string) JSONSchema {
	typ := fmt.Sprintf("%v", t.Type)
	if typ == "" || t.Type == nil {
		typ = "object"
	}

	return newJSONSchemaFromProps(t.Properties, typ, name)
}

// NewJSONSchemaFromBodies creates JSON schema from raml bodies
func NewJSONSchemaFromBodies(b raml.Bodies, name string) JSONSchema {
	if b.ApplicationJSON.Type != "" {
		var t raml.Type
		if err := json.Unmarshal([]byte(b.ApplicationJSON.Type), &t); err == nil {
			return NewJSONSchema(t, name)
		}
	}
	return newJSONSchemaFromProps(b.ApplicationJSON.Properties, "object", name)
}

// newJSONSchemaFromProps creates json schmema
// from a map of properties
func newJSONSchemaFromProps(properties map[string]interface{}, typ, name string) JSONSchema {
	var required []string

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
		Schema:     "http://json-schema.org/schema#",
		Name:       name,
		Type:       typ,
		Properties: props,
		Required:   required,
	}
}

func (js JSONSchema) Supported() bool {
	return js.Type == "object"
}
func (js JSONSchema) String() string {
	// force the type to `object` type
	// we do it because we can only support `object` type now
	// TODO: fix it
	if js.Type != "object" {
		js.Type = "object"
	}

	b, err := json.MarshalIndent(&js, "", "\t")
	if err != nil {
		return "{}"
	}
	return string(b)
}

func (js JSONSchema) Generate(dir string) error {
	filename := filepath.Join(dir, js.Name+"_schema.json")
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

type arrayItem struct {
	Type string `json:"type,omitempty"`
	Ref  string `json:"$ref,omitempty"`
}

func newArrayItem(typ string) *arrayItem {
	if _, isScalar := scalarTypes[typ]; isScalar {
		return &arrayItem{
			Type: typ,
		}
	}
	return &arrayItem{
		Ref: typ + "_schema.json",
	}
}

func newProperty(rp raml.Property) property {
	_, isScalar := scalarTypes[rp.Type]
	if rp.Type != "" && !isScalar && !rp.IsArray() && !rp.IsBidimensiArray() {
		return property{
			Name:     rp.Name,
			Ref:      rp.Type + "_schema.json",
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

	if rp.IsArray() && !rp.IsBidimensiArray() && isPropTypeSupported(rp) {
		p.Type = "array"
		p.Items = newArrayItem(rp.ArrayType())
	}
	return p
}

func isPropTypeSupported(p raml.Property) bool {
	return !p.IsBidimensiArray()
}
