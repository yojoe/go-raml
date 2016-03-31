package codegen

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

type pythonField struct {
	Name       string
	Type       string
	Validators map[string][]string
}

func (pf *pythonField) addValidator(name, arg string, val interface{}) {
	pf.Validators[name] = append(pf.Validators[name], fmt.Sprintf("%v=%v", arg, val))
}

func (pf *pythonField) buildValidators(p raml.Property) {
	pf.Validators = map[string][]string{}
	// string
	if p.MinLength != nil {
		pf.addValidator("Length", "min", *p.MinLength)
	}
	if p.MaxLength != nil {
		pf.addValidator("Length", "max", *p.MaxLength)
	}
}

func (pf pythonField) ValidatorsString() string {
	var v string
	for name, args := range pf.Validators {
		v += fmt.Sprintf("%v(%v)", name, strings.Join(args, ", "))
	}
	return v
}

type pythonClass struct {
	T           raml.Type
	Name        string
	Description string
	Fields      []pythonField
}

func newPythonClass(name, description string, properties map[string]interface{}) pythonClass {
	pc := pythonClass{
		Name:        name,
		Description: description,
	}

	// generate fields
	for k, v := range properties {
		p := raml.ToProperty(k, v)
		field := pythonField{
			Name: p.Name,
			Type: toWtformsType(p.Type),
		}

		if field.Type == "" { // type is not supported, no need to generate the field
			continue
		}

		field.buildValidators(p)
		pc.Fields = append(pc.Fields, field)
	}
	return pc
}

func newPythonClassFromType(T raml.Type, name string) pythonClass {
	pc := newPythonClass(name, T.Description, T.Properties)
	pc.T = T
	return pc
}

func (pc *pythonClass) generate(dir string) error {
	fileName := filepath.Join(dir, pc.Name+".py")
	return generateFile(pc, "./templates/class_python.tmpl", "class_python", fileName, false)
}

// convert from raml Type to python wtforms type
func toWtformsType(t string) string {
	switch t {
	case "string":
		return "TextField"
	case "file":
		return "FileField"
	case "number":
		return "FloatField"
	case "integer":
		return "IntegerField"
	case "boolean":
		return "BooleanField"
	case "date":
		return "DateField"
	default:
		return ""
	}
}

// generate all python classes from an RAML document
func generatePythonClasses(apiDef *raml.APIDefinition, dir string) error {
	for k, t := range apiDef.Types {
		pc := newPythonClassFromType(t, k)
		if err := pc.generate(dir); err != nil {
			return err
		}
	}
	return nil
}
