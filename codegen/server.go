package codegen

import (
	"errors"

	"github.com/Jumpscale/go-raml/raml"

	log "github.com/Sirupsen/logrus"
)

var (
	errInvalidLang = errors.New("invalid language")
)

const (
	serverMainTmplFile       = "./templates/server_main.tmpl"
	serverMainTmplName       = "server_main_template"
	serverPythonMainTmplFile = "./templates/server_python_main.tmpl"
	serverPythonMainTmplName = "server_python_main_template"
)

// base server definition
type server struct {
	apiDef       *raml.APIDefinition
	ResourcesDef []resourceDef
	PackageName  string // Name of the package this server resides in
	withMain     bool
}

type goServer struct {
	server
}

type pythonServer struct {
	server
}

// generate Go server files
func (gs goServer) generate(dir string) error {
	// generate struct validator
	if err := generateInputValidator(gs.PackageName, dir); err != nil {
		return err
	}

	// generate all Type structs
	if err := generateStructs(gs.apiDef, dir, gs.PackageName); err != nil {
		return err
	}

	// generate all request & response body
	if err := generateBodyStructs(gs.apiDef, dir, gs.PackageName); err != nil {
		return err
	}

	// security scheme
	if err := generateSecurity(gs.apiDef, dir, gs.PackageName, langGo); err != nil {
		log.Errorf("failed to generate security scheme:%v", err)
		return err
	}

	// genereate resources
	rds, err := generateServerResources(gs.apiDef, dir, gs.PackageName, langGo)
	if err != nil {
		return err
	}
	gs.ResourcesDef = rds

	// generate main
	if gs.withMain {
		return generateFile(gs, serverMainTmplFile, serverMainTmplName, dir+"/main.go", true)
	}

	return nil
}

func (ps pythonServer) generate(dir string) error {
	// security scheme
	if err := generateSecurity(ps.apiDef, dir, ps.PackageName, langPython); err != nil {
		log.Errorf("failed to generate security scheme:%v", err)
		return err
	}

	// genereate resources
	rds, err := generateServerResources(ps.apiDef, dir, ps.PackageName, langPython)
	if err != nil {
		return err
	}
	ps.ResourcesDef = rds

	// generate main
	if ps.withMain {
		return generateFile(ps, serverPythonMainTmplFile, serverPythonMainTmplName, dir+"/app.py", true)
	}
	return nil

}

// GenerateServer generates API server files
func GenerateServer(apiDef *raml.APIDefinition, dir, packageName, lang string, generateMain bool) error {
	if err := checkCreateDir(dir); err != nil {
		return err
	}

	sd := server{
		PackageName: packageName,
		apiDef:      apiDef,
		withMain:    generateMain,
	}
	switch lang {
	case langGo:
		gs := goServer{server: sd}
		return gs.generate(dir)
	case langPython:
		ps := pythonServer{server: sd}
		return ps.generate(dir)
	default:
		return errInvalidLang
	}
}
