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
	Name string
	Type string
}

// StructDef defines a struct
type structDef struct {
	Name        string
	PackageName string
	Fields      map[string]fieldDef
}

//create struct from type node
func newStructDefFromType(t raml.Type, sName, packageName string) structDef {
	structField := make(map[string]fieldDef)
	for k, v := range t.Properties {
		//upper first char
		fieldName := strings.Title(k)

		//convert to internal field type
		fieldType := convertToGoType(v.Type)

		fd := fieldDef{
			Name: fieldName,
			Type: fieldType,
		}

		structField[fieldName] = fd
	}

	//handle inheritance on raml1.0
	addStructInheritance(structField, t)

	structDef := structDef{
		Name:        strings.Title(sName),
		PackageName: packageName,
		Fields:      structField,
	}
	return structDef
}

//create struct from body node
func newStructDefFromBody(body *raml.Bodies, structNamePrefix, packageName string, isGenerateRequest bool) structDef {
	structField := make(map[string]fieldDef)

	for k, v := range body.ApplicationJson.Properties {
		//upper first char
		fieldName := strings.Title(k)

		//convert to internal field type
		fieldType := convertToGoType(v.Type)

		fd := fieldDef{
			Name: fieldName,
			Type: fieldType,
		}

		structField[fieldName] = fd
	}

	structName := strings.Title(structNamePrefix + respBodySuffix)
	if isGenerateRequest {
		structName = strings.Title(structNamePrefix + reqBodySuffix)
	}

	return structDef{
		Name:        structName,
		Fields:      structField,
		PackageName: packageName,
	}
}

func (sd structDef) generate(dir string) error {
	fileName := dir + "/" + strings.ToLower(sd.Name) + ".go"
	if err := generateFile(sd, structTemplateLocation, "struct_template", fileName, false); err != nil {
		return err
	}
	return runGoFmt(fileName)
}

//GenerateStruct generate struct
func GenerateStruct(apiDefinition *raml.APIDefinition, dir string, packageName string) error {
	if err := checkCreateDir(dir); err != nil {
		return err
	}
	for k, v := range apiDefinition.Types {
		sd := newStructDefFromType(v, k, packageName)
		if err := sd.generate(dir); err != nil {
			return err
		}
	}
	return nil
}

//addStructInheritance add inheritance type into structField
//example:
//  Mammal:
//    type: Animal
//    properties:
//      name:
//        type: string
// the additional fieldDef would be Animal Animal
func addStructInheritance(structField map[string]fieldDef, t raml.Type) {
	if t.Type != nil {
		strType := interfaceToString(t.Type)

		if strings.ToLower(strType) == "object" {
			return
		}
		//handle multiple([A , B]) or common inheritance
		if len(strings.Split(strType, ",")) > 1 {
			constructMultipleInheritance(structField, strType)
		} else if strings.HasSuffix(strType, "[]") {
			constructArrayTypeInheritance(structField, strType)
		} else {
			fd := fieldDef{
				Name: strings.Title(strType),
				Type: strings.Title(strType),
			}

			structField[strType] = fd
		}
	}
}

//Construct multiple inheritance construct multiple inheritance to gotype
//example :
//Anggora:
//	type: [ Animal , Cat ]
//	properties:
//		color:
//			type: string
// The additional fielddef would be :
// Animal Animal and Cat Cat
func constructMultipleInheritance(structField map[string]fieldDef, strType string) {
	splitByComa := strings.Split(strType, ",")
	for i := 0; i < len(splitByComa); i++ {
		fieldType := strings.TrimSpace(splitByComa[i])
		fieldName := strings.Title(fieldType)

		fd := fieldDef{
			Name: strings.Title(fieldName),
			Type: fieldType,
		}

		structField[fieldName] = fd
	}
}

//Construct array specialization from inheritance spec
//example :
//MyType:
//	type: Person[]
//Additional structField would be :
// person []Person
func constructArrayTypeInheritance(structField map[string]fieldDef, strType string) {
	replaceBracketAndTrimmed := strings.TrimSpace(strings.Replace(strType, "[]", "", -1))
	fieldName := replaceBracketAndTrimmed
	fieldType := "[]" + strings.Title(replaceBracketAndTrimmed)

	fd := fieldDef{
		Name: strings.Title(fieldName),
		Type: fieldType,
	}

	structField[fieldName] = fd
}
