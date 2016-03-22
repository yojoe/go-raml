package commands

import (
	"fmt"
	"path"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
)

const (
	// Oauth2 string
	Oauth2 = "OAuth 2.0"
)

// securityDef define a security scheme, we only support oauth2 now.
// we generate middleware that checking for oauth2 credential
type securityDef struct {
	Name string
	*raml.SecurityScheme
	PackageName string
	Header      *raml.Header
	QueryParams *raml.NamedParameter
	apiDef      *raml.APIDefinition
}

// create securityDef object
func newSecurityDef(apiDef *raml.APIDefinition, ss *raml.SecurityScheme, name, packageName string) securityDef {
	sd := securityDef{
		SecurityScheme: ss,
		apiDef:         apiDef,
	}
	sd.Name = securitySchemeName(name)
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

func (sd *securityDef) generate(dir string) error {
	// we only support oauth2
	if sd.Type != Oauth2 {
		return nil
	}

	// generate oauth token checking middleware
	fileName := path.Join(dir, "oauth2_"+sd.Name+".go")
	if err := generateFile(sd, "./templates/oauth2_middleware.tmpl", "oauth2_middleware", fileName, true); err != nil {
		return err
	}

	return nil
}

// generate security related code
func generateSecurity(apiDef *raml.APIDefinition, dir, packageName string) error {
	if err := checkCreateDir(dir); err != nil {
		return err
	}

	// generate oauth2 middleware
	for _, v := range apiDef.SecuritySchemes {
		for k, ss := range v {
			if ss.Type != Oauth2 {
				continue
			}
			sd := newSecurityDef(apiDef, &ss, k, packageName)
			if err := sd.generate(dir); err != nil {
				log.Errorf("generateSecurity() failed to generate %v, err=%v", k, err)
				return err
			}
		}
	}
	// generate oauth2 scope matching middleware
	log.Infof("generate oauth2_scope_match middleware")
	for _, r := range apiDef.Resources {
		if err := generateScopeMatching(apiDef, &r, packageName, dir); err != nil {
			return err
		}
	}
	return nil
}

type scopeMatcher struct {
	PackageName string
	Scopes      string
	Name        string
}

func newScopeMatcher(oauth2Name, packageName string, scopes []string) scopeMatcher {
	quoted := make([]string, 0, len(scopes))
	for _, s := range scopes {
		quoted = append(quoted, fmt.Sprintf(`"%v"`, s))
	}
	return scopeMatcher{
		Name:        oauth2Name + "_" + normalizeScope(strings.Join(scopes, "")),
		PackageName: packageName,
		Scopes:      strings.Join(quoted, ", "),
	}
}

func generateScopeMatching(apiDef *raml.APIDefinition, res *raml.Resource, packageName, dir string) error {
	if err := methodScopeMatching(apiDef, res.Get, packageName, dir); err != nil {
		return err
	}
	if err := methodScopeMatching(apiDef, res.Post, packageName, dir); err != nil {
		return err
	}
	if err := methodScopeMatching(apiDef, res.Put, packageName, dir); err != nil {
		return err
	}
	if err := methodScopeMatching(apiDef, res.Patch, packageName, dir); err != nil {
		return err
	}
	if err := methodScopeMatching(apiDef, res.Delete, packageName, dir); err != nil {
		return err
	}
	for _, v := range res.Nested {
		if err := generateScopeMatching(apiDef, v, packageName, dir); err != nil {
			return err
		}
	}
	return nil
}

func methodScopeMatching(apiDef *raml.APIDefinition, m *raml.Method, packageName, dir string) error {
	if m == nil {
		return nil
	}
	for _, sb := range m.SecuredBy {
		if !validateSecurityScheme(sb.Name, apiDef) { // check if it is valid to generate
			continue
		}

		scopes, err := getSecurityScopes(sb)
		if err != nil {
			return err
		}
		if len(scopes) == 0 {
			continue
		}

		sm := newScopeMatcher(sb.Name, packageName, scopes)
		fileName := "oauth2_" + sb.Name + "_" + normalizeScope(strings.Join(scopes, "")) + ".go"
		fileName = path.Join(dir, fileName)
		if err := generateFile(sm, "./templates/oauth2_scopes_match.tmpl", "oauth2_scopes_match", fileName, false); err != nil {
			return err
		}
	}
	return nil
}

// get scopes of a security scheme as []string
func getSecurityScopes(ss raml.DefinitionChoice) ([]string, error) {
	scopes := []string{}

	// check if there is scopes
	v, ok := ss.Parameters["scopes"]
	if !ok {
		return scopes, nil
	}
	scopesArr, ok := v.([]interface{})
	if !ok {
		return scopes, fmt.Errorf("scopes must be array")
	}
	for _, s := range scopesArr {
		scopes = append(scopes, s.(string))
	}
	return scopes, nil
}

// return security scheme name that could be used in code
func securitySchemeName(name string) string {
	return strings.Replace(name, " ", "", -1)
}

// validate security scheme:
// - not empty
// - not 'null'
// - oauth2 -> we only support oauth2 now
func validateSecurityScheme(name string, apiDef *raml.APIDefinition) bool {
	if name == "" || name == "null" {
		return false
	}
	for _, v := range apiDef.SecuritySchemes {
		if ss, ok := v[name]; ok {
			return ss.Type == Oauth2
		}
	}
	return false
}

// TODO : make it only alphanumeric
func normalizeScope(s string) string {
	ret := strings.Replace(s, " ", "", -1)
	ret = strings.Replace(ret, ":", "", -1)
	return ret
}
