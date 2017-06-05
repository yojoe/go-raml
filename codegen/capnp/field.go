package capnp

import (
	"github.com/Jumpscale/go-raml/raml"
)

type field struct {
	Name string
	Type string
	Num  int
	Enum *enum
}

func newField(structName string, prop raml.Property, lang, pkg string) field {
	fd := field{
		Name: prop.Name,
		Type: toCapnpType(prop.TypeString(), prop.CapnpType),
		Num:  prop.CapnpFieldNumber,
	}
	if isEnum(prop) {
		fd.Enum = newEnum(structName, prop, lang, pkg)
		fd.Type = fd.Enum.Name
	}
	return fd
}
