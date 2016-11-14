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

func newEnum(structName string, prop raml.Property, pkg string, fromStruct bool) *enum {
	e := enum{
		Name: strings.Title(structName) + strings.Title(prop.Name),
		Type: convertToGoType(prop.Type),
		Pkg:  pkg,
	}
	if !fromStruct {
		e.Name = "Enum" + e.Name
	}
	for _, v := range prop.Enum.([]interface{}) {
		e.Fields = append(e.Fields, newEnumField(v, e))
	}
	return &e
}

func newEnumFromStruct(sd *structDef) *enum {
	prop := raml.Property{
		Type: fmt.Sprint(sd.T.Type),
		Name: "",
		Enum: sd.T.Enum,
	}
	return newEnum(sd.Name, prop, sd.PackageName, true)
}
func newEnumField(f interface{}, e enum) enumField {
	var val string

	name := fmt.Sprintf("%v%v", e.Name, f)
	switch v := f.(type) {
	case string:
		val = fmt.Sprintf(`"%v"`, v)
	case int:
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
