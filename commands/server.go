package commands

import (
	"github.com/Jumpscale/go-raml/codegen"

	log "github.com/Sirupsen/logrus"
)

// ServerCommand is executed to generate a go server from a RAML specification
type ServerCommand struct {
	Language         string // target language
	Kind             string
	Dir              string //target dir
	RamlFile         string //raml file
	PackageName      string //package name in the generated go source files
	NoMainGeneration bool   //do not generate a main.go file
	ImportPath       string // root import path of the code, such as : github.com/jumpscale/restapi
	NoAPIDocs        bool   // do not generate API Docs in /apidocs/ endpoint
}

// Execute generates a Go server from an RAML specification
func (command *ServerCommand) Execute() error {
	var apiDocsDir string

	log.Infof("Generating a %v server", command.Language)

	if !command.NoAPIDocs {
		apiDocsDir = "apidocs"
	}

	return codegen.GenerateServer(command.RamlFile, command.Kind, command.Dir, command.PackageName,
		command.Language, apiDocsDir, command.ImportPath, !command.NoMainGeneration)
}
