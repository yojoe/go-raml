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

	// generate oauth2 scope matching middleware
	for _, r := range sd.apiDef.Resources {
		if err := sd.scopeMatching(&r); err != nil {
			return err
		}
	}
	return nil
}

// generate security related code
func generateSecurity(apiDef *raml.APIDefinition, dir, packageName string) error {
	if err := checkCreateDir(dir); err != nil {
		return err
	}

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
	return nil
}

type scopeMatcher struct {
	PackageName string
	Scopes      []string
}

func (sd *securityDef) scopeMatching(res *raml.Resource) error {
	if err := sd.methodScopeMatching(res.Get); err != nil {
		return err
	}
	if err := sd.methodScopeMatching(res.Post); err != nil {
		return err
	}
	if err := sd.methodScopeMatching(res.Put); err != nil {
		return err
	}
	if err := sd.methodScopeMatching(res.Patch); err != nil {
		return err
	}
	if err := sd.methodScopeMatching(res.Delete); err != nil {
		return err
	}
	for _, v := range res.Nested {
		if err := sd.scopeMatching(v); err != nil {
			return err
		}
	}
	return nil
}

func (sd *securityDef) methodScopeMatching(m *raml.Method) error {
	for _, sb := range m.SecuredBy {
		if !validateSecurityScheme(sb.Name, sd.apiDef) { // check if it is valid to generate
			continue
		}

		scopes, ok := sb.Parameters["scopes"] // check if it has scope
		if !ok {
			continue
		}

		// make sure the scopes is an array
		scopesArr, ok := scopes.([]string)
		if !ok {
			return fmt.Errorf("invalid scopes `%v`, it must be array of string", scopes)
		}

		sm := scopeMatcher{
			PackageName: sd.PackageName,
			Scopes:      scopesArr,
		}
		fileName := "oauth2_" + sd.Name + "_" + normalizeScope(strings.Join(scopesArr, "")) + "_middleware.go"
		if err := generateFile(sm, "./templates/oauth2_scope_match.tmpl", "oauth2_scope_match", fileName, false); err != nil {
			return err
		}
	}
	return nil
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
