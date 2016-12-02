package resource

import (
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

type MethodInterface interface {
	Verb() string
	Resource() *raml.Resource
	EndpointStr() string
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
	RAMLResource *raml.Resource // resource object of this method
	Params       string         // methods params
	FuncComments []string
	SecuredBy    []raml.DefinitionChoice
}

func (m Method) Verb() string {
	return m.verb
}

func (m Method) Resource() *raml.Resource {
	return m.RAMLResource
}

func (m Method) EndpointStr() string {
	return m.Endpoint
}

type SetBodyName func(raml.Bodies, string, string) string

func NewMethod(r *raml.Resource, rd *Resource, m *raml.Method, methodName string, sbn SetBodyName) Method {
	method := Method{
		Method:       m,
		Endpoint:     r.FullURI(),
		verb:         strings.ToUpper(methodName),
		RAMLResource: r,
	}

	// set request body
	method.ReqBody = sbn(m.Bodies, method.Endpoint+methodName, commons.ReqBodySuffix)

	//set response body
	for k, v := range m.Responses {
		code := commons.AtoiOrPanic(string(k))
		if code >= 200 && code < 300 {
			method.RespBody = sbn(v.Bodies, method.Endpoint+methodName, commons.RespBodySuffix)
		}
	}

	// set func comment
	if len(m.Description) > 0 {
		method.FuncComments = commons.ParseDescription(m.Description)
	}

	return method
}

type ByEndpoint []MethodInterface

func (b ByEndpoint) Len() int      { return len(b) }
func (b ByEndpoint) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByEndpoint) Less(i, j int) bool {
	return strings.Compare(b[i].EndpointStr(), b[j].EndpointStr()) < 0
}
