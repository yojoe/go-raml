package resource

import (
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
)

const (
	maxCommentPerLine = 80
)

type ResourceInterface interface{}

// Resource is Go code representation of a resource
type Resource struct {
	APIDef      *raml.APIDefinition
	Name        string            // resource name
	Endpoint    string            // root endpoint
	Methods     []MethodInterface // all methods of this resource
	IsServer    bool              // true if it is resource definition for server
	PackageName string            // Name of the package this resource resides in
}

// New creates a resource definition
func New(apiDef *raml.APIDefinition, endpoint, packageName string) Resource {
	return Resource{
		Endpoint:    endpoint,
		APIDef:      apiDef,
		Name:        strings.Title(commons.NormalizeURI(endpoint)),
		PackageName: packageName,
	}
}

type ServerMethodConstructor func(*raml.APIDefinition, *raml.Resource, *Resource, *raml.Method, string, string) MethodInterface
type ClientMethodConstructor func(*raml.Resource, *Resource, *raml.Method, string, string) (MethodInterface, error)

// add a method to resource definition
func (rd *Resource) addMethod(r *raml.Resource, m *raml.Method, methodName, lang string,
	smc ServerMethodConstructor, cmc ClientMethodConstructor) {
	var im MethodInterface
	var err error

	if m == nil {
		return
	}

	if rd.IsServer {
		im = smc(rd.APIDef, r, rd, m, methodName, lang)
	} else {
		im, err = cmc(r, rd, m, methodName, lang)
		if err != nil {
			log.Errorf("client interface method error, err = %v", err)
			return
		}
	}
	rd.Methods = append(rd.Methods, im)
}

// GenerateMethods generates all methods of a resource recursively
func (rd *Resource) GenerateMethods(r *raml.Resource, lang string, smc ServerMethodConstructor, cmc ClientMethodConstructor) {
	rd.addMethod(r, r.Get, "Get", lang, smc, cmc)
	rd.addMethod(r, r.Post, "Post", lang, smc, cmc)
	rd.addMethod(r, r.Put, "Put", lang, smc, cmc)
	rd.addMethod(r, r.Patch, "Patch", lang, smc, cmc)
	rd.addMethod(r, r.Delete, "Delete", lang, smc, cmc)
	rd.addMethod(r, r.Options, "Options", lang, smc, cmc)

	for _, v := range r.Nested {
		rd.GenerateMethods(v, lang, smc, cmc)
	}
}
