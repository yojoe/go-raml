package nim

import (
	"strings"
)

var (
	typeMap = map[string]string{
		"string":   "string",
		"file":     "string",
		"number":   "float64",
		"integer":  "int",
		"boolean":  "bool",
		"datetime": "Time",
	}
)

func toNimType(t string) string {
	if v, ok := typeMap[t]; ok {
		return v
	}
	// other types that need some processing
	switch {
	case strings.HasSuffix(t, "[]"): // array
		return "seq[" + toNimType(t[:len(t)-2]) + "]"
	}

	return t
}
