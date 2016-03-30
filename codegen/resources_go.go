package codegen

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

const (
	resourceIfTemplate  = "./templates/server_resources_interface.tmpl" // resource interface template
	resourceAPITemplate = "./templates/server_resources_api.tmpl"       // resource API template
)

type goResource struct {
	*resourceDef
	WithMiddleware bool // this resource need middleware, we need to import github/justinas/alice
	NeedJSON       bool // if true, the API implementation to import encoding/json package
}

// generate interface file of a resource
func (gr *goResource) generateInterfaceFile(directory string) error {
	filename := directory + "/" + strings.ToLower(gr.Name) + "_if.go"
	return generateFile(gr, resourceIfTemplate, "resource_if_template", filename, true)
}

// generate API file of a resource
func (gr *goResource) generateAPIFile(directory string) error {
	filename := directory + "/" + strings.ToLower(gr.Name) + "_api.go"
	return generateFile(gr, resourceAPITemplate, "resource_api_template", filename, false)
}

// generate Go representation of server's resource.
// A resource have two kind of files:
// - interface file
//		contains API interface and routing code
//		always regenerated
// - API implementation
//		implementation of the API interface.
//		Don't generate if the file already exist
func (gr *goResource) generate(r *raml.Resource, URI, dir string) error {
	gr.generateMethods(r, "", URI, "go")
	gr.setImport()
	if err := gr.generateInterfaceFile(dir); err != nil {
		return err
	}
	return gr.generateAPIFile(dir)
}

func (gr *goResource) setImport() {
	for _, v := range gr.Methods {
		gm := v.(goServerMethod)

		// if there is request/response body, then it needs to import encoding/json
		if gm.RespBody != "" || gm.ReqBody != "" {
			gr.NeedJSON = true
		}

		// if has middleware, we need to import middleware lib
		if len(gm.Middlewares) > 0 {
			gr.WithMiddleware = true
		}
	}
}
