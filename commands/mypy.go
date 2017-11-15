package commands

import (
	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen"
	"github.com/Jumpscale/go-raml/raml"
)

// CapnpCommand is executed to generate capnpm model from RAML specification
type MyPyCommand struct {
	Dir      string //target dir
	RAMLFile string //raml file
}

//Execute generates a client from a RAML specification
func (command *MyPyCommand) Execute() error {
	var apiDef raml.APIDefinition

	log.Debug("Generating mypy models")

	err := raml.ParseFile(command.RAMLFile, &apiDef)
	if err != nil {
		return err
	}
	return codegen.GenerateMyPy(&apiDef, command.Dir)
}
