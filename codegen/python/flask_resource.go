package python

import (
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/resource"
)

func (fs FlaskServer) generateResources(dir string) error {
	for _, rdi := range fs.ResourcesDef {

		pr := newResourceFromDef(rdi.(resource.Resource), fs.APIDef, newServerMethodFlask)

		filename := filepath.Join(dir, strings.ToLower(pr.Name)+".py")

		if err := pr.generate(filename, "./templates/python_server_resource.tmpl", "resource_python_template", dir); err != nil {
			return err
		}
	}
	return nil
}
