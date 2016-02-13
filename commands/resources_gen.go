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
	IsServer bool
	NeedJSON bool
}

// create a resource definition
func newResourceDef(endpoint string) resourceDef {
	rd := resourceDef{
		Endpoint: endpoint,
	}
	rd.Name = strings.Title(normalizeURI(endpoint))
	return rd
}

// method of resource's interface
type interfaceMethod struct {
	*raml.Method
	MethodName   string
	Endpoint     string
	Verb         string
	ReqBody      string         // request body type
	RespBody     string         // response body type
	ResourcePath string         // normalized resource path
	Resource     *raml.Resource // resource object of this method
	MethodParam  string
}

// create an interfaceMethod object
func newInterfaceMethod(r *raml.Resource, rd *resourceDef, m *raml.Method, methodName, parentEndpoint, curEndpoint string) interfaceMethod {
	im := interfaceMethod{
		Method:   m,
		Endpoint: parentEndpoint + curEndpoint,
		Verb:     strings.ToUpper(methodName),
		Resource: r,
	}

	if m.Bodies.Type != "" {
		im.ReqBody = m.Bodies.Type
		rd.NeedJSON = true
	}

	//set response body
	for k, v := range m.Responses {
		if k >= 200 && k < 300 {
			im.RespBody = assignBodyName(v.Bodies, normalizeURITitle(parentEndpoint)+normalizeURITitle(curEndpoint)+methodName, "RespBody")
		}
	}

	return im
}

func newServerInterfaceMethod(r *raml.Resource, rd *resourceDef, m *raml.Method, methodName, parentEndpoint, curEndpoint string) interfaceMethod {
	im := newInterfaceMethod(r, rd, m, methodName, parentEndpoint, curEndpoint)

	name := normalizeURI(parentEndpoint) + normalizeURI(curEndpoint)
	im.MethodName = name[len(rd.Name):] + methodName

	return im
}

func newClientInterfaceMethod(r *raml.Resource, rd *resourceDef, m *raml.Method, methodName, parentEndpoint, curEndpoint string) interfaceMethod {
	im := newInterfaceMethod(r, rd, m, methodName, parentEndpoint, curEndpoint)

	postBuildParams := func(r *raml.Resource, bodyType string) string {
		paramsStr := strings.Join(getResourceParams(r), ",")
		if len(paramsStr) > 0 {
			paramsStr += " string"
		}

		//append body type
		if len(bodyType) > 0 {
			if len(paramsStr) > 0 {
				paramsStr += ", "
			}
			paramsStr += strings.ToLower(bodyType) + " " + bodyType
		}

		//append header
		if len(paramsStr) > 0 {
			paramsStr += ","
		}
		paramsStr += "headers,queryParams map[string]interface{}"

		return paramsStr
	}

	im.ResourcePath = templatingResourcePath(parentEndpoint + curEndpoint)

	name := normalizeURITitle(parentEndpoint + curEndpoint)
	im.MethodName = strings.Title(name + methodName)

	im.ReqBody = assignBodyName(m.Bodies, name+methodName, "ReqBody")
	im.MethodParam = postBuildParams(r, im.ReqBody)

	return im
}

//assignBodyName assign bodies by bodies.Type or bodies.ApplicationJson
//if bodiesType generated from bodies.Type we dont need append prefix and suffix
//example : bodies.Type = City, so bodiesType = City
//if bodiesType generated from bodies.ApplicationJson, we get that value from prefix and suffix
//suffix = [ReqBody | RespBody] and prefix should be uri + method name.
//example prefix could be UsersUserIdDelete
func assignBodyName(bodies raml.Bodies, prefix, suffix string) string {
	var bodiesType string

	if len(bodies.Type) > 0 {
		bodiesType = bodies.Type
	} else if bodies.ApplicationJson != nil {
		bodiesType = prefix + suffix
	}

	return bodiesType
}

// add a method to resource definition
func (rd *resourceDef) addMethod(r *raml.Resource, m *raml.Method, methodName, parentEndpoint, curEndpoint string) {
	if m == nil {
		return
	}
	var im interfaceMethod
	if rd.IsServer {
		im = newServerInterfaceMethod(r, rd, m, methodName, parentEndpoint, curEndpoint)
	} else {
		im = newClientInterfaceMethod(r, rd, m, methodName, parentEndpoint, curEndpoint)
	}
	rd.Methods = append(rd.Methods, im)
}

// generate all methods of a resource recursively
func (rd *resourceDef) generateMethods(r *raml.Resource, parentEndpoint, curEndpoint string) {
	rd.addMethod(r, r.Get, "Get", parentEndpoint, curEndpoint)
	rd.addMethod(r, r.Post, "Post", parentEndpoint, curEndpoint)
	rd.addMethod(r, r.Put, "Put", parentEndpoint, curEndpoint)
	rd.addMethod(r, r.Patch, "Patch", parentEndpoint, curEndpoint)
	rd.addMethod(r, r.Delete, "Delete", parentEndpoint, curEndpoint)

	for k, v := range r.Nested {
		rd.generateMethods(v, parentEndpoint+curEndpoint, k)
	}
}
