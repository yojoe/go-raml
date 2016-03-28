package commands

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

const (
	resourcePyTemplate = "./templates/python_server_resource.tmpl"
)

// generate flask representation of an RAML resource
// It has one file : an API route and implementation
func (rd *resourceDef) generatePython(r *raml.Resource, URI, dir string) error {
	rd.generateMethods(r, "", URI, "python")
	filename := dir + "/" + strings.ToLower(rd.Name) + ".py"
	return generateFile(rd, resourcePyTemplate, "resource_python_template", filename, true)
}
