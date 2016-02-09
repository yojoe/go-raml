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
	Methods []interfaceMethod
}

func newClientDef(name string) clientDef {
	return clientDef{Name: name}
}

func (cd clientDef) generate(dir string) error {
	if err := cd.generateHelperFile(dir); err != nil {
		return err
	}
	return cd.generateClientFile(dir)
}

//GenerateClient generate client code
func GenerateClient(apiDef *raml.APIDefinition, dir string) error {
	//check create dir
	if err := checkCreateDir(dir); err != nil {
		return err
	}

	cd := newClientDef(normalizeURI(apiDef.Title))

	for k, v := range apiDef.Resources {
		rd := newResourceDef(normalizeURITitle(apiDef.Title))
		rd.generateMethods(&v, "", k)
		cd.Methods = append(cd.Methods, rd.Methods...)
	}

	if err := cd.generate(dir); err != nil {
		return err
	}
	return nil
}

func (cd *clientDef) generateHelperFile(dir string) error {
	return generateFile(cd, clientHelperResourceTemplate, "client_helper_resources", cd.clientHelperName(dir), false)
}

func (cd *clientDef) generateClientFile(dir string) error {
	return generateFile(cd, clientResourceTemplate, "client_resources2", cd.clientFileName(dir), false)
}

func (cd *clientDef) clientFileName(dir string) string {
	return dir + "/client_" + strings.ToLower(cd.Name) + ".go"
}

func (cd *clientDef) clientHelperName(dir string) string {
	return dir + "/client_utils.go"
}
