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
	switch {
	case commons.IsBidimensiArray(tip): // bidimensional array
		return "[][]" + commons.BidimensiArrayType(tip)
	case commons.IsArray(tip):
		return "[]" + convertToGoType(commons.ArrayType(tip))
	case strings.HasSuffix(tip, "{}"): // map
		return "map[string]" + convertToGoType(tip[:len(tip)-2])
	case strings.Index(tip, "|") > 0:
		return convertUnion(tip)
	}
	return commons.NormalizePkgName(tip)
}
