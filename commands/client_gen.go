package commands

import (
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
	Methods []interfaceMethod
}

// create client definition from RAML API definition
func newClientDef(apiDef *raml.APIDefinition) clientDef {
	cd := clientDef{
		Name:    normalizeURI(apiDef.Title),
		BaseURI: apiDef.BaseUri,
	}
	if strings.Index(cd.BaseURI, "{version}") > 0 {
		cd.BaseURI = strings.Replace(cd.BaseURI, "{version}", apiDef.Version, -1)
	}
	return cd
}

// generate client files
func (cd clientDef) generate(apiDef *raml.APIDefinition, dir, lang string) error {
	if lang == "python" {
		return cd.generatePython(dir)
	}
	return cd.generateGo(apiDef, dir)
}

// generate Go client files
func (cd clientDef) generateGo(apiDef *raml.APIDefinition, dir string) error {
	// generate struct
	if err := generateStructs(apiDef, dir, "client"); err != nil {
		return err
	}

	// generate strucs from bodies
	if err := generateBodyStructs(apiDef, dir, "client"); err != nil {
		return err
	}

	if err := cd.generateHelperFile(dir); err != nil {
		return err
	}
	return cd.generateClientFile(dir)
}

// generate client library
func generateClient(apiDef *raml.APIDefinition, dir, lang string) error {
	//check create dir
	if err := checkCreateDir(dir); err != nil {
		return err
	}

	cd := newClientDef(apiDef)

	for k, v := range apiDef.Resources {
		rd := newResourceDef(normalizeURITitle(apiDef.Title), "main")
		rd.generateMethods(&v, "", k, lang)
		cd.Methods = append(cd.Methods, rd.Methods...)
	}

	if err := cd.generate(apiDef, dir, lang); err != nil {
		return err
	}
	return nil
}

// generate client helper
func (cd *clientDef) generateHelperFile(dir string) error {
	return generateFile(cd, clientHelperResourceTemplate, "client_helper_resources", cd.clientHelperName(dir), false)
}

// generate main client lib file
func (cd *clientDef) generateClientFile(dir string) error {
	return generateFile(cd, clientResourceTemplate, "client_resource", cd.clientFileName(dir), false)
}

func (cd *clientDef) clientFileName(dir string) string {
	return dir + "/client_" + strings.ToLower(cd.Name) + ".go"
}

func (cd *clientDef) clientHelperName(dir string) string {
	return dir + "/client_utils.go"
}
