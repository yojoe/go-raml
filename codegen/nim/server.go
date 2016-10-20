package nim

import (
	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

type Server struct {
	APIDef     *raml.APIDefinition
	Dir        string
	Title      string
	APIDocsDir string
	Resources  []resource
}

func (s *Server) Generate() error {
	s.Resources = getAllResources(s.APIDef)

	// generate all objects
	if err := GenerateObjects(s.APIDef.Types, s.Dir); err != nil {
		return err
	}

	// main file
	if err := s.generateMain(); err != nil {
		return err
	}
	return nil
}

// main generates main file
func (s *Server) generateMain() error {
	filename := filepath.Join(s.Dir, "main.nim")
	return commons.GenerateFile(s, "./templates/server_main_nim.tmpl", "server_main_nim", filename, true)
}
