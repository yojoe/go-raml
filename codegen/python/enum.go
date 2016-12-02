package python

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

type enumField struct {
	Name  string
	Value string
}
type enum struct {
	Name   string
	Type   string
	Fields []enumField
}

func newEnum(name string, prop raml.Property, fromClass bool) *enum {
	enumName := strings.Title(name) + strings.Title(prop.Name)
	if !fromClass {
		enumName = "Enum" + enumName
	}
	e := enum{
		Name: enumName,
	}

	for _, v := range prop.Enum.([]interface{}) {
		e.Fields = append(e.Fields, newEnumField(v, e))
	}
	return &e
}

func newEnumFromClass(pc *class) *enum {
	prop := raml.Property{
		Type: fmt.Sprint(pc.T.Type),
		Name: "",
		Enum: pc.T.Enum,
	}
	return newEnum(pc.Name, prop, true)
}

func newEnumField(f interface{}, e enum) enumField {
	var name, val string

	switch v := f.(type) {
	case string:
		name = fmt.Sprintf("%v", v)
		val = fmt.Sprintf(`"%v"`, v)
	case int, float64, float32:
		name = fmt.Sprintf("e%v", v)
		val = fmt.Sprintf("%v", v)
	}
	return enumField{
		Name:  name,
		Value: val,
	}
}

func (e *enum) generate(dir string) error {
	filename := filepath.Join(dir, e.Name+".py")
	return commons.GenerateFile(e, "./templates/enum_python.tmpl", "enum_python", filename, true)
}
