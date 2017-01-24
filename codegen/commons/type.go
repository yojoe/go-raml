package commons

import (
	"fmt"
	"strings"
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

func ArrayType(t interface{}) string {
	return strings.TrimSuffix(fmt.Sprint(t), "[]")
}

func BidimensiArrayType(t interface{}) string {
	return strings.TrimSuffix(fmt.Sprint(t), "[][]")
}
