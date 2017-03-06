package angular

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

// pythons class's field
type field struct {
	Name         string
	Required     bool
	ramlType     string // the original raml type
	Model        string
	Form         string
	FormBuilder string
}

func newField(className string, prop raml.Property) (field, error) {
	f := field{
		Name:     prop.Name,
		Required: prop.Required,
	}

	f.ramlType = prop.Type

	if err := f.set_model(); err != nil {
		return f, err
	}
	if err := f.set_form(); err != nil {
		return f, err
	}
	if err := f.set_formbuilder(); err != nil {
		return f, err
	}
	return f, nil
}

func (f *field) set_formbuilder() error {
	f.FormBuilder = translate_raml_formbuilder_type(f.Name, f.ramlType)
	return nil
}
func (f *field) set_form() error {
	f.Form = translate_raml_form_type(f.Name, f.ramlType)
	return nil
}
func (f *field) set_model() error {
	f.Model = translate_raml_model_type(f.Name, f.ramlType)
	return nil
}

func translate_raml_formbuilder_type(name, raml string) string {
	switch raml {
	case "string":
		return "''"
	case "number":
		return "0.0"
	case "integer":
		return "0"
	case "date":
		return "''"
	case "file":
		return "??"
	case "boolean":
		return "false"
	}

	// other types that need some processing
	switch {
	case strings.HasSuffix(raml, "[]"): // array
		return "this.fb.array([])"
	default:
		return "'we do not support nested types yet'"
	}
}

func translate_raml_form_type(name, raml string) string {
	switch raml {
	case "string":
		fallthrough
	case "number":
		fallthrough
	case "integer":
		fallthrough
	case "date":
		return `<div class="form-group">
		<label class="center-block">` + name + `:
		<input class="form-control" formControlName="` + name + `">
		</label>
		</div>`
	case "file":
		return "??"
	case "boolean":
		return `<div class="form-group">
		<label class="center-block">` + name + `:
		<input type="checkbox" formControlName="` + name + `">
		</label>
		</div>`
	}

	// other types that need some processing
	switch {
	default:
		return "<p>we don't support nested types yet</p>"
	}
}
// WTFType return wtforms type of a field
func translate_raml_model_type(name, raml string) string {
	switch raml {
	case "string":
		return "string"
	case "file":
		return "??"
	case "number":
		return "number"
	case "integer":
		return "number"
	case "boolean":
		return "boolean"
	case "date":
		return "date"
	}

	// other types that need some processing
	switch {
	case strings.HasSuffix(raml, "[]"): // array
		return translate_raml_model_type(name, raml[:len(raml)-2]) + "[]"
	default:
		return raml
	}
}
