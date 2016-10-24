package codegen

import (
	"path/filepath"
	"sort"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

// python client definition
type pythonClient struct {
	clientDef
	Services map[string]*ClientService
}

func newPythonClient(cd clientDef, apiDef *raml.APIDefinition) pythonClient {
	services := map[string]*ClientService{}
	for k, v := range apiDef.Resources {
		rd := resource.New(apiDef, commons.NormalizeURITitle(apiDef.Title), "")
		rd.GenerateMethods(&v, langPython, newServerMethod, newPythonClientMethod)
		services[k] = &ClientService{
			lang:         langPython,
			rootEndpoint: k,
			Methods:      rd.Methods,
		}
	}
	return pythonClient{
		clientDef: cd,
		Services:  services,
	}
}

// generate empty __init__.py without overwrite it
func generateEmptyInitPy(dir string) error {
	return commons.GenerateFile(nil, "./templates/init_py.tmpl", "init_py", filepath.Join(dir, "__init__.py"), false)
}

// generate python lib files
func (pc pythonClient) generate(dir string) error {
	// generate empty __init__.py
	if err := generateEmptyInitPy(dir); err != nil {
		return err
	}

	// generate helper
	if err := commons.GenerateFile(nil, "./templates/client_utils_python.tmpl", "client_utils_python", filepath.Join(dir, "client_utils.py"), false); err != nil {
		return err
	}

	if err := pc.generateServices(dir); err != nil {
		return err
	}
	// generate main client lib file
	return commons.GenerateFile(pc, "./templates/client_python.tmpl", "client_python", filepath.Join(dir, "client.py"), true)
}

func (pc pythonClient) generateServices(dir string) error {
	for _, s := range pc.Services {
		sort.Sort(resource.ByEndpoint(s.Methods))
		if err := commons.GenerateFile(s, "./templates/client_service_python.tmpl", "client_service_python", s.filename(dir), false); err != nil {
			return err
		}
	}
	return nil
}
