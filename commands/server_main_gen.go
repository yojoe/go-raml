package commands

import (
	"github.com/Jumpscale/go-raml/raml"
)

const (
	serverMainTmplFile = "./templates/server_main.tmpl"
	serverMainTmplName = "server_main_template"
)

type serverDef struct {
	ResourcesDef []resourceDef
}

func (sd serverDef) generate(dir string) error {
	return generateFile(sd, serverMainTmplFile, serverMainTmplName, dir+"/main.go", true)
}

func ServerMainGen(apiDef *raml.APIDefinition, dir string) error {
	// generate all Type structs
	if err := GenerateStruct(apiDef, dir, "main"); err != nil {
		return err
	}

	// generate all request & response body
	if err := GenerateBodyStruct(apiDef, dir, "main"); err != nil {
		return err
	}

	// genereate resources
	rds, err := ServerResourcesGen(apiDef.Resources, dir)
	if err != nil {
		return err
	}

	sd := serverDef{ResourcesDef: rds}

	return sd.generate(dir)
}
