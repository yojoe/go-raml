package commands

import (
	"github.com/Jumpscale/go-raml/raml"

	log "github.com/Sirupsen/logrus"
)

const (
	serverMainTmplFile       = "./templates/server_main.tmpl"
	serverMainTmplName       = "server_main_template"
	serverPythonMainTmplFile = "./templates/server_python_main.tmpl"
	serverPythonMainTmplName = "server_python_main_template"
)

// API server definition
type serverDef struct {
	ResourcesDef []resourceDef
	PackageName  string // Name of the package this server resides in
}

// generate Go server main file
func (sd serverDef) generateGo(dir string) error {
	return generateFile(sd, serverMainTmplFile, serverMainTmplName, dir+"/main.go", true)
}

func (sd serverDef) generatePython(dir string) error {
	return generateFile(sd, serverPythonMainTmplFile, serverPythonMainTmplName, dir+"/app.py", true)
}

func (sd serverDef) generate(dir, lang string) error {
	if lang == "go" {
		return sd.generateGo(dir)
	}
	return sd.generatePython(dir)
}

// generate API server files
func generateServer(apiDef *raml.APIDefinition, dir, packageName, lang string, generateMain bool) error {

	if err := checkCreateDir(dir); err != nil {
		return err
	}

	if lang == "go" {
		// generate struct validator
		if err := generateInputValidator(packageName, dir); err != nil {
			return err
		}

		// generate all Type structs
		if err := generateStructs(apiDef, dir, packageName); err != nil {
			return err
		}

		// generate all request & response body
		if err := generateBodyStructs(apiDef, dir, packageName); err != nil {
			return err
		}
	}

	// security scheme
	if err := generateSecurity(apiDef, dir, packageName, lang); err != nil {
		log.Errorf("failed to generate security scheme:%v", err)
		return err
	}

	// genereate resources
	rds, err := generateServerResources(apiDef, dir, packageName, lang)
	if err != nil {
		return err
	}

	if generateMain {
		sd := serverDef{ResourcesDef: rds, PackageName: packageName}
		err = sd.generate(dir, lang)
	}
	return err
}
