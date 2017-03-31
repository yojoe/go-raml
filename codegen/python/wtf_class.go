package python

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/types"
	"github.com/Jumpscale/go-raml/raml"
)

// wtfClass defines a python wtform class
type wtfClass struct {
	T           raml.Type
	Name        string
	Description []string
	Fields      map[string]wtfField
	Enum        *enum
}

// create a python wtfClass representations
func newWtfClass(name, description string, properties map[string]interface{}) wtfClass {
	pc := wtfClass{
		Name:        name,
		Description: commons.ParseDescription(description),
		Fields:      map[string]wtfField{},
	}

	// generate fields
	for k, v := range properties {
		field, err := newWtfField(name, raml.ToProperty(k, v))
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
func generateAllWtfClasses(apiDef *raml.APIDefinition, dir string) error {
	for name, t := range types.AllTypes(apiDef, "") {
		switch tip := t.Type.(type) {
		case string:
			// TODO
		case types.TypeInBody:
			if tip.ReqResp == types.HTTPRequest {
				methodName := setServerMethodName(tip.Endpoint.Method.DisplayName, tip.Endpoint.Verb, tip.Endpoint.Resource)
				wtfClass := newWtfClass(setReqBodyName(methodName), "", tip.Properties)
				if err := wtfClass.generate(dir); err != nil {
					return err
				}

			}
		case raml.Type:
			pc := newWtfClassFromType(tip, name)
			if err := pc.generate(dir); err != nil {
				return err
			}
		}
	}
	return nil
}

func generateWtfClassesFromTypes(ts map[string]raml.Type, dir string) error {
	apiDef := raml.APIDefinition{
		Types: ts,
	}
	return generateAllWtfClasses(&apiDef, dir)
}
