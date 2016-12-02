package codegen

import (
	"errors"
	"fmt"
	"path/filepath"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen/apidocs"
	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/golang"
	"github.com/Jumpscale/go-raml/codegen/nim"
	"github.com/Jumpscale/go-raml/codegen/python"
	"github.com/Jumpscale/go-raml/raml"
)

var (
	errInvalidLang = errors.New("invalid language")
)

// GenerateServer generates API server files
func GenerateServer(ramlFile, dir, packageName, lang, apiDocsDir, rootImportPath string, generateMain bool) error {
	apiDef := new(raml.APIDefinition)
	// parse the raml file
	ramlBytes, err := raml.ParseReadFile(ramlFile, apiDef)
	if err != nil {
		return err
	}

	// create directory if needed
	if err := commons.CheckCreateDir(dir); err != nil {
		return err
	}

	switch lang {
	case langGo:
		if rootImportPath == "" {
			return fmt.Errorf("invalid import path = empty")
		}
		gs := golang.NewServer(apiDef, packageName, apiDocsDir, rootImportPath, generateMain)
		err = gs.Generate(dir)
	case langPython:
		ps := python.NewServer(apiDef, apiDocsDir, generateMain)
		err = ps.Generate(dir)
	case langNim:
		ns := nim.NewServer(apiDef, apiDocsDir, dir)
		err = ns.Generate()
	default:
		return errInvalidLang
	}
	if err != nil {
		return err
	}

	if apiDocsDir == "" {
		return nil
	}

	if lang == langNim {
		apiDocsDir = "public/" + apiDocsDir
	}

	log.Infof("Generating API Docs to %v", apiDocsDir)

	return apidocs.Generate(apiDef, ramlFile, ramlBytes, filepath.Join(dir, apiDocsDir))
}
