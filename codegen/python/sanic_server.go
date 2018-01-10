package python

import (
	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

// SanicServer represents a sanic asyncio server
type SanicServer struct {
	APIDef       *raml.APIDefinition
	Title        string
	ResourcesDef []pythonResource
	withMain     bool
	APIDocsDir   string
	Template     serverTemplate
}

// NewSanicServer creates new sanic server from an RAML file
func NewSanicServer(apiDef *raml.APIDefinition, apiDocsDir string, withMain bool,
	libRootURLs []string) *SanicServer {
	var prs []pythonResource
	for _, rd := range getServerResourcesDefs(apiDef) {
		pr := newResource(rd, apiDef, serverKindSanic)
		prs = append(prs, pr)
	}

	// TODO : get rid of this global variables
	globAPIDef = apiDef
	globLibRootURLs = libRootURLs

	templates := templates(serverKindSanic)

	return &SanicServer{
		APIDef:       apiDef,
		Title:        apiDef.Title,
		APIDocsDir:   apiDocsDir,
		withMain:     withMain,
		ResourcesDef: prs,
		Template:     templates,
	}
}

// Generate generates sanic server code
func (s *SanicServer) Generate(dir string) error {
	if err := generateJSONSchema(s.APIDef, filepath.Join(dir, handlersDir())); err != nil {
		return err
	}
	if err := s.generateResources(dir); err != nil {
		return err
	}
	if err := s.generateOauth2(s.APIDef.SecuritySchemes, dir); err != nil {
		return err
	}

	// python classes and it's helper
	if err := commons.GenerateFile(nil, "./templates/python/client_support.tmpl",
		"client_support", filepath.Join(dir, "types", "client_support.py"), true); err != nil {
		return err
	}

	_, err := GenerateAllClasses(s.APIDef, filepath.Join(dir, typesDir()), false)
	if err != nil {
		return err
	}

	return s.generateMain(dir)
}

func (s *SanicServer) generateMain(dir string) error {
	if !s.withMain {
		return nil
	}

	// html front page
	if err := commons.GenerateFile(s, "./templates/index.html.tmpl", "index.html",
		filepath.Join(dir, "index.html"), true); err != nil {
		return err
	}

	return commons.GenerateFile(s, "./templates/python/server_main_sanic.tmpl", "server_main_sanic",
		filepath.Join(dir, "app.py"), true)
}
