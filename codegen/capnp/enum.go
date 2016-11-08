package capnp

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

type enum struct {
	ID     string
	Name   string
	Fields []field
	lang   string
	pkg    string
}

func newEnum(prop raml.Property, lang, pkg string) *enum {
	e := enum{
		ID:   getID(),
		Name: "Enum" + strings.Title(prop.Name),
		lang: lang,
		pkg:  pkg,
	}
	for k, v := range prop.Enum.([]interface{}) {
		f := field{
			Name: v.(string),
			Num:  k,
		}
		e.Fields = append(e.Fields, f)
	}
	return &e
}

func (e *enum) generate(dir string) error {
	filename := filepath.Join(dir, e.Name+".capnp")
	return commons.GenerateFile(e, "./templates/enum_capnp.tmpl", "enum_capnp", filename, true)
}

func (e *enum) Imports() string {
	switch e.lang {
	case "go":
		return `using Go = import "/go.capnp";`
	default:
		return ""
	}

}

func (e *enum) Annotations() []string {
	switch e.lang {
	case "go":
		return e.goAnnotations()
	default:
		return []string{}
	}
}

func (e *enum) goAnnotations() []string {
	return []string{
		fmt.Sprintf(`$Go.package("%v")`, e.pkg),
		fmt.Sprintf(`$Go.import("%v")`, e.pkg),
	}
}
func isEnum(prop raml.Property) bool {
	return prop.Enum != nil
}
