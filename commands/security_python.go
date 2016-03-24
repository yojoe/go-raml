package commands

import (
	"path"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

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

func (sd *securityDef) generatePython(dir string) error {
	// we only support oauth2
	if sd.Type != Oauth2 {
		return nil
	}

	// generate oauth token checking middleware
	fileName := path.Join(dir, "oauth2_"+sd.Name+".py")
	if err := generateFile(sd, "./templates/oauth2_middleware_python.tmpl", "oauth2_middleware_python", fileName, false); err != nil {
		return err
	}
	return nil
}
