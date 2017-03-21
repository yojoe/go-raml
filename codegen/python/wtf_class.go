package python

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/chuckpreslar/inflect"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

// wtfClass defines a python wtform class
type wtfClass struct {
	T           raml.Type
	Name        string
	Description []string
	Fields      map[string]field
	Enum        *enum
}

// create a python wtfClass representations
func newWtfClass(name, description string, properties map[string]interface{}) wtfClass {
	pc := wtfClass{
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

func newWtfClassFromType(T raml.Type, name string) wtfClass {
	pc := newWtfClass(name, T.Description, T.Properties)
	pc.T = T
	pc.handleAdvancedType()
	return pc
}

// generate a python wtfClass file
func (pc *wtfClass) generate(dir string) error {
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
	return commons.GenerateFile(pc, "./templates/wtf_class_python.tmpl", "wtf_class_python", fileName, false)
}

func (pc *wtfClass) EmptyField() bool {
	return len(pc.Fields) == 0
}

func (pc *wtfClass) handleAdvancedType() {
	if pc.T.Type == nil {
		pc.T.Type = "object"
	}
	if pc.T.IsEnum() {
		pc.Enum = newEnumFromWtfClass(pc)
	}
}

// generate all wtfClasses from all  methods request/response bodies
func (fs FlaskServer) generateWtfClassesFromBodies(dir string) error {
	for _, r := range fs.ResourcesDef {
		for _, mi := range r.Methods {
			m := mi.(serverMethod)
			if err := generateWtfClassesFromMethod(m, dir); err != nil {
				return err
			}
		}
	}
	return nil
}

// generate wtfClasses from a method
//
// TODO:
// we currently camel case instead of snake case because of mistake in previous code
// and we might need to maintain backward compatibility. Fix this!
func generateWtfClassesFromMethod(m serverMethod, dir string) error {
	// request body
	if commons.HasJSONBody(&m.Bodies) {
		name := inflect.UpperCamelCase(m.MethodName + "ReqBody")
		wtfClass := newWtfClass(name, "", m.Bodies.ApplicationJSON.Properties)
		if err := wtfClass.generate(dir); err != nil {
			return err
		}
	}
	return nil
}

// return list of import statements
func (pc wtfClass) Imports() []string {
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

// generate all python wtfClasses from an RAML document
func generateWtfClasses(types map[string]raml.Type, dir string) error {
	for k, t := range types {
		pc := newWtfClassFromType(t, k)
		if err := pc.generate(dir); err != nil {
			return err
		}
	}
	return nil
}
