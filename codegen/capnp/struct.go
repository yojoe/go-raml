package capnp

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

type field struct {
	Name string
	Type string
	Num  int
}

type Struct struct {
	ID            string
	Name          string
	Fields        map[string]field
	Description   []string
	OrderedFields []field
	T             raml.Type
	lang          string
	pkg           string
}

func NewStruct(t raml.Type, name, lang, pkg string) (Struct, error) {
	// generate fields from type properties
	fields := make(map[string]field)

	for k, v := range t.Properties {
		prop := raml.ToProperty(k, v)
		fd := field{
			Name: prop.Name,
			Type: toCapnpType(prop.Type, prop.CapnpType),
			Num:  prop.CapnpFieldNumber,
		}
		fields[prop.Name] = fd
	}

	ID, err := getID()
	if err != nil {
		return Struct{}, err
	}
	s := Struct{
		Name:        name,
		Fields:      fields,
		Description: commons.ParseDescription(t.Description),
		T:           t,
		ID:          ID,
		pkg:         pkg,
		lang:        lang,
	}
	if err := s.checkValidCapnp(); err != nil {
		return s, err
	}
	return s, s.orderFields()
}

func (s *Struct) Generate(dir string) error {
	filename := filepath.Join(dir, s.Name+".capnp")
	return commons.GenerateFile(s, "./templates/struct_capnp.tmpl", "struct_capnp", filename, true)

}

func (s *Struct) ImportLang() string {
	switch s.lang {
	case "go":
		return `using Go = import "/go.capnp";`
	default:
		return ""
	}
}

func (s *Struct) Annotations() []string {
	switch s.lang {
	case "go":
		return s.goAnnotations()
	default:
		return []string{}
	}
}

func (s *Struct) checkValidCapnp() error {
	if strings.Title(s.Name) != s.Name {
		return fmt.Errorf("invalid type name:%v. Type names must begin with a capital letter", s.Name)
	}
	return nil
}

func (s *Struct) goAnnotations() []string {
	pkg := fmt.Sprintf(`$Go.package("%v")`, s.pkg)
	return []string{pkg}
}

func (s *Struct) orderFields() error {
	findField := func(num int) (field, bool) {
		for _, f := range s.Fields {
			if f.Num == num {
				return f, true
			}
		}
		return field{}, false
	}

	for i := 0; i < len(s.Fields); i++ {
		f, ok := findField(i)
		if !ok {
			return fmt.Errorf("can't find field number %v of `%v`", i, s.Name)
		}
		s.OrderedFields = append(s.OrderedFields, f)
	}
	return nil
}
