package angular

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

// class defines a angular class
type class struct {
	T           raml.Type
	Name        string
	LName       string
	Description []string
	Fields      map[string]field
}

// create a angular class representations
func newClass(name, description string, properties map[string]interface{}) class {
	pc := class{
		Name:        name,
		Description: commons.ParseDescription(description),
		Fields:      map[string]field{},
	}
	pc.LName = strings.ToLower(pc.Name)

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

// generate a angular class file
func (pc *class) generate(dir string) error {
	// generate enums
	fileName := filepath.Join(dir, pc.LName+".ts")
	if err := commons.GenerateFile(pc, "./templates/angular/templates/class_model_angular.tmpl", "class_model_angular", fileName, false); err != nil {
		return err
	}
	fileName = filepath.Join(dir, pc.LName+"-form.component.ts")
	if err := commons.GenerateFile(pc, "./templates/angular/templates/class_component_angular.tmpl", "class_component_angular", fileName, false); err != nil {
		return err
	}
	fileName = filepath.Join(dir, pc.LName+"-form.component.html")
	if err := commons.GenerateFile(pc, "./templates/angular/templates/class_template_angular.tmpl", "class_template_angular", fileName, false); err != nil {
		return err
	}
	return nil
}

func (pc *class) handleAdvancedType() {
	if pc.T.Type == nil {
		pc.T.Type = "object"
	}
}

// generate all angular classes from an RAML document
func GenerateClasses(classes map[string]class, dir string) error {
	for k, pc := range classes {
		if err := GenerateClass(k, pc, dir); err != nil {
			return err
		}
	}
	return nil
}

func GenerateClass(name string, pc class, dir string) error {
	classdir := filepath.Join(dir, name)
	os.Mkdir(classdir, 0755)
	return pc.generate(classdir)
}
