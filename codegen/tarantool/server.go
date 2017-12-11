package tarantool

import (
	"os/exec"
	"path"
	"sort"

	"fmt"

	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/capnp"
	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
	"os"
	"strings"
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
	schemasPath := path.Join(s.TargetDir, "schemas")
	err := capnp.GenerateCapnp(s.apiDef, schemasPath, "plain", "")
	if err != nil {
		return err
	}

	args, err := filepath.Glob(path.Join(schemasPath, "*.capnp"))
	if err != nil {
		return err
	}

	// lua generator used the name of the first file in the list to determine the name of the compiled lua file
	compiledFile := strings.Replace(args[0], ".capnp", "_capnp.lua", 1)

	args = append([]string{"compile", "-olua"}, args...)
	cmd := exec.Command("capnp", args...)
	err = cmd.Run()
	if err != nil {
		return err
	}

	if _, err = os.Stat(compiledFile); os.IsNotExist(err) {
		return fmt.Errorf(
			"Can't find the compiled lua schema file expected at %v", compiledFile)
	}

	// rename the compiled lua file to a more generic name
	err = os.Rename(compiledFile, path.Join(
		schemasPath, "schema.lua"))
	if err != nil {
		return err
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
