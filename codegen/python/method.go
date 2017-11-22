package python

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/chuckpreslar/inflect"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/codegen/security"
	"github.com/Jumpscale/go-raml/raml"
)

type method struct {
	resource.Method
	ReqBody  string
	RespBody string // TODO : fix it properly as part of https://github.com/Jumpscale/go-raml/issues/350
}

func newMethod(rm resource.Method) *method {
	// resp body
	// TODO : fix it properly as part of https://github.com/Jumpscale/go-raml/issues/350
	var respBody string
	for code, resp := range rm.Responses {
		code := commons.AtoiOrPanic(string(code))
		if code >= 200 && code < 300 {
			respBody = setBodyName(resp.Bodies, rm.Endpoint+rm.VerbTitle(), commons.RespBodySuffix)
			break
		}
	}

	normalizedEndpoint := commons.NormalizeURITitle(rm.Endpoint)
	return &method{
		Method:   rm,
		ReqBody:  setBodyName(rm.Bodies, normalizedEndpoint+rm.VerbTitle(), commons.ReqBodySuffix),
		RespBody: respBody,
	}
}

// python server method
type serverMethod struct {
	*method
	MiddlewaresArr []middleware
	SecuredBy      []raml.DefinitionChoice
}

func setServerMethodName(displayName, verb string, resource *raml.Resource) string {
	if len(displayName) > 0 {
		return commons.DisplayNameToFuncName(displayName)
	}
	return snakeCaseResourceURI(resource) + "_" + strings.ToLower(verb)
}

func setReqBodyName(methodName string) string {
	return inflect.UpperCamelCase(methodName + "ReqBody")
}

// setup sets all needed variables
func (sm *serverMethod) setup(apiDef *raml.APIDefinition, r *raml.Resource, rd *resource.Resource, resourceParams []string) error {
	sm.MethodName = setServerMethodName(sm.DisplayName, sm.Verb(), r)

	if commons.HasJSONBody(&(sm.Bodies)) {
		sm.ReqBody = setReqBodyName(sm.MethodName)
	}

	sm.Params = strings.Join(resourceParams, ", ")
	sm.Endpoint = strings.Replace(sm.Endpoint, "{", "<", -1)
	sm.Endpoint = strings.Replace(sm.Endpoint, "}", ">", -1)

	// security middlewares
	for _, v := range sm.SecuredBy {
		if !security.ValidateScheme(v.Name, apiDef) {
			continue
		}
		// oauth2 middleware
		m, err := newPythonOauth2Middleware(v)
		if err != nil {
			log.Errorf("error creating middleware for method.err = %v", err)
			return err
		}
		sm.MiddlewaresArr = append(sm.MiddlewaresArr, m)
	}
	return nil
}

func newServerMethod(apiDef *raml.APIDefinition, rd *resource.Resource, rm resource.Method, kind string) serverMethod {
	meth := newMethod(rm)
	sm := serverMethod{
		method:    meth,
		SecuredBy: security.GetMethodSecuredBy(apiDef, rm.Resource, rm.Method),
	}
	params := resource.GetResourceParams(rm.Resource)
	if kind == serverKindSanic {
		params = append([]string{"request"}, params...)
	}
	sm.setup(apiDef, rm.Resource, rd, params)
	return sm

}

// defines a python client lib method
type clientMethod struct {
	*method
	ResourcePath string
	PRArgs       string // python requests's args
	PRCall       string // the way we call python request
}

func newClientMethod(rm resource.Method) clientMethod {
	meth := newMethod(rm)
	cm := clientMethod{
		method:       meth,
		ResourcePath: commons.ParamizingURI(rm.Endpoint, "+"),
	}
	cm.setup()
	return cm
}

func (pcm *clientMethod) setup() {
	prArgs := []string{"uri"}  // prArgs are arguments we supply to python request
	params := []string{"self"} // params are method signature params

	// for method with request body, we add `data` argument
	if !pcm.Bodies.IsEmpty() || pcm.Verb() == "PUT" || pcm.Verb() == "POST" || pcm.Verb() == "PATCH" {
		params = append(params, "data")
		prArgs = append(prArgs, "data")
	} else {
		prArgs = append(prArgs, "None")
	}

	// construct prArgs string from the array
	prArgs = append(prArgs, "headers", "query_params", "content_type")
	pcm.PRArgs = strings.Join(prArgs, ", ")

	// construct method signature
	params = append(params, resource.GetResourceParams(pcm.Resource)...)
	params = append(params, "headers=None", "query_params=None", `content_type="application/json"`)
	pcm.Params = strings.Join(params, ", ")

	// python request call
	// we encapsulate the call to put, post, and patch.
	// To be able to accept plain string or dict.
	// if it is a dict, we encode it to json
	if pcm.Verb() == "GET" || pcm.Verb() == "PUT" || pcm.Verb() == "POST" || pcm.Verb() == "PATCH" || pcm.Verb() == "DELETE" {
		pcm.PRCall = fmt.Sprintf("self.client.%v", strings.ToLower(pcm.Verb()))
	} else {
		pcm.PRCall = fmt.Sprintf("self.client.session.%v", strings.ToLower(pcm.Verb()))
	}

	if len(pcm.DisplayName) > 0 {
		pcm.MethodName = commons.DisplayNameToFuncName(pcm.DisplayName)
	} else {
		pcm.MethodName = snakeCaseResourceURI(pcm.Resource) + "_" + strings.ToLower(pcm.Verb())
	}
}

// IsArrayResponse returns true if the response body is
// of array type
func (pcm clientMethod) IsArrayResponse() bool {
	t := raml.Type{Type: pcm.RespBody}
	return t.IsArray() || t.IsBidimensiArray()
}

// RespBodyBasicType returns basic type of the response body.
// example : string[] -> string
func (pcm clientMethod) RespBodyBasicType() string {
	return commons.GetBasicType(pcm.RespBody)
}

func (pcm clientMethod) imports() []string {
	var imports []string

	// import from response body
	if pcm.HasRespBody() {
		basicType := pcm.RespBodyBasicType()
		imports = append(imports, "from ."+basicType+" import "+basicType)
	}
	return imports
}

// HasRespBody returns true if this method has response body
func (pcm clientMethod) HasRespBody() bool {
	return pcm.RespBody != ""
}

// create snake case function name from a resource URI
func snakeCaseResourceURI(r *raml.Resource) string {
	return _snakeCaseResourceURI(r, "")
}

func _snakeCaseResourceURI(r *raml.Resource, completeURI string) string {
	if r == nil {
		return completeURI
	}
	var snake string
	if len(r.URI) > 0 {
		uri := commons.NormalizeURI(r.URI)
		if len(uri) > 0 {
			if r.Parent != nil { // not root resource, need to add "_"
				snake = "_"
			}

			if strings.HasPrefix(r.URI, "/{") {
				snake += "by" + strings.ToUpper(uri[:1])
			} else {
				snake += strings.ToLower(uri[:1])
			}

			if len(uri) > 1 { // append with the rest of uri
				snake += uri[1:]
			}
		}
	}
	return _snakeCaseResourceURI(r.Parent, snake+completeURI)
}

// setBodyName set name of method's request/response body.
//
// Rules:
//  - use bodies.Type if not empty and not `object`
//  - use bodies.ApplicationJSON.Type if not empty and not `object`
//  - use prefix+suffix if:
//      - not meet previous rules
//      - previous rules produces JSON string
func setBodyName(bodies raml.Bodies, prefix, suffix string) string {
	var tipe string
	prefix = commons.NormalizeURITitle(prefix)

	if len(bodies.Type) > 0 && bodies.Type != "object" {
		tipe = bodies.Type
	} else if bodies.ApplicationJSON != nil {
		if bodies.ApplicationJSON.TypeString() != "" && bodies.ApplicationJSON.TypeString() != "object" {
			tipe = bodies.ApplicationJSON.TypeString()
		} else {
			tipe = prefix + suffix
		}
	}

	if commons.IsJSONString(tipe) {
		tipe = prefix + suffix
	}

	return tipe

}
