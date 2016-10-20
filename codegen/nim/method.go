package nim

import (
	"fmt"
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

// JesterEndpoint returns endpoint in jester format
func (m method) JesterEndpoint() string {
	e := strings.Replace(m.Endpoint, "{", "@", -1)
	return strings.Replace(e, "}", "", -1)
}

// ProcParams are params of this jester endpoint handler
func (m method) ProcParams() string {
	var params []string

	for _, p := range cr.GetResourceParams(m.Resource()) {
		params = append(params, p+": string")
	}

	params = append(params, "req: PRequest")
	return strings.Join(params, ", ")
}

// CallParams are params when we call jester handler
func (m method) CallParams() string {
	var params []string

	for _, p := range cr.GetResourceParams(m.Resource()) {
		params = append(params, fmt.Sprintf(`@"%v"`, p))
	}

	params = append(params, "request")
	return strings.Join(params, ", ")
}

func (m method) ContentRetval() string {
	retval := m.RespBody
	if retval == "" {
		retval = "string"
	}
	return retval
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

	prefix = commons.NormalizeURI(prefix)

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
