package resource

import (
	"regexp"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
)

var (
	reResource = regexp.MustCompile("({{1}[\\w\\s]+}{1})")
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

type ServerMethodConstructor func(*raml.APIDefinition, *raml.Resource, *Resource, *raml.Method, string) MethodInterface
type ClientMethodConstructor func(*raml.Resource, *Resource, *raml.Method, string) (MethodInterface, error)

// add a method to resource definition
func (rd *Resource) addMethod(r *raml.Resource, m *raml.Method, methodName string,
	smc ServerMethodConstructor, cmc ClientMethodConstructor) {
	var im MethodInterface
	var err error

	if m == nil {
		return
	}

	if rd.IsServer {
		im = smc(rd.APIDef, r, rd, m, methodName)
	} else {
		im, err = cmc(r, rd, m, methodName)
		if err != nil {
			log.Errorf("client interface method error, err = %v", err)
			return
		}
	}
	rd.Methods = append(rd.Methods, im)
}

// GenerateMethods generates all methods of a resource recursively
func (rd *Resource) GenerateMethods(r *raml.Resource, lang string, smc ServerMethodConstructor, cmc ClientMethodConstructor) {
	rd.addMethod(r, r.Get, "Get", smc, cmc)
	rd.addMethod(r, r.Post, "Post", smc, cmc)
	rd.addMethod(r, r.Put, "Put", smc, cmc)
	rd.addMethod(r, r.Patch, "Patch", smc, cmc)
	rd.addMethod(r, r.Delete, "Delete", smc, cmc)
	rd.addMethod(r, r.Options, "Options", smc, cmc)

	for _, v := range r.Nested {
		rd.GenerateMethods(v, lang, smc, cmc)
	}
}

// _getResourceParams is the recursive function of getResourceParams
func _getResourceParams(r *raml.Resource, params []string) []string {
	if r == nil {
		return params
	}

	matches := reResource.FindAllString(r.URI, -1)
	for _, v := range matches {
		params = append(params, v[1:len(v)-1])
	}

	return _getResourceParams(r.Parent, params)
}

// GetResourceParams get all params of a resource
// examples:
// /users  							  : no params
// /users/{userId}					  : params 1 = userId
// /users/{userId}/address/{addressId : params 1= userId, param 2= addressId
func GetResourceParams(r *raml.Resource) []string {
	params := []string{}
	return _getResourceParams(r, params)
}
