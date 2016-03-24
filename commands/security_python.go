package commands

import (
	"fmt"
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

// generate python scope matcher
func generatePythonScopeMatcher(apiDef *raml.APIDefinition, packageName, dir string) error {
	// generate oauth2 scope matching middleware of root document
	if err := securedByScopeMatching(apiDef, apiDef.SecuredBy, packageName, dir); err != nil {
		return err
	}

	// generate oauth2 scope matching middleware for all resource
	for _, r := range apiDef.Resources {
		if err := generateResourceScopeMatcher(apiDef, &r, packageName, dir); err != nil {
			return err
		}
	}
	return nil
}

// scope matcher middleware definition
type scopeMatcher struct {
	PackageName string
	Scopes      string
	Name        string
}

// create scopeMatcher
func newScopeMatcher(oauth2Name, packageName string, scopes []string) scopeMatcher {
	quoted := make([]string, 0, len(scopes))
	for _, s := range scopes {
		quoted = append(quoted, fmt.Sprintf(`"%v"`, s))
	}
	return scopeMatcher{
		Name:        scopeMatcherName(oauth2Name, scopes),
		PackageName: packageName,
		Scopes:      strings.Join(quoted, ", "),
	}
}

// generate scope matcher middleware name from oauth2 security scheme name and scopes
func scopeMatcherName(oauth2Name string, scopes []string) string {
	return securitySchemeName(oauth2Name) + "_" + replaceNonAlphanumerics(strings.Join(scopes, ""))
}

// generate scope matching midleware needed by a resource
func generateResourceScopeMatcher(apiDef *raml.APIDefinition, res *raml.Resource, packageName, dir string) error {
	if err := securedByScopeMatching(apiDef, res.SecuredBy, packageName, dir); err != nil {
		return err
	}

	if err := methodScopeMatching(apiDef, res.Get, packageName, dir); err != nil {
		return err
	}
	if err := methodScopeMatching(apiDef, res.Post, packageName, dir); err != nil {
		return err
	}
	if err := methodScopeMatching(apiDef, res.Put, packageName, dir); err != nil {
		return err
	}
	if err := methodScopeMatching(apiDef, res.Patch, packageName, dir); err != nil {
		return err
	}
	if err := methodScopeMatching(apiDef, res.Delete, packageName, dir); err != nil {
		return err
	}
	for _, v := range res.Nested {
		if err := generateResourceScopeMatcher(apiDef, v, packageName, dir); err != nil {
			return err
		}
	}
	return nil
}

// generate scope matching middleware needed by a method
func methodScopeMatching(apiDef *raml.APIDefinition, m *raml.Method, packageName, dir string) error {
	if m == nil {
		return nil
	}
	return securedByScopeMatching(apiDef, m.SecuredBy, packageName, dir)
}

// generate secure matcher of a SecuredBy field
func securedByScopeMatching(apiDef *raml.APIDefinition, sbs []raml.DefinitionChoice, packageName, dir string) error {
	generatePython := func(sm scopeMatcher) error {
		fileName := path.Join(dir, "oauth2_"+sm.Name+".py")
		return generateFile(sm, "./templates/oauth2_scopes_match_python.tmpl", "oauth2_scopes_match_python", fileName, false)
	}
	for _, sb := range sbs {
		if !validateSecurityScheme(sb.Name, apiDef) { // check if it is valid to generate
			continue
		}

		scopes, err := getSecurityScopes(sb)
		if err != nil {
			return err
		}
		if len(scopes) == 0 {
			continue
		}

		sm := newScopeMatcher(sb.Name, packageName, scopes)
		if err := generatePython(sm); err != nil {
			return err
		}
	}
	return nil
}
