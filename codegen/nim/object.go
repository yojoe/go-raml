package nim

import (
	"encoding/json"
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

// GenerateObjects generates Nim objects from RAML types
func GenerateObjects(types map[string]raml.Type, dir string) error {
	for name, t := range types {
		if err := generateObject(t, name, dir); err != nil {
			fmt.Printf("failed : %v\n", err)
		}
	}
	return nil
}

// GenerateObjectFromBody generate a Nim object from an RAML Body
func GenerateObjectFromBody(methodName string, body *raml.Bodies, isReq bool, dir string) error {
	if !commons.HasJSONBody(body) {
		return nil
	}
	obj, err := newObjectFromBody(methodName, body, isReq)
	if err != nil {
		return err
	}
	return obj.generate(dir)
}

// generates Nim object from an RAML type
func generateObject(t raml.Type, name, dir string) error {
	obj, err := newObjectFromType(t, name)
	if err != nil {
		return err
	}
	return obj.generate(dir)
}

func newObjectFromBody(methodName string, body *raml.Bodies, isReq bool) (object, error) {
	name := methodName + "RespBody"
	if isReq {
		name = methodName + "ReqBody"
	}

	if body.ApplicationJSON.Type != "" {
		var t raml.Type
		if err := json.Unmarshal([]byte(body.ApplicationJSON.Type), &t); err == nil {
			return newObjectFromType(t, name)
		}
	}

	return newObject(name, "", body.ApplicationJSON.Properties)
}

func newObjectFromType(t raml.Type, name string) (object, error) {
	obj, err := newObject(name, t.Description, t.Properties)
	obj.T = t
	return obj, err
}

func newObject(name, description string, properties map[string]interface{}) (object, error) {
	// generate fields from type properties
	fields := make(map[string]field)

	for k, v := range properties {
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
		Description: commons.ParseDescription(description),
	}, nil
}

// generate nim object representation
func (o *object) generate(dir string) error {
	filename := filepath.Join(dir, o.Name+".nim")
	return commons.GenerateFile(o, "./templates/object_nim.tmpl", "object_nim", filename, true)
}
