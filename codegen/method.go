package codegen

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

type methodInterface interface {
	Verb() string
	Resource() *raml.Resource
	EndpointStr() string
}

// Method defines base Method struct
type method struct {
	*raml.Method
	MethodName   string
	Endpoint     string
	verb         string
	ReqBody      string         // request body type
	RespBody     string         // response body type
	ResourcePath string         // normalized resource path
	resource     *raml.Resource // resource object of this method
	Params       string         // methods params
	FuncComments []string
	SecuredBy    []raml.DefinitionChoice
}

func (m method) Verb() string {
	return m.verb
}

func (m method) Resource() *raml.Resource {
	return m.resource
}

func (m method) EndpointStr() string {
	return m.Endpoint
}

func newMethod(r *raml.Resource, rd *resourceDef, m *raml.Method, methodName string) method {
	method := method{
		Method:   m,
		Endpoint: r.FullURI(),
		verb:     strings.ToUpper(methodName),
		resource: r,
	}

	// set request body
	method.ReqBody = setBodyName(m.Bodies, normalizeURITitle(method.Endpoint)+methodName, reqBodySuffix)

	//set response body
	for k, v := range m.Responses {
		if k >= 200 && k < 300 {
			method.RespBody = setBodyName(v.Bodies, normalizeURITitle(method.Endpoint)+methodName, respBodySuffix)
		}
	}

	// set func comment
	if len(m.Description) > 0 {
		method.FuncComments = commentBuilder(m.Description)
	}

	return method
}

// create server resource's method
func newServerMethod(apiDef *raml.APIDefinition, r *raml.Resource, rd *resourceDef, m *raml.Method,
	methodName, lang string) methodInterface {

	method := newMethod(r, rd, m, methodName)

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
			method: &method,
		}
		gm.setup(apiDef, r, rd, methodName)
		return gm
	case langPython:
		pm := pythonServerMethod{
			method: &method,
		}
		pm.setup(apiDef, r, rd)
		return pm
	default:
		panic("invalid language:" + lang)
	}
}

// create client resource's method
func newClientMethod(r *raml.Resource, rd *resourceDef, m *raml.Method, methodName, lang string) (methodInterface, error) {
	method := newMethod(r, rd, m, methodName)

	method.ResourcePath = paramizingURI(method.Endpoint)

	name := normalizeURITitle(method.Endpoint)

	method.ReqBody = setBodyName(m.Bodies, name+methodName, "ReqBody")

	switch lang {
	case langGo:
		gcm := goClientMethod{method: &method}
		err := gcm.setup(methodName)
		return gcm, err
	case langPython:
		pcm := pythonClientMethod{method: method}
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
//	- use prefix+suffix
func setBodyName(bodies raml.Bodies, prefix, suffix string) string {
	var bodiesType string

	if len(bodies.Type) > 0 && bodies.Type != "object" {
		bodiesType = convertToGoType(bodies.Type)
	} else if bodies.ApplicationJSON != nil {
		if bodies.ApplicationJSON.Type != "" && bodies.ApplicationJSON.Type != "object" {
			bodiesType = convertToGoType(bodies.ApplicationJSON.Type)
		} else {
			bodiesType = prefix + suffix
		}
	}

	return bodiesType
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

type byEndpoint []methodInterface

func (b byEndpoint) Len() int      { return len(b) }
func (b byEndpoint) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byEndpoint) Less(i, j int) bool {
	return strings.Compare(b[i].EndpointStr(), b[j].EndpointStr()) < 0
}
