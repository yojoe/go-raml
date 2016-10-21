package nim

import (
	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

// Server represents a Nim server
type Server struct {
	APIDef     *raml.APIDefinition
	Dir        string
	Title      string
	APIDocsDir string
	Resources  []resource
}

// Generate generates all Nim server files
func (s *Server) Generate() error {
	s.Resources = getAllResources(s.APIDef)

	// generate all objects from all RAML types
	if err := generateObjects(s.APIDef.Types, s.Dir); err != nil {
		return err
	}

	// generate all objects from request/response body
	if _, err := generateObjectsFromBodies(s.Resources, s.Dir); err != nil {
		return err
	}

	// main file
	if err := s.generateMain(); err != nil {
		return err
	}

	// API implementation
	if err := s.generateResourceAPIs(); err != nil {
		return err
	}

	// HTML front page
	if err := commons.GenerateFile(s, "./templates/index.html.tmpl", "index.html", filepath.Join(s.Dir, "index.html"), false); err != nil {
		return err
	}

	return nil
}

// main generates main file
func (s *Server) generateMain() error {
	filename := filepath.Join(s.Dir, "main.nim")
	return commons.GenerateFile(s, "./templates/server_main_nim.tmpl", "server_main_nim", filename, true)
}

// Imports returns array of modules that need to be imported by server's main file
func (s *Server) Imports() []string {
	imports := map[string]struct{}{}

	for _, r := range s.Resources {
		imports[r.apiName()] = struct{}{}
	}
	return commons.MapToSortedStrings(imports)
}

func (s *Server) generateResourceAPIs() error {
	for _, r := range s.Resources {
		if err := r.generate(s.Dir); err != nil {
			return err
		}
	}
	return nil
}
