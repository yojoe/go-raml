package commands

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

const (
	clientResourceTemplate       = "./templates/client_resource.tmpl"
	clientHelperResourceTemplate = "./templates/client_helper_resource.tmpl"
)

type clientDef struct {
	Name    string
	BaseURI string
	Methods []interfaceMethod
}

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

func (cd clientDef) generate(apiDef *raml.APIDefinition, dir, lang string) error {
	if lang == "python" {
		return cd.generatePython(dir)
	}
	return cd.generateGo(apiDef, dir)
}

func (cd clientDef) generateGo(apiDef *raml.APIDefinition, dir string) error {
	//generate struct
	if err := GenerateStruct(apiDef, dir, "client"); err != nil {
		return err
	}

	//generate body struct
	if err := GenerateBodyStruct(apiDef, dir, "client"); err != nil {
		return err
	}

	if err := cd.generateHelperFile(dir); err != nil {
		return err
	}
	return cd.generateClientFile(dir)
}

//GenerateClient generate client code
func GenerateClient(apiDef *raml.APIDefinition, dir, lang string) error {
	//check create dir
	if err := checkCreateDir(dir); err != nil {
		return err
	}

	cd := newClientDef(apiDef)

	for k, v := range apiDef.Resources {
		rd := newResourceDef(normalizeURITitle(apiDef.Title))
		rd.generateMethods(&v, "", k)
		cd.Methods = append(cd.Methods, rd.Methods...)
	}

	if err := cd.generate(apiDef, dir, lang); err != nil {
		return err
	}
	return nil
}

func (cd *clientDef) generateHelperFile(dir string) error {
	return generateFile(cd, clientHelperResourceTemplate, "client_helper_resources", cd.clientHelperName(dir), false)
}

func (cd *clientDef) generateClientFile(dir string) error {
	return generateFile(cd, clientResourceTemplate, "client_resource", cd.clientFileName(dir), false)
}

func (cd *clientDef) clientFileName(dir string) string {
	return dir + "/client_" + strings.ToLower(cd.Name) + ".go"
}

func (cd *clientDef) clientHelperName(dir string) string {
	return dir + "/client_utils.go"
}
