package commands

import (
	"sort"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

const (
	resourceIfTemplate  = "./templates/server_resources_interface.tmpl" // resource interface template
	resourceAPITemplate = "./templates/server_resources_api.tmpl"       // resource API template
)

// generate interface file of a resource
func (rd *resourceDef) generateInterfaceFile(directory string) error {
	filename := directory + "/" + strings.ToLower(rd.Name) + "_if.go"
	return generateFile(rd, resourceIfTemplate, "resource_if_template", filename, true)
}

// generate API file of a resource
func (rd *resourceDef) generateAPIFile(directory string) error {
	filename := directory + "/" + strings.ToLower(rd.Name) + "_api.go"
	return generateFile(rd, resourceAPITemplate, "resource_api_template", filename, false)
}

// generate Go representation of server's resource.
// A resource have two kind of files:
// - interface file
//		contains API interface and routing code
//		always regenerated
// - API implementation
//		implementation of the API interface.
//		Don't generate if the file already exist
func (rd *resourceDef) generateGo(r *raml.Resource, URI, dir string) error {
	rd.generateMethods(r, "", URI, "go")
	if err := rd.generateInterfaceFile(dir); err != nil {
		return err
	}
	return rd.generateAPIFile(dir)
}

func (rd *resourceDef) generate(r *raml.Resource, URI, dir, lang string) error {
	if lang == langGo {
		return rd.generateGo(r, URI, dir)
	}
	return rd.generatePython(r, URI, dir)
}

// generate Server's Go representation of RAML resources
func generateServerResources(apiDef *raml.APIDefinition, directory, packageName, lang string) ([]resourceDef, error) {
	var rds []resourceDef

	rs := apiDef.Resources

	if err := checkCreateDir(directory); err != nil {
		return rds, err
	}
	// sort the keys, so we have resource sorted by keys.
	// the generated code actually don't need it to be sorted.
	// but test fixture need it
	var keys []string
	for k := range rs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// create resource def
	for _, k := range keys {
		r := rs[k]
		rd := newResourceDef(apiDef, k, packageName)
		rd.IsServer = true
		if err := rd.generate(&r, k, directory, lang); err != nil {
			return rds, err
		}
		rds = append(rds, rd)
	}
	return rds, nil
}
