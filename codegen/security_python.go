package codegen

import (
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
