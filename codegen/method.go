package codegen

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

type methodInterface interface {
	Verb() string
	Resource() *raml.Resource
}

// Method defines base Method struct
type Method struct {
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

func (m Method) Verb() string {
	return m.verb
}

func (m Method) Resource() *raml.Resource {
	return m.resource
}

func newMethod(r *raml.Resource, rd *resourceDef, m *raml.Method, methodName, parentEndpoint, curEndpoint string) Method {
	method := Method{
		Method:   m,
		Endpoint: parentEndpoint + curEndpoint,
		verb:     strings.ToUpper(methodName),
		resource: r,
	}

	// set request body
	method.ReqBody = assignBodyName(m.Bodies, normalizeURITitle(method.Endpoint)+methodName, "ReqBody")

	//set response body
	for k, v := range m.Responses {
		if k >= 200 && k < 300 {
			method.RespBody = assignBodyName(v.Bodies, normalizeURITitle(method.Endpoint)+methodName, "RespBody")
		}
	}

	// set func comment
	if len(m.Description) > 0 {
		method.FuncComments = commentBuilder(m.Description)
	}

	return method
}

func newServerMethod(apiDef *raml.APIDefinition, r *raml.Resource, rd *resourceDef, m *raml.Method,
	methodName, parentEndpoint, curEndpoint, lang string) methodInterface {

	method := newMethod(r, rd, m, methodName, parentEndpoint, curEndpoint)

	// security scheme
	switch {
	case len(m.SecuredBy) > 0: // use secured by from this method
		method.SecuredBy = m.SecuredBy
	case len(r.SecuredBy) > 0: // use securedby from resource
		method.SecuredBy = r.SecuredBy
	default:
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

func newClientMethod(r *raml.Resource, rd *resourceDef, m *raml.Method, methodName, parentEndpoint, curEndpoint, lang string) (methodInterface, error) {
	method := newMethod(r, rd, m, methodName, parentEndpoint, curEndpoint)

	method.ResourcePath = paramizingURI(method.Endpoint)

	name := normalizeURITitle(method.Endpoint)

	method.ReqBody = assignBodyName(m.Bodies, name+methodName, "ReqBody")

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

// assignBodyName assign method's request body by bodies.Type or bodies.ApplicationJson
// if bodiesType generated from bodies.Type we dont need append prefix and suffix
// 		example : bodies.Type = City, so bodiesType = City
// if bodiesType generated from bodies.ApplicationJson, we get that value from prefix and suffix
//		suffix = [ReqBody | RespBody] and prefix should be uri + method name.
//		example prefix could be UsersUserIdDelete
func assignBodyName(bodies raml.Bodies, prefix, suffix string) string {
	var bodiesType string

	if len(bodies.Type) > 0 {
		bodiesType = convertToGoType(bodies.Type)
	} else if bodies.ApplicationJson != nil {
		if bodies.ApplicationJson.Type != "" {
			bodiesType = convertToGoType(bodies.ApplicationJson.Type)
		} else {
			bodiesType = prefix + suffix
		}
	}

	return bodiesType
}
