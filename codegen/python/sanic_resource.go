package python

import (
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/resource"
)

func (s SanicServer) generateResources(dir string) error {
	for _, rdi := range s.ResourcesDef {

		pr := newResourceFromDef(rdi.(resource.Resource), s.APIDef)

		filename := filepath.Join(dir, strings.ToLower(pr.Name)+"_api.py")

		// API implementation
		if err := pr.generate(filename, "./templates/server_resources_api_python_sanic.tmpl", "server_resources_api_python_sanic", dir); err != nil {
			return err
		}
	}
	return nil
}

type sanicRouteView struct {
}
