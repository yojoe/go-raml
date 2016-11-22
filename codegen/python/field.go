package python

import (
	"fmt"
	"sort"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/raml"
)

// pythons class's field
type field struct {
	Name        string
	Type        string
	Required    bool
	Validators  string
	Enum        *enum
	ramlType    string // the original raml type
	isFormField bool
	isList      bool                // it is a list field
	validators  map[string][]string // array of validators, only used to build `Validators` field
}

func newField(className string, prop raml.Property) (field, error) {
	f := field{
		Name:     prop.Name,
		Required: prop.Required,
	}

	if prop.IsEnum() {
		f.Enum = newEnum(className, prop, false)
		f.Type = f.Enum.Name
	} else {
		f.setType(prop.Type)
		if f.Type == "" {
			return f, fmt.Errorf("unsupported type:%v", prop.Type)
		}
		f.buildValidators(prop)
	}

	return f, nil
}

// convert from raml Type to python wtforms type
func (pf *field) setType(t string) {
	pf.ramlType = t
	switch t {
	case "string":
		pf.Type = "TextField"
	case "file":
		pf.Type = "FileField"
	case "number":
		pf.Type = "FloatField"
	case "integer":
		pf.Type = "IntegerField"
	case "boolean":
		pf.Type = "BooleanField"
	case "date":
		pf.Type = "DateField"
	}

	if pf.Type != "" { // type already set, no need to go down
		return
	}

	// other types that need some processing
	switch {
	case strings.HasSuffix(t, "[][]"): // bidimensional array
		log.Info("validator has no support for bidimensional array, ignore it")
	case strings.HasSuffix(t, "[]"): // array
		pf.isList = true
		pf.setType(t[:len(t)-2])
	case strings.HasSuffix(t, "{}"): // map
		log.Info("validator has no support for map, ignore it")
	case strings.Index(t, "|") > 0:
		log.Info("validator has no support for union, ignore it")
	case strings.Index(t, ".") > 1:
		pf.Type = t[strings.Index(t, ".")+1:]
		pf.isFormField = true
	default:
		pf.isFormField = true
		pf.Type = t
	}

}

func (pf *field) addValidator(name, arg string, val interface{}) {
	pf.validators[name] = append(pf.validators[name], fmt.Sprintf("%v=%v", arg, val))
}

// build validators string
func (pf *field) buildValidators(p raml.Property) {
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
	if p.MultipleOf != nil {
		pf.addValidator("multiple_of", "mult", *p.MultipleOf)
	}

	// required
	if p.Required {
		pf.addValidator("DataRequired", "message", `""`)
	}

	if p.MinItems != nil {
		pf.Validators += fmt.Sprintf(",min_entries=%v", *p.MinItems)
	}
	if p.MaxItems != nil {
		pf.Validators += fmt.Sprintf(",max_entries=%v", *p.MaxItems)
	}
	if len(pf.Validators) > 0 {
		pf.Validators = pf.Validators[1:]
	}

	pf.buildValidatorsString()
}

func (pf *field) buildValidatorsString() {
	var v []string
	if pf.Validators != "" {
		return
	}
	for name, args := range pf.validators {
		v = append(v, fmt.Sprintf("%v(%v)", name, strings.Join(args, ", ")))
	}

	// we actually don't need to sort it to generate correct validators
	// we need to sort it to generate predictable order which needed during the test
	sort.Strings(v)
	pf.Validators = strings.Join(v, ", ")
}

// WTFType return wtforms type of a field
func (pf field) WTFType() string {
	switch {
	case pf.isList && pf.isFormField:
		return fmt.Sprintf("FieldList(FormField(%v))", pf.Type)
	case pf.isList:
		return fmt.Sprintf("FieldList(%v('%v', [required()]), %v)", pf.Type, pf.Name, pf.Validators)
	case pf.isFormField:
		return fmt.Sprintf("FormField(%v)", pf.Type)
	default:
		return fmt.Sprintf("%v(validators=[%v])", pf.Type, pf.Validators)
	}
}
