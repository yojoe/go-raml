package nim

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/types"
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
	Enum        *enum
}

func generateAllObjects(apiDef *raml.APIDefinition, dir string) error {
	objs := []object{}
	for name, t := range types.AllTypes(apiDef, "") {
		switch tip := t.Type.(type) {
		case string:
			//TODO
		case types.TypeInBody:
			suffix := commons.RespBodySuffix
			if tip.ReqResp == types.HTTPRequest {
				suffix = commons.ReqBodySuffix
			}
			verb := strings.Title(strings.ToLower(tip.Endpoint.Verb))
			bodyName := setBodyName(tip.Body(), tip.Endpoint.Addr+verb, suffix)
			fmt.Printf("name=%v, bodyName=%v\n", name, bodyName)
			obj, err := newObject(bodyName, "", tip.Properties)
			if err != nil {
				return err
			}
			registerObject(obj.Name)
			objs = append(objs, obj)

		case raml.Type:
			obj, err := newObjectFromType(tip, t.Name)
			if err != nil {
				return err
			}

			registerObject(obj.Name)

			for _, f := range obj.Fields {
				if f.Enum != nil {
					registerObject(f.Enum.Name)
				}
			}
			objs = append(objs, obj)
		}
	}
	for _, obj := range objs {
		if err := obj.generate(dir); err != nil {
			return err
		}
	}
	return nil
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

	if o.Enum != nil {
		return o.Enum.generate(dir)
	}
	filename := filepath.Join(dir, o.Name+".nim")
	if err := commons.GenerateFile(o, "./templates/object_nim.tmpl", "object_nim", filename, true); err != nil {
		return err
	}
	return nil
}

// get array of all imported modules:
// - `times` module if it use Time
// - other generated modules that used as field
// - parent object
func (o object) Imports() []string {
	ip := map[string]struct{}{}

	for _, f := range o.Fields {
		if name, ok := objectRegistered(f.Type); ok {
			ip[name] = struct{}{}
		}
		if f.Type == "Time" {
			ip["times"] = struct{}{}
		}
	}

	for _, p := range o.Parents {
		if name, ok := objectRegistered(p); ok {
			ip[name] = struct{}{}
		}
	}
	// del reference to our self
	if _, ok := ip[o.Name]; ok {
		delete(ip, o.Name)
	}
	return commons.MapToSortedStrings(ip)
}

// handle RAML advanced data type
func (o *object) handleAdvancedType() {
	if o.T.Type == nil {
		o.T.Type = "object"
	}
	strType := o.T.TypeString()

	switch {
	case len(strings.Split(strType, ",")) > 1: //multiple inheritance
		// TODO
	case o.T.IsEnum():
		o.makeEnum()
	case strings.ToLower(strType) == "object": // plain type
	case o.T.IsArray():
		o.makeArray(strType)
	}
}

func (o *object) makeEnum() {
	o.Enum = newEnumFromObject(o)
}
func (o *object) makeArray(t string) {
	o.Parents = append(o.Parents, toNimType(t))
	o.buildOneLine(toNimType(t))
}

func (o *object) buildOneLine(tipe string) {
	o.OneLineDef = fmt.Sprintf("%v* = %v", o.Name, tipe)
}

func registerObject(name string) {
	name = strings.TrimSpace(name)
	objectsRegister[name] = struct{}{}
}

// regex to find string inside a square bracket
var reSquareBracket = regexp.MustCompile(`\[(.*?)\]`)

// check if an object named `objName` has been generated
func objectRegistered(objName string) (string, bool) {
	objName = extractTypeName(objName)
	_, ok := objectsRegister[objName]
	return objName, ok
}

// extract type name from potentially complex type declaration
func extractTypeName(tip string) string {
	// trim all spaces
	tip = strings.TrimSpace(tip)

	// detect array
	if strings.Index(tip, "]") < 0 {
		return tip
	}

	// find type inside square bracket
	inBracket := reSquareBracket.FindString(tip)
	if inBracket == "" {
		return tip
	}
	last := strings.LastIndex(inBracket, "[")
	inBracket = inBracket[last+1:]
	return strings.TrimSuffix(inBracket, "]")
}
