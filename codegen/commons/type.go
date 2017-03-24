package commons

import (
	"fmt"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

func IsArray(t interface{}) bool {
	if t == nil {
		return false
	}
	return strings.HasSuffix(fmt.Sprint(t), "[]")
}

func IsBidimensiArray(t interface{}) bool {
	if t == nil {
		return false
	}
	return strings.HasSuffix(fmt.Sprint(t), "[][]")
}

func IsUnion(t interface{}) bool {
	if t == nil {
		return false
	}
	return strings.Index(fmt.Sprint(t), "|") > 0
}

func ArrayType(t interface{}) string {
	return strings.TrimSuffix(fmt.Sprint(t), "[]")
}

func BidimensiArrayType(t interface{}) string {
	return strings.TrimSuffix(fmt.Sprint(t), "[][]")
}

func MultipleInheritance(t string) ([]string, bool) {
	tip := raml.Type{
		Type: t,
	}
	return tip.MultipleInheritance()
}
