package golang

import (
	"fmt"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/codegen/security"
	"github.com/Jumpscale/go-raml/raml"
)

type goServerMethod struct {
	*resource.Method
	Middlewares string
}

// setup go server method, initializes all needed variables
func (gm *goServerMethod) setup(apiDef *raml.APIDefinition, r *raml.Resource, rd *resource.Resource, methodName string) error {
	// set method name
	name := commons.NormalizeURI(gm.Endpoint)
	if len(gm.DisplayName) > 0 {
		gm.MethodName = commons.DisplayNameToFuncName(gm.DisplayName)
	} else {
		gm.MethodName = name[len(rd.Name):] + methodName
	}

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
func (gm goServerMethod) libImported(rootImportPath string) map[string]struct{} {
	libs := map[string]struct{}{}

	// req body
	if lib := libImportPath(rootImportPath, gm.ReqBody); lib != "" {
		libs[lib] = struct{}{}
	}
	// resp body
	if lib := libImportPath(rootImportPath, gm.RespBody); lib != "" {
		libs[lib] = struct{}{}
	}
	return libs
}

type goClientMethod struct {
	*resource.Method
}

// create client resource's method
func newGoClientMethod(r *raml.Resource, rd *resource.Resource, m *raml.Method, methodName, lang string) (resource.MethodInterface, error) {
	method := resource.NewMethod(r, rd, m, methodName, setBodyName)

	method.ResourcePath = commons.ParamizingURI(method.Endpoint, "+")

	name := commons.NormalizeURITitle(method.Endpoint)

	method.ReqBody = setBodyName(m.Bodies, name+methodName, "ReqBody")

	gcm := goClientMethod{Method: &method}
	err := gcm.setup(methodName)
	return gcm, err
}

func (gcm *goClientMethod) setup(methodName string) error {
	// build func/method params
	buildParams := func(r *raml.Resource, bodyType string) (string, error) {
		paramsStr := strings.Join(resource.GetResourceParams(r), ",")
		if len(paramsStr) > 0 {
			paramsStr += " string"
		}

		// append request body type
		if len(bodyType) > 0 {
			if len(paramsStr) > 0 {
				paramsStr += ", "
			}
			paramsStr += strings.ToLower(bodyType) + " " + bodyType
		}

		// append header
		if len(paramsStr) > 0 {
			paramsStr += ","
		}
		paramsStr += "headers,queryParams map[string]interface{}"

		return paramsStr, nil
	}

	// method name
	name := commons.NormalizeURITitle(gcm.Endpoint)

	if len(gcm.DisplayName) > 0 {
		gcm.MethodName = commons.DisplayNameToFuncName(gcm.DisplayName)
	} else {
		gcm.MethodName = strings.Title(name + methodName)
	}

	// method param
	methodParam, err := buildParams(gcm.RAMLResource, gcm.ReqBody)
	if err != nil {
		return err
	}
	gcm.Params = methodParam

	return nil
}

// ReturnTypes returns all types returned by this method
func (gcm goClientMethod) ReturnTypes() string {
	var types []string
	if gcm.RespBody != "" {
		types = append(types, gcm.RespBody)
	}
	types = append(types, []string{"*http.Response", "error"}...)

	return fmt.Sprintf("(%v)", strings.Join(types, ","))
}

func (gcm goClientMethod) libImported(rootImportPath string) map[string]struct{} {
	libs := map[string]struct{}{}

	// req body
	if lib := libImportPath(rootImportPath, gcm.ReqBody); lib != "" {
		libs[lib] = struct{}{}
	}
	// resp body
	if lib := libImportPath(rootImportPath, gcm.RespBody); lib != "" {
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
	methodName, lang string) resource.MethodInterface {

	method := resource.NewMethod(r, rd, m, methodName, setBodyName)

	// security scheme
	if len(m.SecuredBy) > 0 {
		method.SecuredBy = m.SecuredBy
	} else if sb := security.FindResourceSecuredBy(r); len(sb) > 0 {
		method.SecuredBy = sb
	} else {
		method.SecuredBy = apiDef.SecuredBy // use secured by from root document
	}

	gm := goServerMethod{
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
		tipe = convertToGoType(bodies.Type)
	} else if bodies.ApplicationJSON != nil {
		if bodies.ApplicationJSON.Type != "" && bodies.ApplicationJSON.Type != "object" {
			tipe = convertToGoType(bodies.ApplicationJSON.Type)
		} else {
			tipe = prefix + suffix
		}
	}

	if commons.IsJSONString(tipe) {
		tipe = prefix + suffix
	}

	return tipe
}
