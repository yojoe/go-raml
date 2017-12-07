package tarantool

import (
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

// Method is a tarantool code representation of a method
type Method struct {
	RamlMethod *raml.Method
	Verb       string
	EndPoint   string
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
		return commons.DisplayNameToFuncName(m.RamlMethod.DisplayName)
	}
	name := commons.ReplaceNonAlphanumerics(commons.NormalizeURI(m.EndPoint))
	return name + m.Verb
}

// URI returns the tarantool URI of the fullURI
func (r *Resource) URI() string {
	fullURI := r.RamlResource.FullURI()
	return strings.Replace(strings.Replace(fullURI, "{", ":", -1), "}", "", -1)

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
	r.Methods = append(r.Methods, &Method{RamlMethod: m, Verb: methodName, EndPoint: r.RamlResource.FullURI()})
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
		tarantoolResources.Resources = append(tarantoolResources.Resources, &tResource)
		tarantoolResources.AddNested(&r)
	}
	return tarantoolResources
}

func (tr *TarantoolResources) AddNested(resource *raml.Resource) {
	for _, nestedResource := range resource.Nested {
		tResource := Resource{RamlResource: nestedResource}
		tResource.generateMethods()
		tr.Resources = append(tr.Resources, &tResource)
		tr.AddNested(nestedResource)
	}
}
