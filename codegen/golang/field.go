package golang

import (
	"fmt"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

// FieldDef defines a field of a struct
type fieldDef struct {
	Name          string // field name
	Type          string // field type
	IsComposition bool   // composition type
	IsOmitted     bool   // omitted empty
	UniqueItems   bool
	Enum          *enum // not nil if this field contains enum

	Validators string
}

func newFieldDef(structName string, prop raml.Property, pkg string) fieldDef {
	fd := fieldDef{
		Name:      strings.Title(prop.Name),
		Type:      convertToGoType(prop.Type),
		IsOmitted: !prop.Required,
	}
	fd.buildValidators(prop)
	if prop.IsEnum() {
		fd.Enum = newEnum(structName, prop, pkg, false)
		fd.Type = fd.Enum.Name
	}

	return fd
}

func (fd *fieldDef) buildValidators(p raml.Property) {
	validators := ""
	// string
	if p.MinLength != nil {
		validators += fmt.Sprintf(",min=%v", *p.MinLength)
	}
	if p.MaxLength != nil {
		validators += fmt.Sprintf(",max=%v", *p.MaxLength)
	}
	if p.Pattern != nil {
		validators += fmt.Sprintf(",regexp=%v", *p.Pattern)
	}

	// Number
	if p.Minimum != nil {
		validators += fmt.Sprintf(",min=%v", *p.Minimum)
	}

	if p.Maximum != nil {
		validators += fmt.Sprintf(",max=%v", *p.Maximum)
	}

	if p.MultipleOf != nil {
		validators += fmt.Sprintf(",multipleOf=%v", *p.MultipleOf)
	}

	//if p.Format != nil {
	//}

	// Array & Map
	if p.MinItems != nil {
		validators += fmt.Sprintf(",min=%v", *p.MinItems)
	}
	if p.MaxItems != nil {
		validators += fmt.Sprintf(",max=%v", *p.MaxItems)
	}
	if p.UniqueItems {
		fd.UniqueItems = true
	}

	// Required
	if !fd.IsOmitted {
		validators += ",nonzero"
	}

	if validators != "" {
		fd.Validators = validators[1:]
	}
}
