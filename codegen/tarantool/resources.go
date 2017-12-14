package tarantool

import (
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
	"github.com/pinzolo/casee"
)

type response struct {
	Code     int
	BodyType string
}

// Method is a tarantool code representation of a method
type Method struct {
	RamlMethod *raml.Method
	Verb       string
	EndPoint   string
	ReqBody    string
	Responses  []response
}

// Resource is tarantool code representation of a resource
type Resource struct {
	RamlResource *raml.Resource
	Methods      []*Method
}

type ResourceByURI []*Resource

type TarantoolResources struct {
	Resources ResourceByURI
}

func (b ResourceByURI) Len() int      { return len(b) }
func (b ResourceByURI) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ResourceByURI) Less(i, j int) bool {
	return b[i].URI() < b[j].URI()
}

// Handler returns the name of the function that handles requests for this method
func (m *Method) Handler() string {

	if len(m.RamlMethod.DisplayName) > 0 {
		return casee.ToSnakeCase(commons.DisplayNameToFuncName(m.RamlMethod.DisplayName))
	}
	name := commons.ReplaceNonAlphanumerics(commons.NormalizeURI(m.EndPoint))
	return casee.ToSnakeCase(name + m.Verb)
}

func formatUri(uri string) string {
	return strings.Replace(strings.Replace(uri, "{", ":", -1), "}", "", -1)
}

// URI returns the tarantool URI of the fullURI
func (m *Method) URI() string {
	return formatUri(m.EndPoint)

}

// URI returns the tarantool URI of the fullURI
func (r *Resource) URI() string {
	fullURI := r.RamlResource.FullURI()
	return formatUri(fullURI)

}

// Handler returns the name of the function handling requests to this resource
func (r *Resource) Handler() string {
	if len(r.RamlResource.DisplayName) > 0 {
		return commons.DisplayNameToFuncName(r.RamlResource.DisplayName) + "_handler"
	}
	name := commons.ReplaceNonAlphanumerics(commons.NormalizeURI(r.RamlResource.FullURI()))
	return name + "_handler"
}

// addMethod adds a method to resource definition
func (r *Resource) addMethod(m *raml.Method, methodName string) {
	if m == nil {
		return
	}
	endPoint := r.RamlResource.FullURI()
	verbTitle := strings.Title(strings.ToLower(methodName))
	normalizedEndpoint := commons.NormalizeURITitle(endPoint)
	reqBody := setBodyName(
		m.Bodies, normalizedEndpoint+verbTitle, commons.ReqBodySuffix)

	responses := make([]response, 0)
	for code, methodResponse := range m.Responses {
		bodyType := setBodyName(methodResponse.Bodies, endPoint+verbTitle, commons.RespBodySuffix)
		if bodyType == "" {
			continue
		}
		resp := response{
			Code:     commons.AtoiOrPanic(string(code)),
			BodyType: bodyType,
		}
		responses = append(responses, resp)

	}
	method := &Method{
		RamlMethod: m,
		Verb:       methodName,
		EndPoint:   endPoint,
		ReqBody:    reqBody,
		Responses:  responses,
	}
	r.Methods = append(r.Methods, method)
}

// generateMethods generates all methods of a resource
func (r *Resource) generateMethods() {
	r.addMethod(r.RamlResource.Get, "GET")
	r.addMethod(r.RamlResource.Post, "POST")
	r.addMethod(r.RamlResource.Put, "PUT")
	r.addMethod(r.RamlResource.Patch, "PATCH")
	r.addMethod(r.RamlResource.Delete, "DELETE")
	r.addMethod(r.RamlResource.Options, "OPTIONS")
}

func flattenResources(resources map[string]raml.Resource) TarantoolResources {
	tarantoolResources := TarantoolResources{}
	for _, resource := range resources {
		r := resource
		tResource := Resource{RamlResource: &r}
		tResource.generateMethods()
		if len(tResource.Methods) > 0 {
			tarantoolResources.Resources = append(tarantoolResources.Resources, &tResource)
		}
		tarantoolResources.AddNested(&r)
	}
	return tarantoolResources
}

func (tr *TarantoolResources) AddNested(resource *raml.Resource) {
	for _, nestedResource := range resource.Nested {
		tResource := Resource{RamlResource: nestedResource}
		tResource.generateMethods()
		if len(tResource.Methods) > 0 {
			tr.Resources = append(tr.Resources, &tResource)
		}
		tr.AddNested(nestedResource)
	}
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
	var bodyName string
	prefix = commons.NormalizeURITitle(prefix)

	if len(bodies.Type) > 0 && bodies.Type != "object" {
		bodyName = bodies.Type
	} else if bodies.ApplicationJSON != nil {
		if bodies.ApplicationJSON.TypeString() != "" && bodies.ApplicationJSON.TypeString() != "object" {
			bodyName = bodies.ApplicationJSON.TypeString()
		} else {
			bodyName = prefix + suffix
		}
	}

	if commons.IsJSONString(bodyName) {
		bodyName = prefix + suffix
	}

	return bodyName

}
