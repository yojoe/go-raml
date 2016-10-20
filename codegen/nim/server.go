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
	objNames, err := GenerateObjects(s.APIDef.Types, s.Dir)
	if err != nil {
		return err
	}
	addGeneratedObjects(objNames)

	if err := s.generateMain(); err != nil {
		return err
	}

	if err := s.generateResourceAPIs(); err != nil {
		return err
	}
	return nil
}

// main generates main file
func (s *Server) generateMain() error {
	filename := filepath.Join(s.Dir, "main.nim")
	return commons.GenerateFile(s, "./templates/server_main_nim.tmpl", "server_main_nim", filename, true)
}

func (s *Server) Imports() []string {
	imports := []string{}

	for _, r := range s.Resources {
		imports = append(imports, r.apiName())
	}
	return imports
}

func (s *Server) generateResourceAPIs() error {
	for _, r := range s.Resources {
		if err := r.generate(s.Dir); err != nil {
			return err
		}
	}
	return nil
}
