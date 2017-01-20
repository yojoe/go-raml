package python

import (
	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

type SanicServer struct {
	APIDef       *raml.APIDefinition
	Title        string
	ResourcesDef []resource.ResourceInterface
	withMain     bool
	APIDocsDir   string
}

func NewSanicServer(apiDef *raml.APIDefinition, apiDocsDir string, withMain bool) *SanicServer {
	var rds []resource.ResourceInterface
	for _, rd := range getServerResourcesDefs(apiDef) {
		rds = append(rds, rd)
	}

	return &SanicServer{
		APIDef:       apiDef,
		Title:        apiDef.Title,
		APIDocsDir:   apiDocsDir,
		withMain:     withMain,
		ResourcesDef: rds,
	}
}

func (s *SanicServer) Generate(dir string) error {
	if err := s.generateJSONSchema(dir); err != nil {
		return err
	}
	if err := s.generateResources(dir); err != nil {
		return err
	}
	if err := s.generateOauth2(s.APIDef.SecuritySchemes, dir); err != nil {
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
		filepath.Join(dir, "index.html"), false); err != nil {
		return err
	}

	return commons.GenerateFile(s, "./templates/server_main_python_sanic.tmpl", "server_main_python_sanic",
		filepath.Join(dir, "app.py"), true)
}
