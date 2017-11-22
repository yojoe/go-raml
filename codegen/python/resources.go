package python

import (
	"sort"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

type pythonResource struct {
	*resource.Resource
	MiddlewaresArr []middleware
	Methods        []serverMethod
}

func (pr *pythonResource) addMiddleware(mwr middleware) {
	// check if already exist
	for _, v := range pr.MiddlewaresArr {
		if v.Name == mwr.Name {
			return
		}
	}
	pr.MiddlewaresArr = append(pr.MiddlewaresArr, mwr)
}

func newResource(rd resource.Resource, apiDef *raml.APIDefinition, kind string) pythonResource {
	r := pythonResource{
		Resource: &rd,
	}
	// generate methods
	for _, rm := range rd.Methods {
		sm := newServerMethod(apiDef, &rd, rm, kind)
		r.Methods = append(r.Methods, sm)
	}

	r.setMiddlewares()
	return r
}

// set middlewares to import
func (pr *pythonResource) setMiddlewares() {
	for _, pm := range pr.Methods {
		for _, m := range pm.MiddlewaresArr {
			pr.addMiddleware(m)
		}
	}
}

// generate flask representation of an RAML resource
// It has one file : an API route and implementation
func (pr *pythonResource) generate(fileName, tmplFile, tmplName, dir string) error {
	return commons.GenerateFile(pr, tmplFile, tmplName, fileName, false)
}

// return array of request body in this resource
func (pr pythonResource) ReqBodies() []string {
	var reqs []string
	for _, pm := range pr.Methods {
		if pm.ReqBody != "" && !commons.IsStrInArray(reqs, pm.ReqBody) {
			reqs = append(reqs, pm.ReqBody)
		}
	}
	sort.Strings(reqs)
	return reqs
}
