package codegen

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

// security define a security scheme, we only support oauth2 now.
// we generate middleware that checking for oauth2 credential
type security struct {
	*raml.SecurityScheme
	Name        string
	PackageName string
	Header      *raml.Header
	QueryParams *raml.NamedParameter
	apiDef      *raml.APIDefinition
}

// create security struct
func newSecurity(apiDef *raml.APIDefinition, ss *raml.SecurityScheme, name, packageName string) security {
	sd := security{
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

// go representation of a security scheme
type goSecurity struct {
	*security
}

// generate Go representation of a security scheme
// it implemented as struct based middleware
func (gs *goSecurity) generate(dir string) error {
	fileName := path.Join(dir, "oauth2_"+gs.Name+"_middleware.go")
	return generateFile(gs, "./templates/oauth2_middleware.tmpl", "oauth2_middleware", fileName, false)
}

type pythonSecurity struct {
	*security
}

func (ps *pythonSecurity) generate(dir string) error {
	fileName := path.Join(dir, "oauth2_"+ps.Name+".py")
	return generateFile(ps, "./templates/oauth2_middleware_python.tmpl", "oauth2_middleware_python", fileName, false)
}

// generate security related code
func generateSecurity(apiDef *raml.APIDefinition, dir, packageName, lang string) error {
	var err error

	// generate oauth2 middleware
	for _, v := range apiDef.SecuritySchemes {
		for k, ss := range v {
			if ss.Type != Oauth2 {
				continue
			}

			sd := newSecurity(apiDef, &ss, k, packageName)

			switch lang {
			case langGo:
				gss := goSecurity{security: &sd}
				err = gss.generate(dir)

			case langPython:
				pss := pythonSecurity{security: &sd}
				err = pss.generate(dir)
			}
			if err != nil {
				log.Errorf("generateSecurity() failed to generate %v, err=%v", k, err)
				return err
			}
		}
	}
	return nil
}

// get oauth2 middleware handler from a security scheme
func getOauth2MwrHandler(ss raml.DefinitionChoice) (string, error) {
	quotedScopes, err := getQuotedSecurityScopes(ss)
	if err != nil {
		return "", err
	}
	scopesArgs := strings.Join(quotedScopes, ", ")
	return fmt.Sprintf(`newOauth2%vMiddleware([]string{%v}).Handler`, securitySchemeName(ss.Name), scopesArgs), nil
}

// get array of security scopes in the form of quoted string
func getQuotedSecurityScopes(ss raml.DefinitionChoice) ([]string, error) {
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
