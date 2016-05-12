package codegen

import (
	"errors"
	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/apidocs"
	"github.com/Jumpscale/go-raml/codegen/date"
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
	APIDocsDir   string // apidocs directory. apidocs won't be generated if it is empty
	withMain     bool
}

type goServer struct {
	server
}

type pythonServer struct {
	server
}

// generate all Go server files
func (gs goServer) generate(dir string) error {
	if err := gs.generateDates(dir); err != nil {
		log.Errorf("generate() failed to generate date files:%v", err)
		return err
	}

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
		return generateFile(gs, serverMainTmplFile, serverMainTmplName, filepath.Join(dir, "main.go"), true)
	}

	return nil
}

// generate all dates files
func (gs goServer) generateDates(dir string) error {
	dates := []struct {
		Type     string
		Format   string
		FileName string
	}{
		{"date-only", "", "date_only.go"},
		{"time-only", "", "time_only.go"},
		{"datetime-only", "", "datetime_only.go"},
		{"datetime", "RFC3339", "datetime.go"},
		{"datetime", "RFC2616", "datetime_rfc2616.go"},
	}
	for _, d := range dates {
		b, err := date.Get(d.Type, d.Format)
		if err != nil {
			return err
		}
		ctx := map[string]interface{}{
			"PackageName": gs.PackageName,
			"Content":     string(b),
		}

		err = generateFile(ctx, "./templates/date.tmpl", "date", filepath.Join(dir, d.FileName), false)
		if err != nil {
			return err
		}
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
		return generateFile(ps, serverPythonMainTmplFile, serverPythonMainTmplName, filepath.Join(dir, "app.py"), true)
	}
	return nil

}

// GenerateServer generates API server files
func GenerateServer(ramlFile, dir, packageName, lang, apiDocsDir string, generateMain bool) error {
	apiDef := new(raml.APIDefinition)
	ramlBytes, err := raml.ParseReadFile(ramlFile, apiDef)
	if err != nil {
		return err
	}

	if err := checkCreateDir(dir); err != nil {
		return err
	}

	sd := server{
		PackageName: packageName,
		apiDef:      apiDef,
		APIDocsDir:  apiDocsDir,
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

	if sd.APIDocsDir == "" {
		return nil
	}

	log.Infof("Generating API Docs to %v endpoint", sd.APIDocsDir)

	return apidocs.Generate(ramlBytes, filepath.Join(dir, sd.APIDocsDir))
}
