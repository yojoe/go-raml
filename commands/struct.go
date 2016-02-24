package commands

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

var (
	structTemplateLocation = "./templates/struct.tmpl"
)

// FieldDef defines a field of a struct
type fieldDef struct {
	Name          string
	Type          string
	IsComposition bool
}

// StructDef defines a struct
type structDef struct {
	Name         string              // struct's name
	Description  []string            // structs description
	PackageName  string              // package name
	Fields       map[string]fieldDef // all struct's fields
	OneLineDef   string              // not empty if this struct can defined in one line
	IsOneLineDef bool
	t            raml.Type // raml.Type of this struct
}

// create new struct def
func newStructDef(name, packageName, description string, properties map[string]raml.Property) structDef {
	// generate struct's fields from type properties
	fields := make(map[string]fieldDef)
	for k, v := range properties {
		goType := convertToGoType(v.Type)
		if v.Enum != nil {
			goType = "[]" + goType
		}

		fd := fieldDef{
			Name: strings.Title(k),
			Type: goType, // convert to internal field type
		}
		fields[k] = fd
	}

	return structDef{
		Name:        strings.Title(name),
		PackageName: packageName,
		Fields:      fields,
		Description: funcCommentBuilder(description),
	}

}

// create struct definition from RAML Type node
func newStructDefFromType(t raml.Type, sName, packageName, description string) structDef {
	sd := newStructDef(sName, packageName, description, t.Properties)
	sd.t = t

	// handle inheritance on raml1.0
	sd.handleAdvancedType(t)

	return sd
}

// create struct definition from RAML Body node
func newStructDefFromBody(body *raml.Bodies, structNamePrefix, packageName string, isGenerateRequest bool) structDef {
	// set struct name based on request or response
	structName := structNamePrefix + respBodySuffix
	if isGenerateRequest {
		structName = structNamePrefix + reqBodySuffix
	}

	return newStructDef(structName, packageName, "", body.ApplicationJson.Properties)
}

// generate Go struct
func (sd structDef) generate(dir string) error {
	fileName := dir + "/" + strings.ToLower(sd.Name) + ".go"
	if err := generateFile(sd, structTemplateLocation, "struct_template", fileName, false); err != nil {
		return err
	}
	return runGoFmt(fileName)
}

// generate all structs from an RAML api definition
func generateStructs(apiDefinition *raml.APIDefinition, dir string, packageName string) error {
	if err := checkCreateDir(dir); err != nil {
		return err
	}
	for k, v := range apiDefinition.Types {
		sd := newStructDefFromType(v, k, packageName, v.Description)
		if err := sd.generate(dir); err != nil {
			return err
		}
	}
	return nil
}

// handle advance type type into structField
// example:
//   Mammal:
//     type: Animal
//     properties:
//       name:
//         type: string
// the additional fieldDef would be Animal composition
func (sd *structDef) handleAdvancedType(t raml.Type) {
	if t.Type == nil {
		return
	}
	strType := interfaceToString(t.Type)

	switch {
	case len(strings.Split(strType, ",")) > 1: //multiple inheritance
		sd.addMultipleInheritance(strType)
	case sd.t.IsArray(): // arary type
		sd.buildArray()
	case sd.t.IsMap(): // map
		sd.buildMap()
	case strings.ToLower(strType) == "object": // plain type
		return
	case sd.t.IsEnum():
		sd.buildEnum()
	default: // single inheritance
		sd.addSingleInheritance(strType)
	}
}

// add single inheritance
// inheritance is implemented as composition
// spec : http://docs.raml.org/specs/1.0/#raml-10-spec-inheritance-and-specialization
func (sd *structDef) addSingleInheritance(strType string) {
	fd := fieldDef{
		Name:          strings.Title(strType),
		Type:          strings.Title(strType),
		IsComposition: true,
	}
	sd.Fields[strType] = fd

}

// construct multiple inheritance to Go type
// example :
// Anggora:
//	 type: [ Animal , Cat ]
//	 properties:
//		color:
//			type: string
// The additional fielddef would be a composition of Animal & Cat
// http://docs.raml.org/specs/1.0/#raml-10-spec-multiple-inheritance
func (sd *structDef) addMultipleInheritance(strType string) {
	for _, s := range strings.Split(strType, ",") {
		fieldType := strings.TrimSpace(s)
		fd := fieldDef{
			Name:          strings.Title(fieldType),
			Type:          fieldType,
			IsComposition: true,
		}

		sd.Fields[fd.Name] = fd
	}
}

// buildEnum based on http://docs.raml.org/specs/1.0/#raml-10-spec-enums
// example result  `type TypeName []data_type`
func (sd *structDef) buildEnum() {
	if _, ok := sd.t.Type.(string); !ok {
		return
	}

	sd.IsOneLineDef = true
	sd.OneLineDef = "type " + sd.Name + "[]" + convertToGoType(sd.t.Type.(string))
}

// build map type based on http://docs.raml.org/specs/1.0/#raml-10-spec-map-types
// result is `type TypeName map[string]something`
func (sd *structDef) buildMap() {
	typeFromSquareBracketProp := func() string {
		var p raml.Property
		for _, v := range sd.t.Properties {
			if v.Type == "" {
				v.Type = "string"
			}
			p = v
			break
		}
		return p.Type
	}
	sd.IsOneLineDef = true
	switch {
	case sd.t.AdditionalProperties != "":
		sd.OneLineDef = "type " + sd.Name + " map[string]" + convertToGoType(sd.t.AdditionalProperties)
	case len(sd.t.Properties) == 1:
		sd.OneLineDef = "type " + sd.Name + " map[string]" + typeFromSquareBracketProp()
	default:
		sd.IsOneLineDef = false
	}
}

// build array type
// spec http://docs.raml.org/specs/1.0/#raml-10-spec-array-types
// example result  `type TypeName []something`
func (sd *structDef) buildArray() {
	sd.IsOneLineDef = true
	sd.OneLineDef = "type " + sd.Name + " " + convertToGoType(sd.t.Type.(string))
}
