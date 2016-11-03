package codegen

import (
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/nim"
	"github.com/Jumpscale/go-raml/codegen/python"
	"github.com/Jumpscale/go-raml/raml"
)

// API client definition
type clientDef struct {
	Name     string
	BaseURI  string
	Services map[string]ClientService
}

// create client definition from RAML API definition
func newClientDef(apiDef *raml.APIDefinition) clientDef {
	cd := clientDef{
		Name:     commons.NormalizeURI(apiDef.Title),
		BaseURI:  apiDef.BaseURI,
		Services: map[string]ClientService{},
	}
	if strings.Index(cd.BaseURI, "{version}") > 0 {
		cd.BaseURI = strings.Replace(cd.BaseURI, "{version}", apiDef.Version, -1)
	}
	return cd
}

// GenerateClient generates client library
func GenerateClient(apiDef *raml.APIDefinition, dir, packageName, lang, rootImportPath string) error {
	//check create dir
	if err := commons.CheckCreateDir(dir); err != nil {
		return err
	}

	// global variables
	globAPIDef = apiDef

	// creates base client struct
	cd := newClientDef(apiDef)

	switch lang {
	case langGo:
		gc, err := newGoClient(cd, apiDef, packageName, rootImportPath)
		if err != nil {
			return err
		}
		return gc.generate(apiDef, dir)
	case langPython:
		pc := python.NewClient(apiDef)
		return pc.Generate(dir)
	case langNim:
		nc := nim.Client{
			APIDef: apiDef,
			Dir:    dir,
		}
		return nc.Generate()
	}
	return errInvalidLang
}
