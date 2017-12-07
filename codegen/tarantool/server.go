package tarantool

import (
	"github.com/Jumpscale/go-raml/codegen/capnp"
	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
	"path"
	"sort"
	"os/exec"
	"fmt"
)

// Server represents a tarantool server
type Server struct {
	apiDef     *raml.APIDefinition
	APIDocsDir string // apidocs directory. apidocs won't be generated if it is empty
	TargetDir  string // root directory of the generated code
	Resources  TarantoolResources
}

// NewServer creates a new tarantool server
func NewServer(apiDef *raml.APIDefinition, apiDocsDir string, targetDir string) Server {
	return Server{
		apiDef:     apiDef,
		APIDocsDir: apiDocsDir,
		TargetDir:  targetDir,
	}
}

// Generate generates all tarantool server files
func (s *Server) Generate() error {
	err := s.generateSchemas()
	if err != nil {
		return err
	}

	s.generateResources()

	err = s.generateMain()
	if err != nil {
		return err
	}

	err = s.generateHandlers()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) generateSchemas() error {
	err := capnp.GenerateCapnp(s.apiDef, path.Join(s.TargetDir, "capnp"), "plain", "")
	if err != nil {
		return err
	}
	// generate types
	for name := range s.apiDef.Types {
		input := path.Join(s.TargetDir, "capnp", fmt.Sprintf("%v.capnp", name))
		cmd := exec.Command(fmt.Sprintf("capnp compile -olua %v", input))
		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

// generateResources generates a list of tarantool resources
func (s *Server) generateResources() {
	resources := flattenResources(s.apiDef.Resources)

	sort.Sort(resources.Resources)
	s.Resources = resources
}

// generateMain generates the main server file
func (s *Server) generateMain() error {
	filename := s.TargetDir + "/" + "main.lua"

	// error helper
	ctx := struct {
		Resources TarantoolResources
	}{
		Resources: s.Resources,
	}
	return commons.GenerateFile(ctx, "./templates/tarantool/server_main.tmpl", "server_main", filename, true)
}

// generateHandlers generates all endpoint handlers
func (s *Server) generateHandlers() error {
	for _, resource := range s.Resources.Resources {
		filename := path.Join(s.TargetDir, "handlers", resource.Handler()+".lua")

		// error helper
		ctx := struct {
			Resource *Resource
		}{
			Resource: resource,
		}
		err := commons.GenerateFile(ctx, "./templates/tarantool/server_resource_handler.tmpl", "server_resource_handler", filename, true)
		if err != nil {
			return err
		}
	}

	return nil
}
