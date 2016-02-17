package commands

import (
	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
)

// ServerCommand is executed to generate a go server from a RAML specification
type ServerCommand struct {
	Dir         string //target dir
	RamlFile    string //raml file
	PackageName string //package name in the generated go source files
}

// Execute generates a Go server from an RAML specification
func (command *ServerCommand) Execute() error {
	log.Debug("Generating a go server")
	apiDef, err := raml.ParseFile(command.RamlFile)
	if err != nil {
		return err
	}
	return generateServer(apiDef, command.Dir, command.PackageName)
}
