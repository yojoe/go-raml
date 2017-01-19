package python

import (
	"github.com/Jumpscale/go-raml/codegen/resource"
)

func (fs FlaskServer) generateResources(dir string) error {
	for _, rdi := range fs.ResourcesDef {
		rd := rdi.(resource.Resource)
		pr := pythonResource{
			Resource: &rd,
		}
		res := fs.APIDef.Resources[pr.Endpoint]
		if err := pr.generate(&res, pr.Endpoint, dir); err != nil {
			return err
		}
	}
	return nil
}
