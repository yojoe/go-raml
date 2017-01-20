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

func NewJSONSchema(t raml.Type, name string) JSONSchema {
	typ := fmt.Sprintf("%v", t.Type)
	if typ == "" || t.Type == nil {
		typ = "object"
	}

	return NewJSONSchemaFromProps(t.Properties, typ, name)
}

func NewJSONSchemaFromProps(properties map[string]interface{}, typ, name string) JSONSchema {
	var required []string

	props := make(map[string]property, len(properties))
	for k, v := range properties {
		p := newProperty(raml.ToProperty(k, v))
		if !isPropTypeSupported(p.Type) {
			continue
		}
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
	MinItems    *int `json:"minItems,omitempty"`
	MaxItems    *int `json:"maxItems,omitempty"`
	UniqueItems bool `json:"uniqueItems,omitempty"`
}

func newProperty(rp raml.Property) property {
	return property{
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
}

var (
	supportedPropTypes = map[string]bool{
		"string":  true,
		"number":  true,
		"boolean": true,
		"integer": true,
	}
)

func isPropTypeSupported(typ string) bool {
	_, ok := supportedPropTypes[typ]
	return ok
}
