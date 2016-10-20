package nim

import (
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	cr "github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

type method struct {
	*cr.Method
}

func newServerMethod(apiDef *raml.APIDefinition, r *raml.Resource, rd *cr.Resource, m *raml.Method,
	methodName, lang string) cr.MethodInterface {

	rm := cr.NewMethod(r, rd, m, methodName, setBodyName)

	// set method name
	if len(rm.DisplayName) > 0 {
		rm.MethodName = strings.Replace(rm.DisplayName, " ", "", -1)
	} else {
		rm.MethodName = commons.NormalizeURI(r.FullURI()) + methodName
	}
	return method{Method: &rm}
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

	if len(bodies.Type) > 0 && bodies.Type != "object" {
		tipe = toNimType(bodies.Type)
	} else if bodies.ApplicationJSON != nil {
		if bodies.ApplicationJSON.Type != "" && bodies.ApplicationJSON.Type != "object" {
			tipe = toNimType(bodies.ApplicationJSON.Type)
		} else {
			tipe = prefix + suffix
		}
	}

	if commons.IsJSONString(tipe) {
		tipe = prefix + suffix
	}

	return tipe
}
