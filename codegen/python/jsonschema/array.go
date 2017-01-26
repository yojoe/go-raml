package jsonschema

import (
	"fmt"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

func newArraySchema(t *raml.Type, typ, name string) JSONSchema {
	js := JSONSchema{
		Name:   name,
		Schema: schemaVer,
		Type:   "array",
		Items:  newArrayItem(getArrayItemsType(t, typ)),
	}
	if t == nil {
		return js
	}
	js.MaxItems = t.MaxItems
	js.MinItems = t.MinItems
	js.UniqueItems = t.UniqueItems
	return js
}

func isTypeArray(typ string) bool {
	return typ == "array" || (commons.IsArray(typ) && !commons.IsBidimensiArray(typ))
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
		Ref: typ + fileSuffix,
	}
}

func getArrayItemsType(t *raml.Type, typ string) string {
	if typ == "array" {
		return fmt.Sprint(t.Items)
	}
	return commons.ArrayType(typ)
}
