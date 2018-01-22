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
	Gevent       bool
	Template     serverTemplate
}

// NewFlaskServer creates new flask server from an RAML file
func NewFlaskServer(apiDef *raml.APIDefinition, apiDocsDir string,
	withMain bool, libRootURLs []string, gevent bool) *FlaskServer {

	// TODO : get rid of this global variables
	globAPIDef = apiDef
	globLibRootURLs = libRootURLs

	// generates resource
	var prs []pythonResource
	for _, rd := range getServerResourcesDefs(apiDef) {
		pr := newResource(rd, apiDef, serverKindFlask)
		prs = append(prs, pr)
	}

	templates := templates(serverKindFlask)

	return &FlaskServer{
		APIDef:       apiDef,
		Title:        apiDef.Title,
		APIDocsDir:   apiDocsDir,
		WithMain:     withMain,
		ResourcesDef: prs,
		Gevent:       gevent,
		Template:     templates,
	}
}

// Generate generates all python server files
func (ps *FlaskServer) Generate(dir string) error {
	// python classes and it's helper
	typesPath := filepath.Join(dir, typesDir)

	if err := commons.GenerateFile(nil, "./templates/python/client_support.tmpl",
		"client_support", filepath.Join(typesPath, "client_support.py"), true); err != nil {
		return err
	}

	if ps.Gevent {
		if err := commons.GenerateFile(ps, "./templates/python/server_gevent.tmpl", "server_gevent", filepath.Join(dir, "server.py"), true); err != nil {
			return err
		}

		if err := GeneratePythonCapnpClasses(ps.APIDef, typesPath); err != nil {
			return nil
		}
	} else {
		_, err := GenerateAllClasses(ps.APIDef, typesPath, ps.Gevent)
		if err != nil {
			return err
		}
	}

	// json schema
	if err := generateJSONSchema(ps.APIDef, filepath.Join(dir, handlersDir)); err != nil {
		return err
	}

	// security scheme
	if err := generateServerSecurity(ps.APIDef.SecuritySchemes, ps.Template, dir); err != nil {
		log.Errorf("failed to generate security scheme:%v", err)
		return err
	}

	// generate resources
	if err := generateResources(ps.ResourcesDef, ps.Template, dir); err != nil {
		return err
	}

	// libraries
	if err := generateLibraries(ps.APIDef.Libraries, dir); err != nil {
		return err
	}

	// requirements.txt file
	if err := commons.GenerateFile(nil, "./templates/python/requirements_python.tmpl", "requirements_python",
		filepath.Join(dir, "requirements.txt"), true); err != nil {
		return err
	}

	// generate main
	if ps.WithMain {
		// generate HTML front page
		if err := commons.GenerateFile(ps, "./templates/index.html.tmpl", "index.html", filepath.Join(dir, "index.html"), true); err != nil {
			return err
		}
		// main file
		if err := commons.GenerateFile(ps, "./templates/python/server_main_flask.tmpl", "server_main_flask", filepath.Join(dir, "app.py"), true); err != nil {
			return err
		}
	}

	return generateEmptyInitPy(dir)
}
