package nim

import (
	"github.com/Jumpscale/go-raml/codegen/commons"
)

var (
	typeMap = map[string]string{
		"string":    "string",
		"file":      "string",
		"number":    "float64",
		"integer":   "int",
		"boolean":   "bool",
		"datetime":  "Time",
		"date-only": "Time",
		"time-only": "Time",
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
