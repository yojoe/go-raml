package commands

import log "github.com/Sirupsen/logrus"

//ServerCommand is executed to generate a go server from a RAML specification
type ServerCommand struct{}

//Execute generates a go server from a RAML specification
func (command *ServerCommand) Execute() error {
	log.Debug("Generating a go server")
	return nil
}
