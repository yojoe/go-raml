package codegen

import (
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

const (
	resourcePyTemplate = "./templates/python_server_resource.tmpl"
)

type pythonResource struct {
	*resource.Resource
	MiddlewaresArr []pythonMiddleware
}

func (pr *pythonResource) addMiddleware(mwr pythonMiddleware) {
	// check if already exist
	for _, v := range pr.MiddlewaresArr {
		if v.Name == mwr.Name {
			return
		}
	}
	pr.MiddlewaresArr = append(pr.MiddlewaresArr, mwr)
}

// set middlewares to import
func (pr *pythonResource) setMiddlewares() {
	for _, v := range pr.Methods {
		pm := v.(pythonServerMethod)
		for _, m := range pm.MiddlewaresArr {
			pr.addMiddleware(m)
		}
	}
}

// generate flask representation of an RAML resource
// It has one file : an API route and implementation
func (pr *pythonResource) generate(r *raml.Resource, URI, dir string) error {
	pr.GenerateMethods(r, "python", newServerMethod, newClientMethod)
	pr.setMiddlewares()
	filename := dir + "/" + strings.ToLower(pr.Name) + ".py"
	return commons.GenerateFile(pr, resourcePyTemplate, "resource_python_template", filename, true)
}

// return array of request body in this resource
func (pr pythonResource) ReqBodies() []string {
	var reqs []string
	for _, m := range pr.Methods {
		pm := m.(pythonServerMethod)
		if pm.ReqBody != "" {
			reqs = append(reqs, pm.ReqBody)
		}
	}
	return reqs
}
