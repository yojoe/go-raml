package codegen

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

type goClient struct {
	clientDef
	libraries      map[string]*raml.Library
	PackageName    string
	RootImportPath string
	Services       map[string]*ClientService
}

// generate Go client files
func (gc goClient) generate(apiDef *raml.APIDefinition, dir string) error {
	// helper package
	gh := goramlHelper{
		packageName: gc.PackageName,
		packageDir:  "",
	}
	if err := gh.generate(dir); err != nil {
		return err
	}

	// generate struct
	if err := generateStructs(apiDef.Types, dir, gc.PackageName, langGo); err != nil {
		return err
	}

	// generate strucs from bodies
	if err := generateBodyStructs(apiDef, dir, gc.PackageName, langGo); err != nil {
		return err
	}

	// libraries
	if err := generateLibraries(gc.libraries, dir); err != nil {
		return err
	}

	if err := gc.generateHelperFile(dir); err != nil {
		return err
	}

	if err := gc.generateServices(dir); err != nil {
		return err
	}
	return gc.generateClientFile(dir)
}

// generate Go client helper
func (gc *goClient) generateHelperFile(dir string) error {
	fileName := filepath.Join(dir, "/client_utils.go")
	return generateFile(gc, "./templates/client_utils_go.tmpl", "client_utils_go", fileName, false)
}

func (gc *goClient) generateServices(dir string) error {
	for _, s := range gc.Services {
		sort.Sort(byEndpoint(s.Methods))
		if err := generateFile(s, "./templates/client_service_go.tmpl", "client_service_go", s.filename(dir), false); err != nil {
			return err
		}
	}
	return nil
}

// generate Go client lib file
func (gc *goClient) generateClientFile(dir string) error {
	fileName := filepath.Join(dir, "/client_"+strings.ToLower(gc.Name)+".go")
	return generateFile(gc, "./templates/client_go.tmpl", "client_go", fileName, false)
}
