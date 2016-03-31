package codegen

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

type pythonField struct {
	Name       string
	Type       string
	Required   bool
	Validators string
	validators map[string][]string
}

func (pf *pythonField) addValidator(name, arg string, val interface{}) {
	pf.validators[name] = append(pf.validators[name], fmt.Sprintf("%v=%v", arg, val))
}

func (pf *pythonField) buildValidators(p raml.Property) {
	pf.validators = map[string][]string{}
	// string
	if p.MinLength != nil {
		pf.addValidator("Length", "min", *p.MinLength)
	}
	if p.MaxLength != nil {
		pf.addValidator("Length", "max", *p.MaxLength)
	}
	if p.Pattern != nil {
		pf.addValidator("Regexp", "regex", `"`+*p.Pattern+`"`)
	}

	// number
	if p.Minimum != nil {
		pf.addValidator("NumberRange", "min", *p.Minimum)
	}
	if p.Maximum != nil {
		pf.addValidator("NumberRange", "max", *p.Maximum)
	}

	// required
	if p.Required {
		pf.addValidator("DataRequired", "message", `""`)
	}
	pf.buildValidatorsString()
}

func (pf *pythonField) buildValidatorsString() {
	var v []string
	for name, args := range pf.validators {
		v = append(v, fmt.Sprintf("%v(%v)", name, strings.Join(args, ", ")))
	}

	// we actually don't need to sort it to generate correct validators
	// we need to sort it to generate predictable order which needed during the test
	sort.Strings(v)
	pf.Validators = strings.Join(v, ", ")
}

type pythonClass struct {
	T           raml.Type
	Name        string
	Description []string
	Fields      map[string]pythonField
}

func newPythonClass(name, description string, properties map[string]interface{}) pythonClass {
	pc := pythonClass{
		Name:        name,
		Description: commentBuilder(description),
		Fields:      map[string]pythonField{},
	}

	// generate fields
	for k, v := range properties {
		p := raml.ToProperty(k, v)
		field := pythonField{
			Name:     p.Name,
			Type:     toWtformsType(p.Type),
			Required: p.Required,
		}

		if field.Type == "" { // type is not supported, no need to generate the field
			continue
		}

		field.buildValidators(p)
		pc.Fields[p.Name] = field

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
