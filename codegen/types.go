package codegen

import (
	"strings"
)

var (
	typeMap = map[string]string{
		"string":        "string",
		"file":          "string",
		"number":        "float64",
		"integer":       "int",
		"boolean":       "bool",
		"date":          "goraml.Date",
		"date-only":     "goraml.DateOnly",
		"time-only":     "goraml.TimeOnly",
		"datetime-only": "goraml.DatetimeOnly",
		"datetime":      "goraml.DateTime",
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

	// other types that need some processing
	switch {
	case strings.HasSuffix(tip, "[][]"): // bidimensional array
		return "[][]" + convertToGoType(tip[:len(tip)-4])
	case strings.HasSuffix(tip, "[]"): // array
		return "[]" + convertToGoType(tip[:len(tip)-2])
	case strings.HasSuffix(tip, "{}"): // map
		return "map[string]" + convertToGoType(tip[:len(tip)-2])
	case strings.Index(tip, "|") > 0:
		return convertUnion(tip)
	}
	return normalizePkgName(tip)
}
