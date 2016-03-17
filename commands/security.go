package commands

import (
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
}

// create securityDef object
func newSecurityDef(ss *raml.SecurityScheme, name, packageName string) securityDef {
	sd := securityDef{
		SecurityScheme: ss,
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
	fileName := path.Join(dir, "oauth2_"+sd.Name+".go")
	return generateFile(sd, "./templates/oauth2_middleware.tmpl", "oauth2_middleware", fileName, true)
}

func generateSecurity(apiDef *raml.APIDefinition, dir, packageName string) error {
	if err := checkCreateDir(dir); err != nil {
		return err
	}

	for _, v := range apiDef.SecuritySchemes {
		for k, ss := range v {
			if ss.Type != Oauth2 {
				continue
			}
			sd := newSecurityDef(&ss, k, packageName)
			if err := sd.generate(dir); err != nil {
				log.Errorf("generateSecurity() failed to generate %v, err=%v", k, err)
				return err
			}
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
