package golang

import (
	"fmt"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

type clientMethod struct {
	*resource.Method
}

type respBody struct {
	Code int
	Type string
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
	gcm.MethodName = commons.ReplaceNonAlphanumerics(strings.Title(gcm.MethodName))

	// method param
	gcm.Params = buildParams(gcm.RAMLResource, gcm.ReqBody)
}

// return true if this method need to import encoding/json
func (gcm clientMethod) needImportEncodingJSON() bool {
	return gcm.RespBody != ""
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

// ReturnTypes returns all types returned by this method
func (gcm clientMethod) ReturnTypes() string {
	var types []string
	for _, resp := range gcm.SuccessRespBodyTypes() {
		types = append(types, resp.Type)
	}
	types = append(types, []string{"*http.Response", "error"}...)

	return fmt.Sprintf("(%v)", strings.Join(types, ","))
}

func (gcm clientMethod) HasRespBody() bool {
	return len(gcm.RespBodyTypes()) > 0
}

// RespBodyTypes returns all possible type of response body
func (gcm clientMethod) RespBodyTypes() (resps []respBody) {
	for code, resp := range gcm.Responses {
		resp := respBody{
			Code: commons.AtoiOrPanic(string(code)),
			Type: setBodyName(resp.Bodies, gcm.Endpoint+gcm.VerbTitle(), commons.RespBodySuffix),
		}
		if resp.Type != "" {
			resps = append(resps, resp)
		}
	}
	return
}

// FailedRespBodyTypes return all response body that considered a failed response
// i.e. non 2xx status code
func (gcm clientMethod) FailedRespBodyTypes() (resps []respBody) {
	for _, resp := range gcm.RespBodyTypes() {
		if resp.Code < 200 || resp.Code >= 300 {
			resps = append(resps, resp)
		}
	}
	return
}

// SuccessRespBodyTypes returns all response body that considered as success
// i.e. 2xx status code
func (gcm clientMethod) SuccessRespBodyTypes() (resps []respBody) {
	for _, resp := range gcm.RespBodyTypes() {
		if resp.Code >= 200 && resp.Code < 300 {
			resps = append(resps, resp)
		}
	}
	return
}

// ReqBodyType returns type of the request body
func (gcm clientMethod) ReqBodyType() string {
	return setBodyName(gcm.Bodies, gcm.Endpoint+gcm.VerbTitle(), commons.ReqBodySuffix)
}

func (gcm clientMethod) needImportGoraml() bool {
	return gcm.HasRespBody()
}
