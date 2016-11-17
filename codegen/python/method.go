package python

import (
	"fmt"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/codegen/security"
	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
)

// python server method
type serverMethod struct {
	*resource.Method
	MiddlewaresArr []middleware
}

// setup sets all needed variables
func (sm *serverMethod) setup(apiDef *raml.APIDefinition, r *raml.Resource, rd *resource.Resource) error {
	// method name
	if len(sm.DisplayName) > 0 {
		sm.MethodName = commons.DisplayNameToFuncName(sm.DisplayName)
	} else {
		sm.MethodName = snakeCaseResourceURI(r) + "_" + strings.ToLower(sm.Verb())
	}
	sm.Params = strings.Join(resource.GetResourceParams(r), ", ")
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

// defines a python client lib method
type clientMethod struct {
	resource.Method
	PRArgs string // python requests's args
	PRCall string // the way we call python request
}

func newClientMethod(r *raml.Resource, rd *resource.Resource, m *raml.Method, methodName string) (resource.MethodInterface, error) {
	method := resource.NewMethod(r, rd, m, methodName, setBodyName)

	method.ResourcePath = commons.ParamizingURI(method.Endpoint, "+")

	name := commons.NormalizeURITitle(method.Endpoint)

	method.ReqBody = setBodyName(m.Bodies, name+methodName, "ReqBody")

	pcm := clientMethod{Method: method}
	pcm.setup()
	return pcm, nil
}

func (pcm *clientMethod) setup() {
	prArgs := []string{"uri"}  // prArgs are arguments we supply to python request
	params := []string{"self"} // params are method signature params

	// for method with request body, we add `data` argument
	if pcm.Verb() == "PUT" || pcm.Verb() == "POST" || pcm.Verb() == "PATCH" {
		params = append(params, "data")
		prArgs = append(prArgs, "data")
	}

	// construct prArgs string from the array
	prArgs = append(prArgs, "headers=headers", "params=query_params")
	pcm.PRArgs = strings.Join(prArgs, ", ")

	// construct method signature
	params = append(params, resource.GetResourceParams(pcm.Resource())...)
	params = append(params, "headers=None", "query_params=None")
	pcm.Params = strings.Join(params, ", ")

	// python request call
	// we encapsulate the call to put, post, and patch.
	// To be able to accept plain string or dict.
	// if it is a dict, we encode it to json
	if pcm.Verb() == "PUT" || pcm.Verb() == "POST" || pcm.Verb() == "PATCH" {
		pcm.PRCall = fmt.Sprintf("self.client.%v", strings.ToLower(pcm.Verb()))
	} else {
		pcm.PRCall = fmt.Sprintf("self.client.session.%v", strings.ToLower(pcm.Verb()))
	}

	if len(pcm.DisplayName) > 0 {
		pcm.MethodName = commons.DisplayNameToFuncName(pcm.DisplayName)
	} else {
		pcm.MethodName = snakeCaseResourceURI(pcm.Resource()) + "_" + strings.ToLower(pcm.Verb())
	}
}

// create server resource's method
func newServerMethod(apiDef *raml.APIDefinition, r *raml.Resource, rd *resource.Resource, m *raml.Method,
	methodName string) resource.MethodInterface {

	method := resource.NewMethod(r, rd, m, methodName, setBodyName)
	method.SecuredBy = security.GetMethodSecuredBy(apiDef, r, m)

	pm := serverMethod{
		Method: &method,
	}
	pm.setup(apiDef, r, rd)
	return pm
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
		if bodies.ApplicationJSON.Type != "" && bodies.ApplicationJSON.Type != "object" {
			tipe = bodies.ApplicationJSON.Type
		} else {
			tipe = prefix + suffix
		}
	}

	if commons.IsJSONString(tipe) {
		tipe = prefix + suffix
	}

	return tipe

}
