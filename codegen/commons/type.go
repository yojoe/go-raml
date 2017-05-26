package commons

import (
	"github.com/Jumpscale/go-raml/raml"
)

// IsBuiltinType returns true if the given type
// is builtin type
func IsBuiltinType(t interface{}) bool {
	tip := raml.Type{
		Type: t,
	}
	return tip.IsBuiltin()
}
