package python

import (
	"path/filepath"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

// FlaskServer represents a flask server
type FlaskServer struct {
	APIDef       *raml.APIDefinition
	Title        string
	ResourcesDef []resource.ResourceInterface
	WithMain     bool
	APIDocsDir   string
}

func NewFlaskServer(apiDef *raml.APIDefinition, apiDocsDir string, withMain bool) *FlaskServer {
	var rds []resource.ResourceInterface
	for _, rd := range getServerResourcesDefs(apiDef) {
		rds = append(rds, rd)
	}
	return &FlaskServer{
		APIDef:       apiDef,
		Title:        apiDef.Title,
		APIDocsDir:   apiDocsDir,
		WithMain:     withMain,
		ResourcesDef: rds,
	}
}

// Generate generates all python server files
func (ps FlaskServer) Generate(dir string) error {

	globAPIDef = ps.APIDef
	// generate input validators helper
	if err := commons.GenerateFile(struct{}{}, "./templates/input_validators_python.tmpl", "input_validators_python",
		filepath.Join(dir, "input_validators.py"), false); err != nil {
		return err
	}

	// generate request body
	if err := generateClassesFromBodies(getAllResources(ps.APIDef, true), dir); err != nil {
		return err
	}

	// python classes
	if err := generateClasses(ps.APIDef.Types, dir); err != nil {
		log.Errorf("failed to generate python clased:%v", err)
		return err
	}

	// security scheme
	if err := ps.generateOauth2(ps.APIDef.SecuritySchemes, dir); err != nil {
		log.Errorf("failed to generate security scheme:%v", err)
		return err
	}

	// genereate resources
	if err := ps.generateResources(dir); err != nil {
		return err
	}

	// libraries
	if err := generateLibraries(ps.APIDef.Libraries, dir); err != nil {
		return err
	}

	// requirements.txt file
	if err := commons.GenerateFile(nil, "./templates/requirements_python.tmpl", "requirements_python", filepath.Join(dir, "requirements.txt"), false); err != nil {
		return err
	}

	// generate main
	if ps.WithMain {
		// generate HTML front page
		if err := commons.GenerateFile(ps, "./templates/index.html.tmpl", "index.html", filepath.Join(dir, "index.html"), false); err != nil {
			return err
		}
		// main file
		return commons.GenerateFile(ps, "./templates/server_main_python.tmpl", "server_main_python", filepath.Join(dir, "app.py"), true)
	}
	return nil

}
