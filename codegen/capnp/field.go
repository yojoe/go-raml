package capnp

import (
	"github.com/Jumpscale/go-raml/raml"
)

type field struct {
	Name string
	Type string
	Num  int
	Enum *enum
	Items string
}

func newField(structName string, prop raml.Property, lang, pkg string) field {
	capnpType, items :=  toCapnpType(prop.TypeString(), prop.CapnpType, prop.Items.Type)
	fd := field{
		Name: prop.Name,
		Type: capnpType,
		Items: items,

	}
	if isEnum(prop) {
		fd.Enum = newEnum(structName, prop, lang, pkg)
		fd.Type = fd.Enum.Name
	}
	return fd
}
