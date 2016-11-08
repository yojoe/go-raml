package commands

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen"
	"github.com/Jumpscale/go-raml/raml"
)

// CapnpCommand is executed to generate capnpm model from RAML specification
type CapnpCommand struct {
	Dir      string //target dir
	RAMLFile string //raml file
	Language string
	Package  string
}

//Execute generates a client from a RAML specification
func (command *CapnpCommand) Execute() error {
	var apiDef raml.APIDefinition

	command.Language = strings.ToLower(command.Language)
	if command.Language != "go" && command.Language != "nim" && command.Language != "python" {
		return fmt.Errorf("canpnp generator only support Go, Python, and Nim")
	}

	log.Debug("Generating capnp models")

	err := raml.ParseFile(command.RAMLFile, &apiDef)
	if err != nil {
		return err
	}
	return codegen.GenerateCapnp(&apiDef, command.Dir, command.Language, command.Package)
}
