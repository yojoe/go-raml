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

func (sd *securityDef) generateGo(dir string) error {
	// generate oauth token checking middleware
	fileName := path.Join(dir, "oauth2_"+sd.Name+"_middleware.go")
	if err := generateFile(sd, "./templates/oauth2_middleware.tmpl", "oauth2_middleware", fileName, false); err != nil {
		return err
	}

	return nil
}

func (sd *securityDef) generate(lang, dir string) error {
	switch lang {
	case langGo:
		return sd.generateGo(dir)
	case langPython:
		return sd.generatePython(dir)
	default:
		return fmt.Errorf("invalid language :%v", lang)
	}
}

// generate security related code
func generateSecurity(apiDef *raml.APIDefinition, dir, packageName, lang string) error {
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
			if err := sd.generate(lang, dir); err != nil {
				log.Errorf("generateSecurity() failed to generate %v, err=%v", k, err)
				return err
			}
		}
	}

	if lang != langPython {
		return nil
	}
	log.Info("generate python security")
	return generatePythonSecurity(apiDef, packageName, dir)
}

// newOauthMiddleware(header, field, scopes).Handler
func getOauth2MwrHandler(ss raml.DefinitionChoice) (string, error) {
	quotedScopes, err := getQuotedSecurityScopes(ss)
	if err != nil {
		return "", err
	}
	scopesArgs := strings.Join(quotedScopes, ", ")
	return fmt.Sprintf(`newOauth2%vMiddleware([]string{%v}).Handler`, securitySchemeName(ss.Name), scopesArgs), nil
}

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
