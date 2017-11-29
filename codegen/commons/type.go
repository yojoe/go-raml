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

// GetBasicType returns basic type of a basic or complex type
func GetBasicType(t string) string {
	tip := raml.Type{
		Type: t,
	}

	switch {
	case tip.IsArray():
		return tip.ArrayType()
	case tip.IsBidimensiArray():
		return tip.BidimensiArrayType()
	default:
		return t
	}
}

// IsArrayType returns true if the given
// string is a RAML array
func IsArrayType(t string) bool {
	tip := raml.Type{
		Type: t,
	}
	return tip.IsArray()
}
