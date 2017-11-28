package python

import (
	"path/filepath"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

// FlaskServer represents a flask server
type FlaskServer struct {
	APIDef       *raml.APIDefinition
	Title        string
	ResourcesDef []pythonResource
	WithMain     bool
	APIDocsDir   string
}

// NewFlaskServer creates new flask server from an RAML file
func NewFlaskServer(apiDef *raml.APIDefinition, apiDocsDir string,
	withMain bool, libRootURLs []string) *FlaskServer {

	// TODO : get rid of this global variables
	globAPIDef = apiDef
	globLibRootURLs = libRootURLs

	// generates resource
	var prs []pythonResource
	for _, rd := range getServerResourcesDefs(apiDef) {
		pr := newResource(rd, apiDef, serverKindFlask)
		prs = append(prs, pr)
	}

	return &FlaskServer{
		APIDef:       apiDef,
		Title:        apiDef.Title,
		APIDocsDir:   apiDocsDir,
		WithMain:     withMain,
		ResourcesDef: prs,
	}
}

// Generate generates all python server files
func (ps FlaskServer) Generate(dir string) error {
	if err := generateJSONSchema(ps.APIDef, dir); err != nil {
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
	if err := commons.GenerateFile(nil, "./templates/python/requirements_python.tmpl", "requirements_python", filepath.Join(dir, "requirements.txt"), false); err != nil {
		return err
	}

	// generate main
	if ps.WithMain {
		// generate HTML front page
		if err := commons.GenerateFile(ps, "./templates/index.html.tmpl", "index.html", filepath.Join(dir, "index.html"), false); err != nil {
			return err
		}
		// main file
		return commons.GenerateFile(ps, "./templates/python/server_main_python.tmpl", "server_main_python", filepath.Join(dir, "app.py"), true)
	}
	return nil

}
