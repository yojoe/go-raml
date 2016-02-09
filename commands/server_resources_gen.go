package commands

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

const (
	resourceIfTemplate  = "./templates/server_resources_interface.tmpl" // resource interface template
	resourceAPITemplate = "./templates/server_resources_api.tmpl"       // resource API template
)

// resource interface filename
func (rd *resourceDef) ifFileName(directory string) string {
	return directory + "/" + strings.ToLower(rd.Name) + "_if.go"
}

// resource API filename
func (rd *resourceDef) apiFileName(directory string) string {
	return directory + "/" + strings.ToLower(rd.Name) + "_api.go"
}

// generate interface file of a resource
func (rd *resourceDef) generateInterfaceFile(directory string) error {
	return generateFile(rd, resourceIfTemplate, "resource_if_template", rd.ifFileName(directory), true)
}

// generate API file of a resource
func (rd *resourceDef) generateAPIFile(directory string) error {
	return generateFile(rd, resourceAPITemplate, "resource_api_template", rd.apiFileName(directory), false)
}

// generate Go representation of server's resource
func (rd *resourceDef) generate(r *raml.Resource, URI, dir string) error {
	rd.generateMethods(r, "", URI)
	if err := rd.generateInterfaceFile(dir); err != nil {
		return err
	}
	return rd.generateAPIFile(dir)
}

// ServerResourceGen generate Server's Go representation of RAML resource
func ServerResourcesGen(rs map[string]raml.Resource, directory string) ([]resourceDef, error) {
	var rds []resourceDef

	if err := checkCreateDir(directory); err != nil {
		return rds, err
	}
	// create resource def
	for k, r := range rs {
		rd := newResourceDef(k)
		rd.IsServer = true
		if err := rd.generate(&r, k, directory); err != nil {
			return rds, err
		}
		rds = append(rds, rd)
	}
	return rds, nil
}
