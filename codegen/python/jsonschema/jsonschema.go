package jsonschema

import (
	"encoding/json"
	"fmt"

	"github.com/Jumpscale/go-raml/raml"
)

type jsonSchema struct {
	T          raml.Type           `json:"-"`
	Name       string              `json:"-"`
	Type       string              `json:"type"`
	Properties map[string]property `json:"properties"`
}

func NewJSONSchema(t raml.Type, name string) jsonSchema {
	props := make(map[string]property, len(t.Properties))

	for k, v := range t.Properties {
		p := newProperty(raml.ToProperty(k, v))
		if !isPropTypeSupported(p.Type) {
			continue
		}
		props[p.Name] = p
	}

	typ := fmt.Sprintf("%v", t.Type)
	if typ == "" || t.Type == nil {
		typ = "object"
	}
	return jsonSchema{
		T:          t,
		Name:       name,
		Type:       typ,
		Properties: props,
	}
}

func (js jsonSchema) Supported() bool {
	return js.Type == "object"
}
func (js jsonSchema) String() string {
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

type property struct {
	Name     string `json:"-"`
	Type     string `json:"type,omitempty"`
	Required bool   `json:"required"`

	// string
	MinLength *int `json:"minLength,omitempty"`
	MaxLength *int `json:"maxLength,omitempty"`

	// number
	Minimum    *float64 `json:"minimum,omitempty"`
	Maximum    *float64 `json:"maximum,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty"`

	// array
	MinItems    *int  `json:"minItems,omitempty"`
	MaxItems    *int  `json:"maxItems,omitempty"`
	UniqueItems *bool `json:"uniqueItems,omitempty"`
}

func newProperty(rp raml.Property) property {
	return property{
		Name:      rp.Name,
		Type:      rp.Type,
		MinLength: rp.MinLength,
	}
}

var (
	supportedPropTypes = map[string]bool{
		"string":  true,
		"number":  true,
		"boolean": true,
	}
)

func isPropTypeSupported(typ string) bool {
	_, ok := supportedPropTypes[typ]
	return ok
}
