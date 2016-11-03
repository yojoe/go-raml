package golang

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

const (
	langGo = "go"
)

type goClient struct {
	apiDef         *raml.APIDefinition
	Name           string
	BaseURI        string
	libraries      map[string]*raml.Library
	PackageName    string
	RootImportPath string
	Services       map[string]*ClientService
}

func NewClient(apiDef *raml.APIDefinition, packageName, rootImportPath string) (goClient, error) {
	// rootImportPath only needed if we use libraries
	if rootImportPath == "" && len(apiDef.Libraries) > 0 {
		return goClient{}, fmt.Errorf("--import-path can't be empty when we use libraries")
	}

	globRootImportPath = rootImportPath
	globAPIDef = apiDef

	services := map[string]*ClientService{}
	for k, v := range apiDef.Resources {
		rd := resource.New(apiDef, commons.NormalizeURITitle(apiDef.Title), packageName)
		rd.GenerateMethods(&v, langGo, newServerMethod, newGoClientMethod)
		services[k] = &ClientService{
			lang:         langGo,
			rootEndpoint: k,
			PackageName:  packageName,
			Methods:      rd.Methods,
		}
	}
	client := goClient{
		apiDef:         apiDef,
		Name:           commons.NormalizeURI(apiDef.Title),
		BaseURI:        apiDef.BaseURI,
		libraries:      apiDef.Libraries,
		PackageName:    packageName,
		RootImportPath: rootImportPath,
		Services:       services,
	}

	if strings.Index(client.BaseURI, "{version}") > 0 {
		client.BaseURI = strings.Replace(client.BaseURI, "{version}", apiDef.Version, -1)
	}
	return client, nil
}

// generate Go client files
func (gc goClient) Generate(dir string) error {
	// helper package
	gh := goramlHelper{
		packageName: gc.PackageName,
		packageDir:  "",
	}
	if err := gh.generate(dir); err != nil {
		return err
	}

	// generate struct
	if err := generateStructs(gc.apiDef.Types, dir, gc.PackageName, langGo); err != nil {
		return err
	}

	// generate strucs from bodies
	if err := generateBodyStructs(gc.apiDef, dir, gc.PackageName, langGo); err != nil {
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
	return commons.GenerateFile(gc, "./templates/client_utils_go.tmpl", "client_utils_go", fileName, false)
}

func (gc *goClient) generateServices(dir string) error {
	for _, s := range gc.Services {
		sort.Sort(resource.ByEndpoint(s.Methods))
		if err := commons.GenerateFile(s, "./templates/client_service_go.tmpl", "client_service_go", s.filename(dir), false); err != nil {
			return err
		}
	}
	return nil
}

// generate Go client lib file
func (gc *goClient) generateClientFile(dir string) error {
	fileName := filepath.Join(dir, "/client_"+strings.ToLower(gc.Name)+".go")
	return commons.GenerateFile(gc, "./templates/client_go.tmpl", "client_go", fileName, false)
}
