package golang

import (
	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

var (
	typeMap = map[string]string{
		"string":  "string",
		"file":    "string",
		"number":  "float64",
		"integer": "int",
		"boolean": "bool",
		"object":  "json.RawMessage",
		"int8":    "int8",
		"int16":   "int16",
		"int32":   "int32",
		"int64":   "int64",
		"int":     "int",
		"long":    "int64",
		"float":   "float64",
		"double":  "float64",
	}
)

// convert from raml type to go type
func convertToGoType(tip, items string) string {
	if v, ok := typeMap[tip]; ok {
		return v
	}
	goramlPkgDir := func() string {
		if globGoramlPkgDir == "" {
			return ""
		}
		return globGoramlPkgDir + "."
	}()
	dateMap := map[string]string{
		"date":          goramlPkgDir + "Date",
		"date-only":     goramlPkgDir + "DateOnly",
		"time-only":     goramlPkgDir + "TimeOnly",
		"datetime-only": goramlPkgDir + "DatetimeOnly",
		"datetime":      goramlPkgDir + "DateTime",
	}

	if v, ok := dateMap[tip]; ok {
		return v
	}

	ramlType := raml.Type{
		Type:  tip,
		Items: items,
	}

	// other types that need some processing
	switch {
	case ramlType.IsBidimensiArray():
		return "[][]" + ramlType.BidimensiArrayType()
	case ramlType.IsArray():
		return "[]" + convertToGoType(ramlType.ArrayType(), "")
	case ramlType.IsUnion():
		return unionNewName(tip)
	case ramlType.IsMultipleInheritance():
		parents, _ := ramlType.MultipleInheritance()
		return multipleInheritanceNewName(parents)
	}
	return commons.NormalizePkgName(tip)
}
