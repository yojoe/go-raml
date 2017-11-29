package python

import (
	"github.com/chuckpreslar/inflect"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/types"
)

// generate type name for
// inline type declared in request/response body
func genTypeName(tip types.TypeInBody) string {
	methodName := setServerMethodName(tip.Endpoint.Method.DisplayName, tip.Endpoint.Verb, tip.Endpoint.Resource)
	suffix := commons.ReqBodySuffix
	if tip.ReqResp == types.HTTPResponse {
		suffix = commons.RespBodySuffix
	}
	return inflect.UpperCamelCase(methodName + suffix)
}
