package python

import (
	"path/filepath"
	"sort"
	"strings"

	"fmt"
	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/types"
	"github.com/Jumpscale/go-raml/raml"
)

// class defines a python class
type class struct {
	T                 raml.Type
	Name              string
	Description       []string
	Fields            map[string]field
	Enum              *enum
	CreateParamString string
	MyPy              bool
}

type objectProperty struct {
	name            string
	required        bool
	datatype        string
	childProperties []objectProperty
}

// create a python class representations
func newClass(T raml.Type, name string, description string, properties map[string]interface{}) class {
	pc := class{
		Name:        name,
		Description: commons.ParseDescription(description),
		Fields:      map[string]field{},
		T:           T,
	}
	types := globAPIDef.Types

	typeHierarchy := getTypeHierarchy(name, T, types)
	ramlTypes := make([]raml.Type, 0)
	for _, v := range typeHierarchy {
		for _, iv := range v {
			ramlTypes = append(ramlTypes, iv)
		}
	}
	mergedProps := getTypeProperties(ramlTypes)
	for k, v := range properties {
		prop := raml.ToProperty(k, v)
		mergedProps[k] = prop
	}

	for propName, propInterface := range mergedProps {
		op := objectProperties(propName, propInterface)
		field, err := newField(name, T, propName, propInterface, types, op, typeHierarchy)
		if err != nil {
			continue
		}

		pc.Fields[field.Name] = field
	}

	return pc
}

func objectProperties(name string, p interface{}) []objectProperty {
	props := make([]objectProperty, 0)

	switch prop := p.(type) {
	case map[interface{}]interface{}:
		ramlProp := raml.ToProperty(name, p)
		if ramlProp.Type == "object" {
			for k, v := range prop {
				switch k {
				case "properties":
					for propName, childProp := range v.(map[interface{}]interface{}) {
						rProp := raml.ToProperty(propName.(string), childProp)
						objprop := objectProperty{
							name:     rProp.Name,
							required: rProp.Required,
							datatype: rProp.TypeString(),
						}
						if rProp.Type == "object" {
							objprop.childProperties = append(objprop.childProperties, objectProperties(propName.(string), childProp)...)
						}
						props = append(props, objprop)
					}
				}
			}
		}
	}

	return props
}

// ChildProperties returns child properties of a property
func ChildProperties(Properties map[string]interface{}) []raml.Property {
	props := make([]raml.Property, 0)

	for propName, propInterface := range Properties {
		props = append(props, raml.ToProperty(propName, propInterface))
	}

	return props
}

func getTypeHierarchy(name string, T raml.Type, types map[string]raml.Type) []map[string]raml.Type {
	typelist := []map[string]raml.Type{map[string]raml.Type{name: T}}

	for _, parent := range T.Parents() {
		parentType, inherited := types[parent]
		if inherited {
			for _, pt := range getTypeHierarchy(parent, parentType, types) {
				typelist = append(typelist, pt)
			}
		}
	}

	return typelist
}

func getTypeProperties(typelist []raml.Type) map[string]raml.Property {
	// get a list of the types in the inheritance chain for T
	// walk it from the top down and add the properties
	properties := map[string]raml.Property{}
	for i := len(typelist) - 1; i >= 0; i-- {
		for k, v := range typelist[i].Properties {
			prop := raml.ToProperty(k, v)
			// we convert it to property here
			// because we need the proper name, sometimes the name has "?" suffix
			//   which mean optional properties
			properties[prop.Name] = prop
		}
	}

	return properties
}

func newMyPyClassFromType(T raml.Type, name string) class {
	pc := newClass(T, name, T.Description, T.Properties)
	pc.T = T
	pc.handleAdvancedType()
	pc.createParamString()
	return pc
}

func newClassFromType(T raml.Type, name string) class {
	pc := newClass(T, name, T.Description, T.Properties)
	pc.T = T
	pc.handleAdvancedType()
	return pc
}

// generate a python class file
func (pc *class) generate(dir string, template string, name string) ([]string, error) {
	// generate enums
	typeNames := make([]string, 0)
	for _, f := range pc.Fields {
		if f.Enum != nil {
			typeNames = append(typeNames, f.Enum.Name)
			if err := f.Enum.generate(dir); err != nil {
				return typeNames, err
			}
		}
	}

	if pc.Enum != nil {
		typeNames = append(typeNames, pc.Enum.Name)
		return typeNames, pc.Enum.generate(dir)
	}

	fileName := filepath.Join(dir, pc.Name+".py")
	typeNames = append(typeNames, pc.Name)
	return typeNames, commons.GenerateFile(pc, template, name, fileName, false)
}

func (pc *class) createParamString() {
	// build the CreateParamString, used as part of the create() staticmethod
	// which is a convenience initializer for the class
	requiredFields := make([]string, 0)
	optionalFields := make([]string, 0)

	for fieldName, field := range pc.Fields {
		if field.Required {
			requiredFields = append(requiredFields, fmt.Sprintf("%s: %s", fieldName, field.MyPyType))
		} else {
			optionalFields = append(optionalFields, fmt.Sprintf("%s: %s=None", fieldName, field.MyPyType))
		}
	}
	// sort them so we have some stability in param order (important for requiredFields)
	sort.Strings(requiredFields)
	sort.Strings(optionalFields)
	pc.CreateParamString = strings.Join(append(requiredFields, optionalFields...), ", ")
}

func (pc *class) handleAdvancedType() {
	if pc.T.Type == nil {
		pc.T.Type = "object"
	}
	if pc.T.IsEnum() {
		pc.Enum = newEnumFromClass(pc)
	}
}

// return list of import statements
func (pc class) Imports() []string {
	// var imports []string
	imports := make(map[string]struct{})

	for _, field := range pc.Fields {
		for _, imp := range field.imports {
			// Ignore mypy imports if this is not a mypy class
			if !pc.MyPy && imp.Module == "typing" || pc.MyPy && imp.Module == "six" {
				continue
			}

			// do not import ourself
			if imp.Name == pc.Name {
				continue
			}

			importString := "from " + imp.Module + " import " + imp.Name
			imports[importString] = struct{}{}
		}
	}

	return commons.MapToSortedStrings(imports)
}

// generate all python classes from a RAML document
func generateAllClasses(apiDef *raml.APIDefinition, dir string) ([]string, error) {
	// array of tip that need to be generated in the end of this
	// process. because it needs other object to be registered first
	delayedMI := []string{} // delayed multiple inheritance
	template := "./templates/python/class_python.tmpl"
	templateName := "class_python"

	names := []string{}
	for name, t := range types.AllTypes(apiDef, "") {
		var errGen error
		var results []string
		switch tip := t.Type.(type) {
		case string:
			rt := raml.Type{
				Type: tip,
			}
			if rt.IsMultipleInheritance() {
				delayedMI = append(delayedMI, tip)
			}
		case types.TypeInBody:
			methodName := setServerMethodName(tip.Endpoint.Method.DisplayName, tip.Endpoint.Verb, tip.Endpoint.Resource)
			pc := newClass(raml.Type{Type: "object"}, setReqBodyName(methodName), "", tip.Properties)
			propNames := []string{}
			for k := range tip.Properties {
				propNames = append(propNames, k)
			}
			results, errGen = pc.generate(dir, template, templateName)
		case raml.Type:
			if name == "UUID" {
				continue
			}
			pc := newClassFromType(tip, name)
			results, errGen = pc.generate(dir, template, templateName)
		}

		if errGen != nil {
			return names, errGen
		}
		names = append(names, results...)
	}

	for _, tip := range delayedMI {
		rt := raml.Type{
			Type: tip,
		}
		if parents, isMult := rt.MultipleInheritance(); isMult {
			pc := newClassFromType(rt, strings.Join(parents, ""))
			results, err := pc.generate(dir, template, templateName)
			if err != nil {
				return names, err
			}
			names = append(names, results...)
		}

	}
	return names, nil

}

// generate all mypy python classes from a RAML document type section.
func generateMyPyClasses(apiDef *raml.APIDefinition, dir string) error {
	// from types
	for name, tip := range apiDef.Types {
		if name == "UUID" {
			continue
		}
		pc := newMyPyClassFromType(tip, name)
		pc.MyPy = true
		if _, errGen := pc.generate(dir, "./templates/python/class_python_mypy.tmpl", "class_python_mypy"); errGen != nil {
			return errGen
		}
	}
	return nil
}
