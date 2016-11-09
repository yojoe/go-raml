package golang

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

type enumField struct {
	Name  string
	Type  string
	Value string
}
type enum struct {
	Name   string
	Type   string
	Fields []enumField
	Pkg    string
}

func newEnum(structName string, prop raml.Property, pkg string) *enum {
	e := enum{
		Name: "Enum" + strings.Title(prop.Name) + strings.Title(structName),
		Type: convertToGoType(prop.Type),
		Pkg:  pkg,
	}
	for _, v := range prop.Enum.([]interface{}) {
		e.Fields = append(e.Fields, newEnumField(v, e))
	}
	return &e
}

func newEnumField(f interface{}, e enum) enumField {
	var val string
	var name string
	switch v := f.(type) {
	case string:
		name = v
		val = fmt.Sprintf(`"%v"`, v)
	case int:
		name = fmt.Sprintf("%v_%v", e.Name, v)
		val = fmt.Sprintf("%v", v)
	}
	return enumField{
		Name:  name,
		Value: val,
		Type:  e.Name,
	}
}

func (e *enum) generate(dir string) error {
	filename := filepath.Join(dir, e.Name+".go")
	return commons.GenerateFile(e, "./templates/enum_go.tmpl", "enum_go", filename, true)

}

func isEnum(prop raml.Property) bool {
	return prop.Enum != nil
}
