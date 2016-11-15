package capnp

import (
	"fmt"
	"strings"
)

var (
	typeMap = map[string]string{
		"string": "Text",
		//"file":      "string",
		"number":  "Float64",
		"integer": "Int64",
		"boolean": "Bool",
		//"datetime":  "Time",
		//"date-only": "Time",
		//"time-only": "Time",
	}
)

func toCapnpType(t, capnpType string) string {
	t = strings.TrimSpace(t)
	capnpType = strings.TrimSpace(capnpType)

	if capnpType != "" { // there is hint in the RAML file
		return capnpType
	}

	if v, ok := typeMap[t]; ok {
		return v
	}
	// other types that need some processing
	switch {
	case strings.HasSuffix(t, "[]"): // array
		return fmt.Sprintf("List(%v)", toCapnpType(t[:len(t)-2], ""))
	}

	return t
}
