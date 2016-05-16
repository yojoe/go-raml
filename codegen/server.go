package codegen

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/apidocs"
	"github.com/Jumpscale/go-raml/raml"

	log "github.com/Sirupsen/logrus"
)

var (
	errInvalidLang = errors.New("invalid language")
)

// global variables
// it is needed for libraries support
var (
	// root import path
	globRootImportPath string

	// global value of API definition
	globAPIDef *raml.APIDefinition
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
	ResourcesDef []resourceInterface
	PackageName  string // Name of the package this server resides in
	APIDocsDir   string // apidocs directory. apidocs won't be generated if it is empty
	withMain     bool
}

type goServer struct {
	server
	RootImportPath string
}

type pythonServer struct {
	server
}

// generate all Go server files
func (gs goServer) generate(dir string) error {
	// helper package
	gh := goramlHelper{
		rootImportPath: gs.RootImportPath,
		packageName:    "goraml",
		packageDir:     "goraml",
	}
	if err := gh.generate(dir); err != nil {
		return err
	}

	// generate all Type structs
	if err := generateStructs(gs.apiDef.Types, dir, gs.PackageName, langGo); err != nil {
		return err
	}

	// generate all request & response body
	if err := generateBodyStructs(gs.apiDef, dir, gs.PackageName, langGo); err != nil {
		return err
	}

	// security scheme
	if err := generateSecurity(gs.apiDef.SecuritySchemes, dir, gs.PackageName, langGo); err != nil {
		log.Errorf("failed to generate security scheme:%v", err)
		return err
	}

	// genereate resources
	rds, err := generateServerResources(gs.apiDef, dir, gs.PackageName, langGo)
	if err != nil {
		return err
	}
	gs.ResourcesDef = rds

	// libraries
	if err := generateLibraries(gs.apiDef.Libraries, dir); err != nil {
		return err
	}

	// generate main
	if gs.withMain {
		return generateFile(gs, serverMainTmplFile, serverMainTmplName, filepath.Join(dir, "main.go"), true)
	}

	return nil
}

func (ps pythonServer) generate(dir string) error {
	// generate input validators helper
	if err := generateFile(struct{}{}, "./templates/input_validators_python.tmpl", "input_validators_python",
		filepath.Join(dir, "input_validators.py"), false); err != nil {
		return err
	}

	// generate request body
	if err := generateBodyStructs(ps.apiDef, dir, "", langPython); err != nil {
		log.Errorf("failed to generate python classes from request body:%v", err)
		return err
	}
	// python classes
	if err := generatePythonClasses(ps.apiDef, dir); err != nil {
		log.Errorf("failed to generate python clased:%v", err)
		return err
	}
	// security scheme
	if err := generateSecurity(ps.apiDef.SecuritySchemes, dir, ps.PackageName, langPython); err != nil {
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
		return generateFile(ps, serverPythonMainTmplFile, serverPythonMainTmplName, filepath.Join(dir, "app.py"), true)
	}
	return nil

}

// GenerateServer generates API server files
func GenerateServer(ramlFile, dir, packageName, lang, apiDocsDir, rootImportPath string, generateMain bool) error {
	apiDef := new(raml.APIDefinition)
	// parse the raml file
	ramlBytes, err := raml.ParseReadFile(ramlFile, apiDef)
	if err != nil {
		return err
	}

	// global variables
	globAPIDef = apiDef
	globRootImportPath = rootImportPath

	// create directory if needed
	if err := checkCreateDir(dir); err != nil {
		return err
	}

	// create base server
	sd := server{
		PackageName: packageName,
		apiDef:      apiDef,
		APIDocsDir:  apiDocsDir,
		withMain:    generateMain,
	}
	switch lang {
	case langGo:
		if rootImportPath == "" {
			return fmt.Errorf("invalid import path = empty")
		}
		gs := goServer{server: sd, RootImportPath: rootImportPath}
		err = gs.generate(dir)
	case langPython:
		ps := pythonServer{server: sd}
		err = ps.generate(dir)
	default:
		return errInvalidLang
	}

	if err != nil {
		return err
	}

	if sd.APIDocsDir == "" {
		return nil
	}

	log.Infof("Generating API Docs to %v endpoint", sd.APIDocsDir)

	return apidocs.Generate(ramlBytes, filepath.Join(dir, sd.APIDocsDir))
}
