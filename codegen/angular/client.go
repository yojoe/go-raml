package angular

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

var (
	globAPIDef *raml.APIDefinition
)

// Client represents a angular client
type Client struct {
	Name     string
	APIDef   *raml.APIDefinition
	BaseURI  string
	Services map[string]*service
	Template clientTemplate
}

// NewClient creates a angular Client
func NewClient(apiDef *raml.APIDefinition) Client {
	services := map[string]*service{}
	for k, v := range apiDef.Resources {
		rd := resource.New(apiDef, commons.NormalizeURITitle(apiDef.Title), "")
		rd.GenerateMethods(&v, "angular", newServerMethod, newClientMethod)
		services[k] = &service{
			rootEndpoint: k,
			Methods:      rd.Methods,
		}
	}

	c := Client{
		Name:     commons.NormalizeURI(apiDef.Title),
		APIDef:   apiDef,
		BaseURI:  apiDef.BaseURI,
		Services: services,
	}
	if strings.Index(c.BaseURI, "{version}") > 0 {
		c.BaseURI = strings.Replace(c.BaseURI, "{version}", apiDef.Version, -1)
	}
	c.initTemplates()
	return c
}

// Generate generates angular client library files
func (c Client) Generate(dir string) error {
	globAPIDef = c.APIDef
	// generate main heirarchy
	if err := commons.RestoreDir(c.Template.staticDir, dir, true); err != nil {
		return err
	}

	appdir := filepath.Join(dir, "src/app")
	if err := c.generateServices(appdir); err != nil {
		return err
	}

	// generate main client lib file
	return commons.GenerateFile(c, c.Template.mainFile, c.Template.mainName, filepath.Join(appdir, "app.module.ts"), true)
}

func (c Client) generateServices(dir string) error {
	for _, s := range c.Services {
		sort.Sort(resource.ByEndpoint(s.Methods))
		if err := commons.GenerateFile(s, c.Template.serviceFile, c.Template.serviceName, s.filename(dir), false); err != nil {
			return err
		}
	}
	return nil
}
