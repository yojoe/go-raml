package golang

import (
	"path/filepath"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

// global variables
// it is needed for libraries support
var (
	// root import path
	globRootImportPath string

	// global value of API definition
	globAPIDef *raml.APIDefinition
)

// Server represents a Go server
type Server struct {
	apiDef         *raml.APIDefinition
	Title          string
	ResourcesDef   []resource.ResourceInterface
	PackageName    string // Name of the package this server resides in
	APIDocsDir     string // apidocs directory. apidocs won't be generated if it is empty
	withMain       bool
	RootImportPath string
}

func NewServer(apiDef *raml.APIDefinition, packageName, apiDocsDir, rootImportPath string, withMain bool) Server {
	// global variables
	globAPIDef = apiDef
	globRootImportPath = rootImportPath

	return Server{
		apiDef:         apiDef,
		Title:          apiDef.Title,
		PackageName:    packageName,
		APIDocsDir:     apiDocsDir,
		withMain:       withMain,
		RootImportPath: rootImportPath,
	}
}

// generate all Go server files
func (gs Server) Generate(dir string) error {
	// helper package
	gh := goramlHelper{
		rootImportPath: gs.RootImportPath,
		packageName:    "goraml",
		packageDir:     "goraml",
	}
	if err := gh.generate(dir); err != nil {
		return err
	}

	// generate all Type structs
	if err := generateStructs(gs.apiDef.Types, dir, gs.PackageName, langGo); err != nil {
		return err
	}

	// generate all request & response body
	if err := generateBodyStructs(gs.apiDef, dir, gs.PackageName, langGo); err != nil {
		return err
	}

	// security scheme
	if err := generateSecurity(gs.apiDef.SecuritySchemes, dir, gs.PackageName); err != nil {
		log.Errorf("failed to generate security scheme:%v", err)
		return err
	}

	// genereate resources
	rds, err := generateServerResources(gs.apiDef, dir, gs.PackageName)
	if err != nil {
		return err
	}
	gs.ResourcesDef = rds

	// libraries
	if err := generateLibraries(gs.apiDef.Libraries, dir); err != nil {
		return err
	}

	// generate main
	if gs.withMain {
		// HTML front page
		if err := commons.GenerateFile(gs, "./templates/index.html.tmpl", "index.html", filepath.Join(dir, "index.html"), false); err != nil {
			return err
		}
		// main file
		return commons.GenerateFile(gs, "./templates/server_main_go.tmpl", "server_main_go", filepath.Join(dir, "main.go"), true)
	}

	return nil
}
