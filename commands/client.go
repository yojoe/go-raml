package commands

import log "github.com/Sirupsen/logrus"

//ClientCommand is executed to generate client from a RAML specification
type ClientCommand struct {
	Language string
}

//Execute generates a client from a RAML specification
func (command *ClientCommand) Execute() error {
	log.Debug("Generating a rest client for ", command.Language)
	return nil
}
