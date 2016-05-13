package codegen

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

type goClient struct {
	clientDef
}

// generate Go client files
func (gc goClient) generate(apiDef *raml.APIDefinition, dir string) error {
	// generate struct
	if err := generateStructs(apiDef.Types, dir, "client", langGo); err != nil {
		return err
	}

	// generate strucs from bodies
	if err := generateBodyStructs(apiDef, dir, "client", langGo); err != nil {
		return err
	}

	if err := gc.generateHelperFile(dir); err != nil {
		return err
	}
	return gc.generateClientFile(dir)
}

// generate Go client helper
func (gc *goClient) generateHelperFile(dir string) error {
	fileName := dir + "/client_utils.go"
	return generateFile(gc, clientHelperResourceTemplate, "client_helper_resources", fileName, false)
}

// generate Go client lib file
func (gc *goClient) generateClientFile(dir string) error {
	fileName := dir + "/client_" + strings.ToLower(gc.Name) + ".go"
	return generateFile(gc, clientResourceTemplate, "client_resource", fileName, false)
}
