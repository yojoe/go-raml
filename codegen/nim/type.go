package nim

import (
	"github.com/Jumpscale/go-raml/codegen/commons"
)

var (
	typeMap = map[string]string{
		"string":    "string",
		"file":      "string",
		"number":    "float",
		"integer":   "int",
		"boolean":   "bool",
		"datetime":  "Time",
		"date-only": "Time",
		"time-only": "Time",
		"int8":      "int8",
		"int16":     "int16",
		"int32":     "int32",
		"int64":     "int64",
		"int":       "int",
		"long":      "int64",
		"float":     "float",
		"double":    "float64",
	}
)

func toNimType(t string) string {
	if v, ok := typeMap[t]; ok {
		return v
	}
	// other types that need some processing
	parents, isMultiple := commons.MultipleInheritance(t)
	switch {
	case commons.IsBidimensiArray(t):
		return "seq[seq[" + toNimType(commons.BidimensiArrayType(t)) + "]]"
	case commons.IsArray(t):
		return "seq[" + toNimType(commons.ArrayType(t)) + "]"
	case isMultiple:
		return multipleInheritanceNewName(parents)
	}
	return t
}
