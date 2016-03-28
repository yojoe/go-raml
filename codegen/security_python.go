package codegen

import (
	"path"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

// go representation of a security scheme
type pythonSecurity struct {
	*security
}

func (ps *pythonSecurity) generate(dir string) error {
	fileName := path.Join(dir, "oauth2_"+ps.Name+".py")
	return generateFile(ps, "./templates/oauth2_middleware_python.tmpl", "oauth2_middleware_python", fileName, false)
}

type pythonMiddleware struct {
	Name string
	Args string
}

func newPythonOauth2Middleware(ss raml.DefinitionChoice) (pythonMiddleware, error) {
	quotedScopes, err := getQuotedSecurityScopes(ss)
	if err != nil {
		return pythonMiddleware{}, err
	}
	return pythonMiddleware{
		Name: "oauth2_" + securitySchemeName(ss.Name),
		Args: strings.Join(quotedScopes, ", "),
	}, nil
}
