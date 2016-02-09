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
	Name   string
	Fields map[string]fieldDef
}

//create struct from type node
func newStructDefFromType(t raml.Type, sName string) structDef {
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

	structDef := structDef{
		Name:   strings.Title(sName),
		Fields: structField,
	}
	return structDef
}

//create struct from body node
func newStructDefFromBody(body *raml.Bodies, structNamePrefix string, isPartial bool) (structDef, structDef) {
	var reqStructDef, respStructDef structDef
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

	if !isPartial {
		//build struct def for req
		reqStructDef = structDef{
			Name:   strings.Title(structNamePrefix + reqBodySuffix),
			Fields: structField,
		}
	}

	//build structdef for resp
	respStructDef = structDef{
		Name:   strings.Title(structNamePrefix + respBodySuffix),
		Fields: structField,
	}
	return reqStructDef, respStructDef
}

func (sd structDef) generate(dir string) error {
	fileName := dir + "/" + sd.Name + ".go"
	if err := generateFile(sd, structTemplateLocation, "struct_template", fileName, false); err != nil {
		return err
	}
	return runGoFmt(fileName)
}

//GenerateStruct generate struct
func GenerateStruct(apiDefinition *raml.APIDefinition, dir string) error {
	if err := checkCreateDir(dir); err != nil {
		return err
	}
	for k, v := range apiDefinition.Types {
		sd := newStructDefFromType(v, k)
		if err := sd.generate(dir); err != nil {
			return err
		}
	}
	return nil
}
