package codegen

import (
	"errors"
	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/apidocs"
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
	ResourcesDef []resourceInterface
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
	// generate docs if needed

	// generate struct validator
	if err := generateInputValidator(gs.PackageName, dir); err != nil {
		return err
	}

	// generate all Type structs
	if err := generateStructs(gs.apiDef, dir, gs.PackageName, langGo); err != nil {
		return err
	}

	// generate all request & response body
	if err := generateBodyStructs(gs.apiDef, dir, gs.PackageName, langGo); err != nil {
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
func GenerateServer(ramlFile, dir, packageName, lang string, generateMain bool) error {
	ramlBytes, apiDef, err := raml.ParseReadFile(ramlFile)
	if err != nil {
		return err
	}

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

	return apidocs.Create(ramlBytes, filepath.Join(dir, "apidocs"))
}
