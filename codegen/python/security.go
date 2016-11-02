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

type pythonMiddleware struct {
	ImportPath string
	Name       string
	Args       string
}

func newPythonOauth2Middleware(ss raml.DefinitionChoice) (pythonMiddleware, error) {
	quotedScopes, err := security.GetQuotedScopes(ss)
	if err != nil {
		return pythonMiddleware{}, err
	}

	importPath, name := pythonOauth2libImportPath(ss.Name)
	return pythonMiddleware{
		ImportPath: importPath,
		Name:       name,
		Args:       strings.Join(quotedScopes, ", "),
	}, nil
}

// get library import path from a type
func pythonOauth2libImportPath(typ string) (string, string) {
	return pythonLibImportPath(security.SecuritySchemeName(typ), "oauth2_")
}
