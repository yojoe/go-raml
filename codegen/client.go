package codegen

import (
	"fmt"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

const (
	clientResourceTemplate       = "./templates/client_resource.tmpl"
	clientHelperResourceTemplate = "./templates/client_helper_resource.tmpl"
)

// API client definition
type clientDef struct {
	Name    string
	BaseURI string
	Methods []methodInterface
}

// create client definition from RAML API definition
func newClientDef(apiDef *raml.APIDefinition) clientDef {
	cd := clientDef{
		Name:    normalizeURI(apiDef.Title),
		BaseURI: apiDef.BaseURI,
	}
	if strings.Index(cd.BaseURI, "{version}") > 0 {
		cd.BaseURI = strings.Replace(cd.BaseURI, "{version}", apiDef.Version, -1)
	}
	return cd
}

// GenerateClient generates client library
func GenerateClient(apiDef *raml.APIDefinition, dir, packageName, lang, rootImportPath string) error {
	//check create dir
	if err := checkCreateDir(dir); err != nil {
		return err
	}

	// global variables
	globAPIDef = apiDef
	globRootImportPath = rootImportPath

	// creates base client struct
	cd := newClientDef(apiDef)

	for _, v := range apiDef.Resources {
		rd := newResourceDef(apiDef, normalizeURITitle(apiDef.Title), "main")
		rd.generateMethods(&v, lang)
		cd.Methods = append(cd.Methods, rd.Methods...)
	}

	switch lang {
	case langGo:
		// rootImportPath only needed if we use libraries
		if rootImportPath == "" && len(apiDef.Libraries) > 0 {
			return fmt.Errorf("--import-path can't be empty when we use libraries")
		}

		gc := goClient{
			clientDef:      cd,
			libraries:      apiDef.Libraries,
			PackageName:    packageName,
			RootImportPath: rootImportPath,
		}
		return gc.generate(apiDef, dir)
	case langPython:
		pc := pythonClient{clientDef: cd}
		return pc.generate(dir)
	}
	return errInvalidLang
}
