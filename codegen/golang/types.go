package golang

import (
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
)

var (
	typeMap = map[string]string{
		"string":  "string",
		"file":    "string",
		"number":  "float64",
		"integer": "int",
		"boolean": "bool",
		"object":  "json.RawMessage",
	}
)

func convertUnion(strType string) string {
	if strings.Index(strType, "[]") > 0 {
		return "[]interface{}"
	}
	return "interface{}"
}

// convert from raml type to go type
func convertToGoType(tip string) string {
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

	// other types that need some processing
	parents, isMultiple := commons.MultipleInheritance(tip)
	switch {
	case commons.IsBidimensiArray(tip):
		return "[][]" + commons.BidimensiArrayType(tip)
	case commons.IsArray(tip):
		return "[]" + convertToGoType(commons.ArrayType(tip))
	case commons.IsUnion(tip):
		return unionNewName(tip)
	case isMultiple:
		return multipleInheritanceNewName(parents)
	}
	return commons.NormalizePkgName(tip)
}
