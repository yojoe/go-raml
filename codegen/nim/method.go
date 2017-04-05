package nim

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen/commons"
	cr "github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/codegen/security"
	"github.com/Jumpscale/go-raml/raml"
)

type method struct {
	*cr.Method
}

func getMethodName(endpoint, displayName, verb string) string {
	if len(displayName) > 0 {
		return strings.Replace(displayName, " ", "", -1)
	} else {
		return commons.NormalizeURI(formatProcName(endpoint)) + strings.Title(strings.ToLower(verb))
	}
}

// creates new Nim method
func newMethod(apiDef *raml.APIDefinition, r *raml.Resource, rd *cr.Resource, m *raml.Method,
	methodName string) (cr.MethodInterface, error) {

	rm := cr.NewMethod(r, rd, m, methodName, setBodyName)

	rm.MethodName = getMethodName(r.FullURI(), rm.DisplayName, methodName)

	rm.ResourcePath = commons.ParamizingURI(rm.Endpoint, "&")
	if apiDef != nil {
		rm.SecuredBy = security.GetMethodSecuredBy(apiDef, r, m)
	}
	return method{Method: &rm}, nil
}

// creates new client method
func newClientMethod(r *raml.Resource, rd *cr.Resource, m *raml.Method,
	methodName string) (cr.MethodInterface, error) {
	return newMethod(nil, r, rd, m, methodName)
}

// creates new server method
func newServerMethod(apiDef *raml.APIDefinition, r *raml.Resource, rd *cr.Resource, m *raml.Method,
	methodName string) cr.MethodInterface {

	// creates generic method
	mi, err := newMethod(apiDef, r, rd, m, methodName)
	if err != nil {
		log.Fatalf("newServerMethod unexpected error:%v", err)
	}

	return mi
}

// JesterEndpoint returns endpoint in jester format
func (m method) JesterEndpoint() string {
	e := strings.Replace(m.Endpoint, "{", "@", -1)
	return strings.Replace(e, "}", "", -1)
}

// ProcParams are params of this jester endpoint handler
func (m method) ServerProcParams() string {
	var params []string

	for _, p := range cr.GetResourceParams(m.Resource()) {
		params = append(params, p+": string")
	}

	params = append(params, "req: Request")
	return strings.Join(params, ", ")
}

// CallParams are params when we call jester handler
func (m method) ServerCallParams() string {
	var params []string

	for _, p := range cr.GetResourceParams(m.Resource()) {
		params = append(params, fmt.Sprintf(`@"%v"`, p))
	}

	params = append(params, "request")
	return strings.Join(params, ", ")
}

func (m method) ClientProcParams() string {
	params := []string{}

	// request body
	if m.ReqBody != "" {
		params = append(params, fmt.Sprintf("reqBody: %v", m.ReqBody))
	}

	// resource params
	for _, p := range cr.GetResourceParams(m.Resource()) {
		params = append(params, fmt.Sprintf("%v: string", p))
	}

	// query params
	params = append(params, `queryParams: Table[string, string] = initTable[string, string]()`)

	str := strings.Join(params, ", ")
	if str != "" {
		str = ", " + str
	}
	return str
}

// ClientCallParams are params when we call jester handler
func (m method) ClientCallParams() string {
	params := []string{
		m.ResourcePath,
		fmt.Sprintf(`"%v"`, m.Verb()),
	}
	if m.ReqBody != "" {
		params = append(params, "$$reqBody")
	}
	params = append(params, "queryParams=queryParams")

	return strings.Join(params, ", ")
}

func (m method) ContentRetval() string {
	retval := m.RespBody
	if retval == "" {
		retval = "string"
	}
	return retval
}

func (m method) Secured() bool {
	return len(m.SecuredBy) > 0
}

// SecurityScopes retuns security scopes of a method as single string
func (m method) SecurityScopes() string {
	if len(m.SecuredBy) == 0 {
		return ""
	}
	s, err := security.GetQuotedScopes(m.SecuredBy[0])
	if err != nil {
		log.Fatalf("failed to get security scopes of method `%v %v` err = %v", m.Verb(), m.ResourcePath, err)
	}
	return strings.Join(s, ", ")
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
		if bodies.ApplicationJSON.TypeString() != "" && bodies.ApplicationJSON.TypeString() != "object" {
			tipe = toNimType(bodies.ApplicationJSON.TypeString())
		} else {
			tipe = prefix + suffix
		}
	}

	if commons.IsJSONString(tipe) {
		tipe = prefix + suffix
	}

	return tipe
}

// format Nim proc name from complete URI
func formatProcName(fullURI string) string {
	// remove leading `/`
	fullURI = fullURI[1:]

	// When meet `/{`
	// - replace it with `By`
	// - make uppercase the first char after `/{`
	spl := strings.Split(fullURI, "/{")
	tmp := []string{}
	for i, v := range spl {
		if i != 0 {
			v = strings.Title(v)
		}
		tmp = append(tmp, v)
	}
	name := strings.Join(tmp, "By")

	// when meet `/`
	// - make uppercase the first char after `/`
	// - remove the `/`
	spl = strings.Split(name, "/")
	tmp = []string{}
	for i, v := range spl {
		if i != 0 {
			v = strings.Title(v)
		}
		tmp = append(tmp, v)
	}

	return strings.Join(tmp, "")
}
