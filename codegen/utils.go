package codegen

import (
	"os"
	"regexp"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

var (
	regex          = regexp.MustCompile("({{1}[\\w\\s]+}{1})")
	regNonAlphanum = regexp.MustCompile("[^A-Za-z0-9]+")
)

// _getResourceParams is the recursive function of getResourceParams
func _getResourceParams(r *raml.Resource, params []string) []string {
	if r == nil {
		return params
	}

	matches := regex.FindAllString(r.URI, -1)
	for _, v := range matches {
		params = append(params, v[1:len(v)-1])
	}

	return _getResourceParams(r.Parent, params)
}

// get all params of a resource
// examples:
// /users  							  : no params
// /users/{userId}					  : params 1 = userId
// /users/{userId}/address/{addressId : params 1= userId, param 2= addressId
func getResourceParams(r *raml.Resource) []string {
	params := []string{}
	return _getResourceParams(r, params)
}

// create parameterized URI
// Input : raw string, ex : /users/{userId}/address/{addressId}
// Output : "/users/"+userId+"/address/"+addressId
func paramizingURI(URI string) string {
	uri := `"` + URI + `"`
	// replace { with "+
	uri = strings.Replace(uri, "{", `"+`, -1)

	// if ended with }/" or }", remove trailing "
	if strings.HasSuffix(uri, `}/"`) || strings.HasSuffix(uri, `}"`) {
		uri = uri[:len(uri)-1]
	}

	// replace } with +"
	uri = strings.Replace(uri, "}", `+"`, -1)

	// clean trailing +"
	if strings.HasSuffix(uri, `+"`) {
		uri = uri[:len(uri)-2]
	}
	return uri
}

// create directory if not exist
func checkCreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0777)
	}
	return nil
}

// convert interface type to string
// example :
// 1. string type, result would be string
// 2. []interface{} type, result would be array of string. ex: a,b,c
// Please add other type as needed
func interfaceToString(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	case []interface{}:
		interfaceArr := data.([]interface{})
		resultStr := ""
		for _, v := range interfaceArr {
			resultStr += interfaceToString(v) + ","
		}
		return resultStr[:len(resultStr)-1]
	default:
		return ""
	}
}

// replace non alphanumerics with "_"
func replaceNonAlphanumerics(s string) string {
	return strings.Trim(regNonAlphanum.ReplaceAllString(s, "_"), "_")
}
