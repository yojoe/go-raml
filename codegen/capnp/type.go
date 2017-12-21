package capnp

import (
	"strings"

	"github.com/pinzolo/casee"
)

var (
	typeMap = map[string]string{
		"array":     "List",
		"string":    "Text",
		"file":      "Data",
		"object":    "Data",
		"number":    "Float64",
		"integer":   "Int64",
		"boolean":   "Bool",
		"datetime":  "Text",
		"date-only": "Text",
		"time-only": "Text",
		"int8":      "Int8",
		"int16":     "Int16",
		"int32":     "Int32",
		"int64":     "Int64",
		"uint8":     "UInt8",
		"uint16":    "UInt16",
		"uint32":    "UInt32",
		"uint64":    "UInt64",
		"int":       "Int16",
		"float":     "Float64",
		"long":      "Int32",
		"double":    "Float64",
	}
)

func builtinType(t string) (string, bool) {
	if v, ok := typeMap[t]; ok {
		return v, ok
	}
	return casee.ToPascalCase(t), false
}

func toCapnpType(t, capnpType string, itemsType string) (string, string) {
	t = strings.TrimSpace(t)
	capnpType = strings.TrimSpace(capnpType)

	if capnpType != "" { // there is hint in the RAML file
		return capnpType, ""
	}

	v, ok := builtinType(t)
	if ok {
		// if the raml type is array, get the items type
		if t == "array" {
			// if not type is defined, return Data type
			if itemsType == "" {
				return v, "Data"
			}
			itemsType, _ := builtinType(itemsType)
			return v, itemsType
		}
		return v, ""
	}

	// other types that need some processing
	switch {
	case strings.HasSuffix(t, "[]"): // array
		v, _ := builtinType(t[:len(t)-2])
		return "List", v
	}

	return casee.ToPascalCase(t), ""
}
