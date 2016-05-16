package codegen

import (
	"github.com/Jumpscale/go-raml/raml"
	"path/filepath"
	"strings"
)

type goClient struct {
	clientDef
	libraries      map[string]*raml.Library
	RootImportPath string
}

// generate Go client files
func (gc goClient) generate(apiDef *raml.APIDefinition, dir string) error {
	pkgName := filepath.Base(gc.RootImportPath)
	// helper package
	gh := goramlHelper{
		packageName: pkgName,
		packageDir:  "",
	}
	if err := gh.generate(dir); err != nil {
		return err
	}

	// generate struct
	if err := generateStructs(apiDef.Types, dir, pkgName, langGo); err != nil {
		return err
	}

	// generate strucs from bodies
	if err := generateBodyStructs(apiDef, dir, pkgName, langGo); err != nil {
		return err
	}

	// libraries
	if err := generateLibraries(gc.libraries, dir); err != nil {
		return err
	}

	if err := gc.generateHelperFile(dir); err != nil {
		return err
	}
	return gc.generateClientFile(dir)
}

// generate Go client helper
func (gc *goClient) generateHelperFile(dir string) error {
	fileName := dir + "/client_utils.go"
	return generateFile(gc, clientHelperResourceTemplate, "client_helper_resource", fileName, false)
}

// generate Go client lib file
func (gc *goClient) generateClientFile(dir string) error {
	fileName := dir + "/client_" + strings.ToLower(gc.Name) + ".go"
	return generateFile(gc, clientResourceTemplate, "client_resource", fileName, false)
}

// LibImportPaths returns all imported lib
func (gc goClient) LibImportPaths() map[string]struct{} {
	ip := map[string]struct{}{}

	// methods
	for _, v := range gc.Methods {
		gm := v.(goClientMethod)
		for lib := range gm.libImported(globRootImportPath) {
			ip[lib] = struct{}{}
		}
	}
	return ip
}
