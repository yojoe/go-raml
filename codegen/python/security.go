package python

import (
	"path"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/security"
	"github.com/Jumpscale/go-raml/raml"
)

// python representation of a security scheme
type pythonSecurity struct {
	*security.Security
}

// generate security related code
func generateSecurity(schemes map[string]raml.SecurityScheme, dir string) error {
	var err error

	// generate oauth2 middleware
	for k, ss := range schemes {
		if ss.Type != security.Oauth2 { // only support oauth2 now
			continue
		}

		sd := security.New(&ss, k, "")

		pss := pythonSecurity{Security: &sd}
		err = pss.generate(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

// generate security schheme representation in python.
// security scheme is generated as a middleware
func (ps *pythonSecurity) generate(dir string) error {
	fileName := path.Join(dir, "oauth2_"+ps.Name+".py")
	return commons.GenerateFile(ps, "./templates/oauth2_middleware_python.tmpl", "oauth2_middleware_python", fileName, false)
}

type middleware struct {
	ImportPath string
	Name       string
	Args       string
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
