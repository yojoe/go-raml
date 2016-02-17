package commands

import (
	"github.com/Jumpscale/go-raml/raml"
)

const (
	serverMainTmplFile = "./templates/server_main.tmpl"
	serverMainTmplName = "server_main_template"
)

// API server definition
type serverDef struct {
	ResourcesDef []resourceDef
	PackageName  string // Name of the package this server resides in
}

// generate server main file
func (sd serverDef) generate(dir string) error {
	return generateFile(sd, serverMainTmplFile, serverMainTmplName, dir+"/main.go", true)
}

// generate API server files
func generateServer(apiDef *raml.APIDefinition, dir, packageName string) error {
	// generate all Type structs
	if err := generateStructs(apiDef, dir, packageName); err != nil {
		return err
	}

	// generate all request & response body
	if err := generateBodyStructs(apiDef, dir, packageName); err != nil {
		return err
	}

	// genereate resources
	rds, err := generateServerResources(apiDef.Resources, dir, packageName)
	if err != nil {
		return err
	}

	sd := serverDef{ResourcesDef: rds, PackageName: packageName}

	return sd.generate(dir)
}
