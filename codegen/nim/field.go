package nim

import (
	"github.com/Jumpscale/go-raml/raml"
)

// field represents a Nim object field
type field struct {
	Name string // field name
	Type string // field type
	Enum *enum
}

func newField(objName string, prop raml.Property) field {
	f := field{
		Name: prop.Name,
		Type: toNimType(prop.Type),
	}
	if prop.IsEnum() {
		f.Enum = newEnum(objName, prop, false)
		f.Type = f.Enum.Name
	}
	return f
}
