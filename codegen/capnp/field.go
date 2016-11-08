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

func newField(prop raml.Property, lang, pkg string) field {
	fd := field{
		Name: prop.Name,
		Type: toCapnpType(prop.Type, prop.CapnpType),
		Num:  prop.CapnpFieldNumber,
	}
	if isEnum(prop) {
		fd.Enum = newEnum(prop, lang, pkg)
		fd.Type = fd.Enum.Name
	}
	return fd
}
