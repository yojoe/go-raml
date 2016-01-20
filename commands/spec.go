package commands

import log "github.com/Sirupsen/logrus"

//SpecCommand is executed to generate a RAML specification from a go server
type SpecCommand struct{}

//Execute generates a RAML specification from a go server
func (command *SpecCommand) Execute() error {
	log.Debug("Generating RAML specification")
	return nil
}
