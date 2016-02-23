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
	Name        string
	Description []string
	PackageName string
	Fields      map[string]fieldDef
}

// create new struct def
func newStructDef(name, packageName, description string, properties map[string]raml.Property) structDef {
	// generate struct's fields from type properties
	fields := make(map[string]fieldDef)
	for k, v := range properties {
		fd := fieldDef{
			Name: strings.Title(k),
			Type: convertToGoType(v.Type), // convert to internal field type
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
	structDef := newStructDef(sName, packageName, description, t.Properties)

	// handle inheritance on raml1.0
	structDef.addInheritance(t)

	return structDef
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

// add inheritance type into structField
// example:
//   Mammal:
//     type: Animal
//     properties:
//       name:
//         type: string
// the additional fieldDef would be Animal composition
func (sd *structDef) addInheritance(t raml.Type) {
	if t.Type == nil {
		return
	}
	strType := interfaceToString(t.Type)

	if strings.ToLower(strType) == "object" {
		return
	}
	if len(strings.Split(strType, ",")) > 1 { //handle multiple([A , B]) or common inheritance
		sd.addMultipleInheritance(strType)
	} else if strings.HasSuffix(strType, "[]") { // array
		sd.addArraySpecialization(strType)
	} else {
		fd := fieldDef{
			Name:          strings.Title(strType),
			Type:          strings.Title(strType),
			IsComposition: true,
		}
		sd.Fields[strType] = fd
	}
}

// construct multiple inheritance to Go type
// example :
// Anggora:
//	 type: [ Animal , Cat ]
//	 properties:
//		color:
//			type: string
// The additional fielddef would be a composition of Animal & Cat
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

// Construct array specialization from inheritance spec
// example :
//  MyType:
//	  type: Person[]
// Additional structField would be :
//  Person []Person
func (sd *structDef) addArraySpecialization(strType string) {
	fieldName := strings.TrimSpace(strings.Replace(strType, "[]", "", -1))
	fieldType := "[]" + strings.Title(fieldName)

	fd := fieldDef{
		Name: strings.Title(fieldName),
		Type: fieldType,
	}

	sd.Fields[fieldName] = fd
}
