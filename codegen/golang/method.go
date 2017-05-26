package golang

import (
	"fmt"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/codegen/security"
	"github.com/Jumpscale/go-raml/raml"
)

type serverMethod struct {
	*resource.Method
	Middlewares string
}

func serverMethodName(endpoint, displayName, verb, resName string) string {
	if len(displayName) > 0 {
		return commons.DisplayNameToFuncName(displayName)
	} else {
		name := commons.NormalizeURI(endpoint)
		return name[len(resName):] + verb
	}
}

// setup go server method, initializes all needed variables
func (gm *serverMethod) setup(apiDef *raml.APIDefinition, r *raml.Resource, rd *resource.Resource, methodName string) error {
	gm.MethodName = serverMethodName(gm.Endpoint, gm.DisplayName, methodName, rd.Name)

	// setting middlewares
	middlewares := []string{}

	// security middlewares
	for _, v := range gm.SecuredBy {
		if !security.ValidateScheme(v.Name, apiDef) {
			continue
		}
		// oauth2 middleware
		m, err := getOauth2MwrHandler(v)
		if err != nil {
			return err
		}
		middlewares = append(middlewares, m)
	}

	gm.Middlewares = strings.Join(middlewares, ", ")

	return nil
}

// return all libs imported by this method
func (gm serverMethod) libImported(rootImportPath string) map[string]struct{} {
	libs := map[string]struct{}{}

	// req body
	if lib := libImportPath(rootImportPath, gm.ReqBody, globLibRootURLs); lib != "" {
		libs[lib] = struct{}{}
	}
	// resp body
	if lib := libImportPath(rootImportPath, gm.RespBody, globLibRootURLs); lib != "" {
		libs[lib] = struct{}{}
	}
	return libs
}

// Imports return all packages needed
// by this method
func (gm serverMethod) Imports() []string {
	ip := map[string]struct{}{
		"net/http": struct{}{},
	}
	if gm.RespBody != "" || gm.ReqBody != "" {
		ip["encoding/json"] = struct{}{}
	}
	for lib := range gm.libImported(globRootImportPath) {
		ip[lib] = struct{}{}
	}
	return sortImportPaths(ip)
}

type clientMethod struct {
	*resource.Method
}

// true if req body need validation code
func (gm serverMethod) ReqBodyNeedValidation() bool {
	// we can't use t.GetBuiltinType here because
	// the reqBody type is already in Go type
	getBuiltinType := func() string {
		switch {
		case strings.HasPrefix(gm.ReqBody, "[][]"): // bidimensional array
			return strings.TrimPrefix(gm.ReqBody, "[][]")
		case strings.HasPrefix(gm.ReqBody, "[]"): // array
			return strings.TrimPrefix(gm.ReqBody, "[]")
		default:
			return gm.ReqBody
		}
	}
	t := raml.Type{
		Type: getBuiltinType(),
	}

	return !t.IsBuiltin()
}

// create client resource's method
func newGoClientMethod(r *raml.Resource, rd *resource.Resource, m *raml.Method,
	methodName string) (resource.MethodInterface, error) {
	method := resource.NewMethod(r, rd, m, methodName, setBodyName)

	method.ResourcePath = commons.ParamizingURI(method.Endpoint, "+")

	name := commons.NormalizeURITitle(method.Endpoint)

	method.ReqBody = setBodyName(m.Bodies, name+methodName, "ReqBody")

	gcm := clientMethod{Method: &method}
	gcm.setup(methodName)
	return gcm, nil
}

func (gcm *clientMethod) setup(methodName string) {
	// build func/method params
	buildParams := func(r *raml.Resource, bodyType string) string {
		params := resource.GetResourceParams(r)

		if len(params) > 0 {
			// all params has string type
			params[len(params)-1] = params[len(params)-1] + " string"
		}

		// append request body type
		if len(bodyType) > 0 {
			params = append(params, "body "+bodyType)
		}

		// append header
		params = append(params, "headers,queryParams map[string]interface{}")

		return strings.Join(params, ", ")
	}

	// method name
	name := commons.NormalizeURITitle(gcm.Endpoint)

	if len(gcm.DisplayName) > 0 {
		gcm.MethodName = commons.DisplayNameToFuncName(gcm.DisplayName)
	} else {
		gcm.MethodName = name + methodName
	}
	gcm.MethodName = strings.Title(gcm.MethodName)

	// method param
	gcm.Params = buildParams(gcm.RAMLResource, gcm.ReqBody)
}

// ReturnTypes returns all types returned by this method
func (gcm clientMethod) ReturnTypes() string {
	var types []string
	if gcm.RespBody != "" {
		types = append(types, gcm.RespBody)
	}
	types = append(types, []string{"*http.Response", "error"}...)

	return fmt.Sprintf("(%v)", strings.Join(types, ","))
}

func (gcm clientMethod) libImported(rootImportPath string) map[string]struct{} {
	libs := map[string]struct{}{}

	// req body
	if lib := libImportPath(rootImportPath, gcm.ReqBody, globLibRootURLs); lib != "" {
		libs[lib] = struct{}{}
	}
	// resp body
	if lib := libImportPath(rootImportPath, gcm.RespBody, globLibRootURLs); lib != "" {
		libs[lib] = struct{}{}
	}
	return libs
}

// get oauth2 middleware handler from a security scheme
func getOauth2MwrHandler(ss raml.DefinitionChoice) (string, error) {
	// construct security scopes
	quotedScopes, err := security.GetQuotedScopes(ss)
	if err != nil {
		return "", err
	}
	scopesArgs := strings.Join(quotedScopes, ", ")

	// middleware name
	// need to handle case where it reside in different package
	var packageName string
	name := ss.Name

	if splitted := strings.Split(name, "."); len(splitted) == 2 {
		packageName = splitted[0]
		name = splitted[1]
	}
	mwr := fmt.Sprintf(`NewOauth2%vMiddleware([]string{%v}).Handler`, name, scopesArgs)
	if packageName != "" {
		mwr = packageName + "." + mwr
	}
	return mwr, nil
}

// create server resource's method
func newServerMethod(apiDef *raml.APIDefinition, r *raml.Resource, rd *resource.Resource, m *raml.Method,
	methodName string) resource.MethodInterface {

	method := resource.NewMethod(r, rd, m, methodName, setBodyName)

	// security scheme
	method.SecuredBy = security.GetMethodSecuredBy(apiDef, r, m)

	gm := serverMethod{
		Method: &method,
	}
	gm.setup(apiDef, r, rd, methodName)
	return gm
}

// setBodyName set name of method's request/response body.
//
// Rules:
//	- use bodies.Type if not empty and not `object`
//	- use bodies.ApplicationJSON.Type if not empty and not `object`
//	- use prefix+suffix if:
//		- not meet previous rules
//		- previous rules produces JSON string
func setBodyName(bodies raml.Bodies, prefix, suffix string) string {
	var tipe string
	prefix = commons.NormalizeURITitle(prefix)

	if len(bodies.Type) > 0 && bodies.Type != "object" {
		tipe = convertToGoType(bodies.Type, "")
	} else if bodies.ApplicationJSON != nil {
		if bodies.ApplicationJSON.TypeString() != "" && bodies.ApplicationJSON.TypeString() != "object" {
			tipe = convertToGoType(bodies.ApplicationJSON.TypeString(), "")
		} else {
			tipe = prefix + suffix
		}
	}

	if commons.IsJSONString(tipe) {
		tipe = prefix + suffix
	}

	return tipe
}
