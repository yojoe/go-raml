package nim

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

var (
	// saves all generated objects
	objectsRegister map[string]struct{}
)

func init() {
	// register of all objects.
	// needed to generate `import`
	objectsRegister = map[string]struct{}{}
}

// object represents a Nim object
type object struct {
	Name        string
	Description []string
	Fields      map[string]field
	T           raml.Type
	OneLineDef  string
	Parents     []string
}

// generates Nim objects from RAML types
func generateObjects(types map[string]raml.Type, dir string) error {
	objs := []object{}
	for name, t := range types {
		obj, err := newObjectFromType(t, name)
		if err != nil {
			return err
		}
		objs = append(objs, obj)
	}

	for _, obj := range objs {
		registerObject(obj.Name)
		for _, f := range obj.Fields {
			if f.Enum != nil {
				registerObject(f.Enum.Name)
			}
		}
	}

	for _, obj := range objs {
		if err := obj.generate(dir); err != nil {
			return err
		}
	}
	return nil
}

// generate objects from method request & response bodies of all resources
func generateObjectsFromBodies(rs []resource, dir string) ([]string, error) {
	names := []string{}
	for _, r := range rs {
		for _, mi := range r.Methods {
			m := mi.(method)
			ns, err := generateObjectFromMethod(r, m, dir)
			if err != nil {
				fmt.Printf("failed : %v\n", err) // TODO : return err if failed
			}
			names = append(names, ns...)
		}
	}
	return names, nil
}

// generate object from a method
func generateObjectFromMethod(r resource, m method, dir string) ([]string, error) {
	names := []string{}

	name, err := generateObjectFromBody(m.MethodName, &m.Bodies, true, dir)
	if err != nil {
		return names, err
	}
	names = append(names, name)

	for _, v := range m.Responses {
		name, err := generateObjectFromBody(m.MethodName, &v.Bodies, false, dir)
		if err != nil {
			return names, err
		}
		names = append(names, name)
	}
	return names, nil
}

// generateObjectFromBody generate a Nim object from an RAML Body
func generateObjectFromBody(methodName string, body *raml.Bodies, isReq bool, dir string) (string, error) {
	if !commons.HasJSONBody(body) {
		return "", nil
	}
	obj, err := newObjectFromBody(methodName, body, isReq)
	if err != nil {
		return "", err
	}
	return obj.Name, obj.generate(dir)
}

// create new object from a method body
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

// create new object from an RAML type
func newObjectFromType(t raml.Type, name string) (object, error) {
	obj, err := newObject(name, t.Description, t.Properties)
	obj.T = t
	obj.handleAdvancedType()
	return obj, err
}

func newObject(name, description string, properties map[string]interface{}) (object, error) {
	// generate fields from type properties
	fields := make(map[string]field)

	for k, v := range properties {
		prop := raml.ToProperty(k, v)
		fd := newField(name, prop)
		if fd.Type == "" {
			return object{}, fmt.Errorf("unsupported type in nim:%v", prop.Type)
		}
		fields[fd.Name] = fd
	}

	return object{
		Name:        name,
		Fields:      fields,
		Description: commons.ParseDescription(description),
	}, nil
}

// generate nim object representation
func (o *object) generate(dir string) error {
	// generate enums
	for _, f := range o.Fields {
		if f.Enum != nil {
			if err := f.Enum.generate(dir); err != nil {
				return err
			}
		}
	}

	filename := filepath.Join(dir, o.Name+".nim")
	if err := commons.GenerateFile(o, "./templates/object_nim.tmpl", "object_nim", filename, true); err != nil {
		return err
	}
	registerObject(o.Name)
	return nil
}

// get array of all imported modules
func (o object) Imports() []string {
	ip := map[string]struct{}{}

	for _, f := range o.Fields {
		if objectRegistered(f.Type) {
			ip[f.Type] = struct{}{}
		}
		if f.Type == "Time" {
			ip["times"] = struct{}{}
		}
	}

	for _, p := range o.Parents {
		if objectRegistered(p) {
			ip[p] = struct{}{}
		}
	}
	return commons.MapToSortedStrings(ip)
}

// handle RAML advanced data type
func (o *object) handleAdvancedType() {
	if o.T.Type == nil {
		o.T.Type = "object"
	}
	strType := commons.InterfaceToString(o.T.Type)

	switch {
	case len(strings.Split(strType, ",")) > 1: //multiple inheritance
		// TODO
	case strings.ToLower(strType) == "object": // plain type
	case o.T.IsArray():
		o.makeArray(strType)
	}
}

func (o *object) makeArray(t string) {
	o.Parents = append(o.Parents, t[:len(t)-2])
	o.buildOneLine(toNimType(t))
}

func (o *object) buildOneLine(tipe string) {
	o.OneLineDef = fmt.Sprintf("%v* = %v", o.Name, tipe)
}

func registerObject(name string) {
	objectsRegister[name] = struct{}{}
}

// check if an object named `obj` has been generated
func objectRegistered(objName string) bool {
	_, ok := objectsRegister[objName]
	return ok
}
