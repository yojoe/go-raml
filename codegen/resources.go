package codegen

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
)

const (
	maxCommentPerLine = 80
)

type resourceInterface interface{}

// resourceDef is Go code representation of a resource
type resourceDef struct {
	APIDef      *raml.APIDefinition
	Name        string            // resource name
	Endpoint    string            // root endpoint
	Methods     []methodInterface // all methods of this resource
	IsServer    bool              // true if it is resource definition for server
	PackageName string            // Name of the package this resource resides in

}

// create a resource definition
func newResourceDef(apiDef *raml.APIDefinition, endpoint, packageName string) resourceDef {
	rd := resourceDef{
		Endpoint: endpoint,
		APIDef:   apiDef,
	}
	rd.Name = strings.Title(normalizeURI(endpoint))
	rd.PackageName = packageName
	return rd
}

// add a method to resource definition
func (rd *resourceDef) addMethod(r *raml.Resource, m *raml.Method, methodName, parentEndpoint, curEndpoint, lang string) {
	var im methodInterface
	var err error

	if m == nil {
		return
	}

	if rd.IsServer {
		im = newServerMethod(rd.APIDef, r, rd, m, methodName, parentEndpoint, curEndpoint, lang)
	} else {
		im, err = newClientMethod(r, rd, m, methodName, parentEndpoint, curEndpoint, lang)
		if err != nil {
			log.Errorf("client interface method error, err = %v", err)
			return
		}
	}
	rd.Methods = append(rd.Methods, im)
}

// generate all methods of a resource recursively
func (rd *resourceDef) generateMethods(r *raml.Resource, parentEndpoint, curEndpoint, lang string) {
	rd.addMethod(r, r.Get, "Get", parentEndpoint, curEndpoint, lang)
	rd.addMethod(r, r.Post, "Post", parentEndpoint, curEndpoint, lang)
	rd.addMethod(r, r.Put, "Put", parentEndpoint, curEndpoint, lang)
	rd.addMethod(r, r.Patch, "Patch", parentEndpoint, curEndpoint, lang)
	rd.addMethod(r, r.Delete, "Delete", parentEndpoint, curEndpoint, lang)

	for k, v := range r.Nested {
		rd.generateMethods(v, parentEndpoint+curEndpoint, k, lang)
	}
}
