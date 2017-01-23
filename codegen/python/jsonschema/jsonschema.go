package jsonschema

import (
	"encoding/json"
	"fmt"

	"github.com/Jumpscale/go-raml/raml"
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
func (js JSONSchema) String() []byte {
	// force the type to `object` type
	// we do it because we can only support `object` type now
	// TODO: fix it
	if js.Type != "object" {
		js.Type = "object"
	}

	b, err := json.MarshalIndent(&js, "", "\t")
	if err != nil {
		return []byte("{}")
	}
	return b
}

type property struct {
	Name     string      `json:"-"`
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
	Type string `json:"type"`
}

func newProperty(rp raml.Property) property {
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
		p.Items = &arrayItem{
			Type: rp.ArrayType(),
		}
	}
	return p
}

var (
	supportedPropTypes = map[string]bool{
		"string":  true,
		"number":  true,
		"boolean": true,
		"integer": true,
	}
)

func isPropTypeSupported(p raml.Property) bool {
	typ := p.Type
	if p.IsArray() && !p.IsBidimensiArray() {
		typ = p.ArrayType()
	}
	_, ok := supportedPropTypes[typ]
	return ok
}
