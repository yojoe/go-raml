package python

import (
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/security"
	"github.com/Jumpscale/go-raml/raml"
)

const (
	middlewareTypeOauth2 = "oauth2"
)

// python representation of a security scheme
type pythonSecurity struct {
	*security.Security
}

func getOauth2Defs(schemes map[string]raml.SecurityScheme) []pythonSecurity {
	var defs []pythonSecurity

	// generate oauth2 middleware
	for k, ss := range schemes {
		if ss.Type != security.Oauth2 { // only support oauth2 now
			continue
		}
		sd := security.New(&ss, k, "")

		defs = append(defs, pythonSecurity{Security: &sd})
	}
	return defs
}

// generate security schheme representation in python.
// security scheme is generated as a middleware
func (ps *pythonSecurity) generate(fileName, tmplFile, tmplName string) error {
	return commons.GenerateFile(ps, tmplFile, tmplName, fileName, false)
}

type middleware struct {
	ImportPath string
	Name       string
	Args       string
	Type       string
}

func newPythonOauth2Middleware(ss raml.DefinitionChoice) (middleware, error) {
	quotedScopes, err := security.GetQuotedScopes(ss)
	if err != nil {
		return middleware{}, err
	}

	importPath, name := pythonOauth2libImportPath(ss.Name)
	return middleware{
		ImportPath: importPath,
		Name:       name,
		Args:       strings.Join(quotedScopes, ", "),
		Type:       middlewareTypeOauth2,
	}, nil
}

// get library import path from a type
func pythonOauth2libImportPath(typ string) (string, string) {
	return libImportPath(security.SecuritySchemeName(typ), "oauth2_")
}

func oauth2ClientName(schemeName string) string {
	return "Oauth2Client" + strings.Title(schemeName)
}

func oauth2ClientFilename(schemeName string) string {
	return "oauth2_client_" + schemeName + ".py"
}
