package codegen

import (
	"path"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

// python representation of a security scheme
type pythonSecurity struct {
	*security
}

// generate security schheme representation in python.
// security scheme is generated as a middleware
func (ps *pythonSecurity) generate(dir string) error {
	fileName := path.Join(dir, "oauth2_"+ps.Name+".py")
	return generateFile(ps, "./templates/oauth2_middleware_python.tmpl", "oauth2_middleware_python", fileName, false)
}

type pythonMiddleware struct {
	ImportPath string
	Name       string
	Args       string
}

func newPythonOauth2Middleware(ss raml.DefinitionChoice) (pythonMiddleware, error) {
	quotedScopes, err := getQuotedSecurityScopes(ss)
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
	return pythonLibImportPath(securitySchemeName(typ), "oauth2_")
}
