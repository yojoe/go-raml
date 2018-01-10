package python

var (
	typeMap = map[string]string{
		"string":   "string_types",
		"integer":  "int",
		"int":      "int",
		"int8":     "int",
		"int16":    "int",
		"int32":    "int",
		"int64":    "int",
		"long":     "int",
		"number":   "float",
		"double":   "float",
		"float":    "float",
		"boolean":  "bool",
		"datetime": "datetime",
		"object":   "dict",
		"UUID":     "UUID",
	}
)

type pythonType struct {
	name         string
	importModule string
	importName   string
}

// convert raml type to python type
// also returns the needed import
func toPythonType(t string) *pythonType {
	tip, ok := typeMap[t]
	if !ok {
		return nil
	}
	pt := pythonType{
		name: tip,
	}

	switch t {
	case "datetime":
		pt.importModule = "datetime"
		pt.importName = "datetime"
	case "uuid":
		pt.importModule = "uuid"
		pt.importName = "UUID"
	case "string":
		pt.importModule = "six"
		pt.importName = "string_types"
	}
	return &pt
}
