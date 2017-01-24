package python

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/chuckpreslar/inflect"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

// class defines a python class
type class struct {
	T           raml.Type
	Name        string
	Description []string
	Fields      map[string]field
	Enum        *enum
}

// create a python class representations
func newClass(name, description string, properties map[string]interface{}) class {
	pc := class{
		Name:        name,
		Description: commons.ParseDescription(description),
		Fields:      map[string]field{},
	}

	// generate fields
	for k, v := range properties {
		field, err := newField(name, raml.ToProperty(k, v))
		if err != nil {
			continue
		}
		pc.Fields[field.Name] = field

	}
	return pc
}

func newClassFromType(T raml.Type, name string) class {
	pc := newClass(name, T.Description, T.Properties)
	pc.T = T
	pc.handleAdvancedType()
	return pc
}

// generate a python class file
func (pc *class) generate(dir string) error {
	// generate enums
	for _, f := range pc.Fields {
		if f.Enum != nil {
			if err := f.Enum.generate(dir); err != nil {
				return err
			}
		}
	}

	if pc.Enum != nil {
		return pc.Enum.generate(dir)
	}

	fileName := filepath.Join(dir, pc.Name+".py")
	return commons.GenerateFile(pc, "./templates/class_python.tmpl", "class_python", fileName, false)
}

func (pc *class) handleAdvancedType() {
	if pc.T.Type == nil {
		pc.T.Type = "object"
	}
	if pc.T.IsEnum() {
		pc.Enum = newEnumFromClass(pc)
	}
}

// generate all classes from all  methods request/response bodies
func generateClassesFromBodies(rs []pythonResource, dir string) error {
	for _, r := range rs {
		for _, mi := range r.Methods {
			m := mi.(serverMethod)
			if err := generateClassesFromMethod(m, dir); err != nil {
				return err
			}
		}
	}
	return nil
}

// generate classes from a method
//
// TODO:
// we currently camel case instead of snake case because of mistake in previous code
// and we might need to maintain backward compatibility. Fix this!
func generateClassesFromMethod(m serverMethod, dir string) error {
	// request body
	if commons.HasJSONBody(&m.Bodies) {
		name := inflect.UpperCamelCase(m.MethodName + "ReqBody")
		class := newClass(name, "", m.Bodies.ApplicationJSON.Properties)
		if err := class.generate(dir); err != nil {
			return err
		}
	}
	return nil
}

// return list of import statements
func (pc class) Imports() []string {
	var imports []string

	for _, v := range pc.Fields {
		if v.isFormField {
			if strings.Index(v.ramlType, ".") > 1 { // it is a library
				importPath, name := libImportPath(v.ramlType, "")
				imports = append(imports, "from "+importPath+" import "+name)
			} else {
				imports = append(imports, "from "+v.Type+" import "+v.Type)
			}
		}
	}
	sort.Strings(imports)
	return imports
}

// generate all python classes from an RAML document
func generateClasses(types map[string]raml.Type, dir string) error {
	for k, t := range types {
		pc := newClassFromType(t, k)
		if err := pc.generate(dir); err != nil {
			return err
		}
	}
	return nil
}
