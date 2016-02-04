package commands

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

const (
	resourceIfTemplate  = "./templates/server_resources_interface.tmpl"
	resourceAPITemplate = "./templates/server_resources_api.tmpl"
)

type resourceDef struct {
	Name     string
	Endpoint string
	Methods  []interfaceMethod
}

func newResourceDef(endpoint string) resourceDef {
	rd := resourceDef{Endpoint: endpoint}
	rd.Name = strings.Title(normalizeURI(endpoint))
	return rd
}

type interfaceMethod struct {
	*raml.Method
	MethodName string
	Endpoint   string
	Verb       string
}

func (im *interfaceMethod) setMethodName(rd *resourceDef, parentEndpoint, curEndpoint, methodName string) {
	name := normalizeURI(parentEndpoint) + normalizeURI(curEndpoint)
	im.MethodName = name[len(rd.Name):] + methodName
}

func (rd *resourceDef) addMethod(m *raml.Method, methodName, parentEndpoint, curEndpoint string) {
	if m == nil {
		return
	}
	im := interfaceMethod{
		Method:   m,
		Endpoint: parentEndpoint + curEndpoint,
		Verb:     strings.ToUpper(methodName),
	}
	im.setMethodName(rd, parentEndpoint, curEndpoint, methodName)
	rd.Methods = append(rd.Methods, im)
}

func (rd *resourceDef) generateMethods(r *raml.Resource, parentEndpoint, curEndpoint string) {
	rd.addMethod(r.Get, "Get", parentEndpoint, curEndpoint)
	rd.addMethod(r.Post, "Post", parentEndpoint, curEndpoint)
	rd.addMethod(r.Put, "Put", parentEndpoint, curEndpoint)
	rd.addMethod(r.Patch, "Patch", parentEndpoint, curEndpoint)
	rd.addMethod(r.Delete, "Delete", parentEndpoint, curEndpoint)
	for k, v := range r.Nested {
		rd.generateMethods(v, parentEndpoint+curEndpoint, k)
	}
}

func (rd *resourceDef) ifFileName(directory string) string {
	return directory + "/" + strings.ToLower(rd.Name) + "_if.go"
}

func (rd *resourceDef) apiFileName(directory string) string {
	return directory + "/" + strings.ToLower(rd.Name) + "_api.go"
}

func (rd *resourceDef) generateInterfaceFile(directory string) error {
	return generateFile(rd, resourceIfTemplate, "resource_if_template", rd.ifFileName(directory))
}

func (rd *resourceDef) generateAPIFile(directory string) error {
	return generateFile(rd, resourceAPITemplate, "resource_api_template", rd.apiFileName(directory))
}

// ServerResourceGen generate Server's Go representation of RAML resource
func ServerResourcesGen(rs map[string]raml.Resource, directory string) error {
	if err := checkCreateDir(directory); err != nil {
		return err
	}
	// create resource def
	for k, r := range rs {
		rd := newResourceDef(k)
		rd.generateMethods(&r, "", k)
		if err := rd.generateInterfaceFile(directory); err != nil {
			return err
		}
		if err := rd.generateAPIFile(directory); err != nil {
			return err
		}
	}
	return nil
}
