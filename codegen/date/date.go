// Package date provides implementation of various RAML Date type.
// code generator will use code in this package to generate RAML date code.
package date

import (
	"fmt"
	"strings"
)

// Get gets Go code of a specific RAML Date type
// The code returned is without `package date` line
func Get(typ, format string) ([]byte, error) {
	switch typ {
	case "date-only":
		return get("date_only.go")
	case "time-only":
		return get("time_only.go")
	case "datetime-only":
		return get("datetime_only.go")
	case "datetime":
		if format == "" || strings.ToUpper(format) == "RFC3339" {
			return get("datetime.go")
		} else if strings.ToUpper(format) == "RFC2616" {
			return get("datetime_rfc2616.go")
		}
	}
	return []byte{}, fmt.Errorf("unrecognized combination of type :%v format : %v", typ, format)
}

// returns []byte representation of a file without it's first line
func get(filename string) ([]byte, error) {
	b, err := Asset(filename)
	if err != nil {
		return []byte{}, err
	}

	// find first line
	var idx int
	for i, c := range b {
		if string(c) == "\n" {
			idx = i
			break
		}
	}
	if idx == 0 || idx == len(b) {
		return []byte{}, fmt.Errorf("invalid file format")
	}

	// get the rest
	return b[idx+1:], nil
}
