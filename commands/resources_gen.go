package commands

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

// resourceDef is Go code representation of a resource
type resourceDef struct {
	Name     string
	Endpoint string
	Methods  []interfaceMethod
}

// create a resource definition
func newResourceDef(endpoint string) resourceDef {
	rd := resourceDef{Endpoint: endpoint}
	rd.Name = strings.Title(normalizeURI(endpoint))
	return rd
}

// method of resource's interface
type interfaceMethod struct {
	*raml.Method
	MethodName string
	Endpoint   string
	Verb       string
	ReqBody    string // request body type
}

// create an interfaceMethod object
func newInterfaceMethod(rd *resourceDef, m *raml.Method, methodName, parentEndpoint, curEndpoint string) interfaceMethod {
	im := interfaceMethod{
		Method:   m,
		Endpoint: parentEndpoint + curEndpoint,
		Verb:     strings.ToUpper(methodName),
	}
	name := normalizeURI(parentEndpoint) + normalizeURI(curEndpoint)
	im.MethodName = name[len(rd.Name):] + methodName
	if m.Bodies.Type != "" {
		im.ReqBody = m.Bodies.Type
	}
	return im
}

// add a method to resource definition
func (rd *resourceDef) addMethod(m *raml.Method, methodName, parentEndpoint, curEndpoint string) {
	if m == nil {
		return
	}
	im := newInterfaceMethod(rd, m, methodName, parentEndpoint, curEndpoint)
	rd.Methods = append(rd.Methods, im)
}

// generate all methods of a resource recursively
func (rd *resourceDef) generateMethods(r *raml.Resource, parentEndpoint, curEndpoint string) {
	rd.addMethod(r.Get, "Get", parentEndpoint, curEndpoint)
	rd.addMethod(r.Post, "Post", parentEndpoint, curEndpoint)
	rd.addMethod(r.Put, "Put", parentEndpoint, curEndpoint)
	rd.addMethod(r.Patch, "Patch", parentEndpoint, curEndpoint)
	rd.addMethod(r.Delete, "Delete", parentEndpoint, curEndpoint)
	for k, v := range r.Nested {
		rd.generateMethods(v, parentEndpoint+curEndpoint, k)
	}
}
