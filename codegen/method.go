package codegen

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
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

type goServerMethod struct {
	*Method
	Middlewares string
}

// setup go server method, initializes all needed variables
func (gm *goServerMethod) setup(apiDef *raml.APIDefinition, r *raml.Resource, rd *resourceDef, methodName string) error {
	// set method name
	name := normalizeURI(gm.Endpoint)
	if len(gm.DisplayName) > 0 {
		gm.MethodName = strings.Replace(gm.DisplayName, " ", "", -1)
	} else {
		gm.MethodName = name[len(rd.Name):] + methodName
	}

	/// if there is request body, we need to import validator
	if gm.ReqBody != "" {
		rd.NeedValidator = true
	}

	// if there is request/response body, then it needs to import encoding/json
	if gm.RespBody != "" || gm.ReqBody != "" {
		rd.NeedJSON = true
	}

	// setting middlewares
	middlewares := []string{}

	// security middlewares
	for _, v := range gm.SecuredBy {
		if !validateSecurityScheme(v.Name, apiDef) {
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

	if len(gm.Middlewares) > 0 {
		rd.WithMiddleware = true
	}
	return nil
}

type pythonServerMethod struct {
	*Method
	MiddlewaresArr []pythonMiddleware // TODO split interfaceMethod to pyton & Go version
}

// setup sets all needed variables
func (pm *pythonServerMethod) setup(apiDef *raml.APIDefinition, r *raml.Resource, rd *resourceDef) error {
	// method name
	if len(pm.DisplayName) > 0 {
		pm.MethodName = strings.Replace(pm.DisplayName, " ", "", -1)
	} else {
		pm.MethodName = snakeCaseResourceURI(r) + "_" + strings.ToLower(pm.Verb())
	}
	pm.Params = strings.Join(getResourceParams(r), ", ")
	pm.Endpoint = strings.Replace(pm.Endpoint, "{", "<", -1)
	pm.Endpoint = strings.Replace(pm.Endpoint, "}", ">", -1)

	// security middlewares
	for _, v := range pm.SecuredBy {
		if !validateSecurityScheme(v.Name, apiDef) {
			continue
		}
		// oauth2 middleware
		m, err := newPythonOauth2Middleware(v)
		if err != nil {
			log.Errorf("error creating middleware for method.err = %v", err)
			return err
		}
		pm.MiddlewaresArr = append(pm.MiddlewaresArr, m)
	}
	for _, v := range pm.MiddlewaresArr {
		rd.addPythonMiddleware(v)
	}
	return nil
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

	if lang == langGo {
		gm := goServerMethod{
			Method: &method,
		}
		gm.setup(apiDef, r, rd, methodName)
		return gm
	} else {
		pm := pythonServerMethod{
			Method: &method,
		}
		pm.setup(apiDef, r, rd)
		return pm
	}
}

func newClientMethod(r *raml.Resource, rd *resourceDef, m *raml.Method, methodName, parentEndpoint, curEndpoint string) (methodInterface, error) {
	im := newMethod(r, rd, m, methodName, parentEndpoint, curEndpoint)

	// build func/method params
	postBuildParams := func(r *raml.Resource, bodyType string) (string, error) {
		paramsStr := strings.Join(getResourceParams(r), ",")
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

	im.ResourcePath = paramizingURI(im.Endpoint)

	name := normalizeURITitle(im.Endpoint)

	if len(m.DisplayName) > 0 {
		im.MethodName = strings.Replace(m.DisplayName, " ", "", -1)
	} else {
		im.MethodName = strings.Title(name + methodName)
	}

	im.ReqBody = assignBodyName(m.Bodies, name+methodName, "ReqBody")

	methodParam, err := postBuildParams(r, im.ReqBody)
	if err != nil {
		return im, err
	}
	im.Params = methodParam

	return im, nil
}
