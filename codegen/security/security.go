package security

import (
	"fmt"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

const (
	// Oauth2 string
	Oauth2 = "OAuth 2.0"
)

// security define a security scheme, we only support oauth2 now.
// we generate middleware that checking for oauth2 credential
type Security struct {
	*raml.SecurityScheme
	Name        string
	PackageName string
	Header      *raml.Header
	QueryParams *raml.NamedParameter
	//apiDef      *raml.APIDefinition
}

// create security struct
func New(ss *raml.SecurityScheme, name, packageName string) Security {
	sd := Security{
		SecurityScheme: ss,
	}
	sd.Name = SecuritySchemeName(name)
	sd.PackageName = packageName

	// assign header, if any
	for k, v := range sd.DescribedBy.Headers {
		sd.Header = &v
		sd.Header.Name = string(k)
		break
	}

	// assign query params if any
	for k, v := range sd.DescribedBy.QueryParameters {
		sd.QueryParams = &v
		sd.QueryParams.Name = string(k)
		break
	}

	return sd
}

// get array of security scopes in the form of quoted string
func GetQuotedScopes(ss raml.DefinitionChoice) ([]string, error) {
	var quoted []string
	scopes, err := getSecurityScopes(ss)
	if err != nil {
		return quoted, err
	}
	for _, s := range scopes {
		quoted = append(quoted, fmt.Sprintf(`"%v"`, s))
	}
	return quoted, nil
}

// get scopes of a security scheme as []string
func getSecurityScopes(ss raml.DefinitionChoice) ([]string, error) {
	scopes := []string{}

	// check if there is scopes
	v, ok := ss.Parameters["scopes"]
	if !ok {
		return scopes, nil
	}

	// cast it to []string, return error if failed
	scopesArr, ok := v.([]interface{})
	if !ok {
		return scopes, fmt.Errorf("scopes must be array")
	}

	// build []string
	for _, s := range scopesArr {
		scopes = append(scopes, s.(string))
	}
	return scopes, nil
}

// return security scheme name that could be used in code
func SecuritySchemeName(name string) string {
	return strings.Replace(name, " ", "", -1)
}

// validate security scheme:
// - not empty
// - not 'null'
// - oauth2 -> we only support oauth2 now
func ValidateScheme(name string, apiDef *raml.APIDefinition) bool {
	if name == "" || name == "null" {
		return false
	}
	if ss, ok := apiDef.GetSecurityScheme(name); ok {
		return ss.Type == Oauth2
	}
	return false
}

// find resource's securedBy recursively
func FindResourceSecuredBy(r *raml.Resource) []raml.DefinitionChoice {
	if len(r.SecuredBy) > 0 {
		return r.SecuredBy
	}
	if r.Parent == nil {
		return []raml.DefinitionChoice{}
	}
	return FindResourceSecuredBy(r.Parent)
}
