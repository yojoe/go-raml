package nim

import (
	"fmt"
	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

// object represents a Nim object
type object struct {
	Name        string
	Description []string
	Fields      map[string]field
	T           raml.Type
}

// field represents a Nim object field
type field struct {
	Name string // field name
	Type string // field type
}

// GenerateObject generates Nim objects from RAML types
func GenerateObjects(types map[string]raml.Type, dir string) error {
	for name, t := range types {
		if err := generateObject(t, name, dir); err != nil {
			fmt.Printf("failed : %v\n", err)
		}
	}
	return nil
}

func generateObject(t raml.Type, name, dir string) error {
	obj, err := newObject(t, name)
	if err != nil {
		return err
	}
	return obj.generate(dir)
}

func newObject(t raml.Type, name string) (object, error) {
	// generate fields from type properties
	fields := make(map[string]field)

	for k, v := range t.Properties {
		prop := raml.ToProperty(k, v)
		fd := field{
			Name: prop.Name,
			Type: toNimType(prop.Type),
		}
		if fd.Type == "" {
			return object{}, fmt.Errorf("unsupported type in nim:%v", prop.Type)
		}
		fields[prop.Name] = fd
	}

	return object{
		Name:        name,
		Fields:      fields,
		Description: commons.ParseDescription(t.Description),
		T:           t,
	}, nil
}

// generate nim object representation
func (o *object) generate(dir string) error {
	filename := filepath.Join(dir, o.Name+".nim")
	return commons.GenerateFile(o, "./templates/object_nim.tmpl", "object_nim", filename, true)
}
