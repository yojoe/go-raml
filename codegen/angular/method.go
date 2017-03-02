package angular

import (
	"fmt"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/codegen/security"
	"github.com/Jumpscale/go-raml/raml"
)

// python server method
type serverMethod struct {
	*resource.Method
}

// setup sets all needed variables
func (sm *serverMethod) setup(apiDef *raml.APIDefinition, r *raml.Resource, rd *resource.Resource, resourceParams []string) error {
	// method name
	if len(sm.DisplayName) > 0 {
		sm.MethodName = commons.DisplayNameToFuncName(sm.DisplayName)
	} else {
		sm.MethodName = snakeCaseResourceURI(r) + "_" + strings.ToLower(sm.Verb())
	}
	sm.Params = strings.Join(resourceParams, ", ")
	sm.Endpoint = strings.Replace(sm.Endpoint, "{", "<", -1)
	sm.Endpoint = strings.Replace(sm.Endpoint, "}", ">", -1)

	return nil
}

// create server resource's method
func newServerMethod(apiDef *raml.APIDefinition, r *raml.Resource, rd *resource.Resource, m *raml.Method,
	methodName string) resource.MethodInterface {

	method := resource.NewMethod(r, rd, m, methodName, setBodyName)
	method.SecuredBy = security.GetMethodSecuredBy(apiDef, r, m)

	pm := serverMethod{
		Method: &method,
	}
	params := resource.GetResourceParams(r)
	pm.setup(apiDef, r, rd, params)
	return pm
}

// defines a python client lib method
type clientMethod struct {
	resource.Method
	PRArgs string // python requests's args
	PRCall string // the way we call python request
}

func newClientMethod(r *raml.Resource, rd *resource.Resource, m *raml.Method, methodName string) (resource.MethodInterface, error) {
	method := resource.NewMethod(r, rd, m, methodName, setBodyName)

	method.ResourcePath = strings.Replace(method.Endpoint, "{", "${", -1)

	name := commons.NormalizeURITitle(method.Endpoint)

	method.ReqBody = setBodyName(m.Bodies, name+methodName, "ReqBody")

	pcm := clientMethod{Method: method}
	pcm.setup()
	return pcm, nil
}

func (pcm *clientMethod) setup() {
	prArgs := []string{"uri"} // prArgs are arguments we supply to python request
	params := []string{}      // params are method signature params

	// for method with request body, we add `data` argument
	if pcm.Verb() == "PUT" || pcm.Verb() == "POST" || pcm.Verb() == "PATCH" {
		params = append(params, "data")
		prArgs = append(prArgs, "data")
	}

	// construct prArgs string from the array
	prArgs = append(prArgs, "options")
	pcm.PRArgs = strings.Join(prArgs, ", ")

	// construct method signature
	params = append(params, resource.GetResourceParams(pcm.Resource())...)
	params = append(params, "headers={}", "query_params={}")
	pcm.Params = strings.Join(params, ", ")

	// python request call
	pcm.PRCall = fmt.Sprintf("this.http.%v", strings.ToLower(pcm.Verb()))

	if len(pcm.DisplayName) > 0 {
		pcm.MethodName = commons.DisplayNameToFuncName(pcm.DisplayName)
	} else {
		pcm.MethodName = snakeCaseResourceURI(pcm.Resource()) + "_" + strings.ToLower(pcm.Verb())
	}
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
