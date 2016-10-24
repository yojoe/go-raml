package codegen

import (
	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

// create server resource's method
func newServerMethod(apiDef *raml.APIDefinition, r *raml.Resource, rd *resource.Resource, m *raml.Method,
	methodName, lang string) resource.MethodInterface {

	method := resource.NewMethod(r, rd, m, methodName, setBodyName)

	// security scheme
	if len(m.SecuredBy) > 0 {
		method.SecuredBy = m.SecuredBy
	} else if sb := findResourceSecuredBy(r); len(sb) > 0 {
		method.SecuredBy = sb
	} else {
		method.SecuredBy = apiDef.SecuredBy // use secured by from root document
	}

	switch lang {
	case langGo:
		gm := goServerMethod{
			Method: &method,
		}
		gm.setup(apiDef, r, rd, methodName)
		return gm
	case langPython:
		pm := pythonServerMethod{
			Method: &method,
		}
		pm.setup(apiDef, r, rd)
		return pm
	default:
		panic("invalid language:" + lang)
	}
}

// create client resource's method
func newClientMethod(r *raml.Resource, rd *resource.Resource, m *raml.Method, methodName, lang string) (resource.MethodInterface, error) {
	method := resource.NewMethod(r, rd, m, methodName, setBodyName)

	method.ResourcePath = commons.ParamizingURI(method.Endpoint, "+")

	name := commons.NormalizeURITitle(method.Endpoint)

	method.ReqBody = setBodyName(m.Bodies, name+methodName, "ReqBody")

	switch lang {
	case langGo:
		gcm := goClientMethod{Method: &method}
		err := gcm.setup(methodName)
		return gcm, err
	case langPython:
		pcm := pythonClientMethod{Method: method}
		pcm.setup()
		return pcm, nil
	default:
		panic("invalid language:" + lang)

	}
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

// find resource's securedBy recursively
func findResourceSecuredBy(r *raml.Resource) []raml.DefinitionChoice {
	if len(r.SecuredBy) > 0 {
		return r.SecuredBy
	}
	if r.Parent == nil {
		return []raml.DefinitionChoice{}
	}
	return findResourceSecuredBy(r.Parent)
}
