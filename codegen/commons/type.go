package commons

import (
	"fmt"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

func IsArray(t interface{}) bool {
	tip := raml.Type{
		Type: t,
	}
	return tip.IsArray()
}

func IsBidimensiArray(t interface{}) bool {
	tip := raml.Type{
		Type: t,
	}
	return tip.IsBidimensiArray()
}

func IsUnion(t interface{}) bool {
	tip := raml.Type{
		Type: t,
	}
	return tip.IsUnion()
}

func UnionTypes(t interface{}) []string {
	tip := raml.Type{
		Type: t,
	}
	uts, ok := tip.Union()
	if !ok {
		return []string{}
	}
	return uts
}

func ArrayType(t interface{}) string {
	return strings.TrimSuffix(fmt.Sprint(t), "[]")
}

func BidimensiArrayType(t interface{}) string {
	return strings.TrimSuffix(fmt.Sprint(t), "[][]")
}

func MultipleInheritance(t interface{}) ([]string, bool) {
	tip := raml.Type{
		Type: t,
	}
	return tip.MultipleInheritance()
}

func IsMultipleInheritance(t interface{}) bool {
	_, ok := MultipleInheritance(t)
	return ok
}

func IsBuiltinType(t interface{}) bool {
	tip := raml.Type{
		Type: t,
	}
	return tip.IsBuiltin()
}
