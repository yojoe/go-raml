package commands

import (
	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
)

//ClientCommand is executed to generate client from a RAML specification
type ClientCommand struct {
	Language string
	Dir      string //target dir
	RamlFile string //raml file
}

//Execute generates a client from a RAML specification
func (command *ClientCommand) Execute() error {
	log.Debug("Generating a rest client for ", command.Language)
	apiDef, err := raml.ParseFile(command.RamlFile)
	if err != nil {
		return err
	}
	return GenerateClient(apiDef, command.Dir, command.Language)
}
