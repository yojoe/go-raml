package golang

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

const (
	resourceIfTemplate  = "./templates/golang/server_resources_interface.tmpl" // resource interface template
	resourceAPITemplate = "./templates/golang/server_resources_api.tmpl"       // resource API template
)

type goResource struct {
	*resource.Resource
}

// generate interface file of a resource
func (gr *goResource) generateInterfaceFile(directory string) error {
	gr.SortMethods()
	filename := directory + "/" + strings.ToLower(gr.Name) + "_if.go"
	return commons.GenerateFile(gr, resourceIfTemplate, "resource_if_template", filename, true)
}

// generate API file of a resource
func (gr *goResource) generateAPIFile(dir string) error {
	filename := filepath.Join(dir, strings.ToLower(gr.Name)+"_api.go")
	return commons.GenerateFile(gr, resourceAPITemplate, "resource_api_template", filename, false)
}

// generate API implementation in one file per method mode
func (gr *goResource) generateAPIImplementations(dir string) error {
	// generate the main API impl file, which only contains struct
	mainCtx := map[string]interface{}{
		"PackageName": gr.PackageName,
		"Name":        gr.Name,
		"Endpoint":    gr.Endpoint,
	}

	mainFile := filepath.Join(dir, strings.ToLower(gr.Name)+"_api")
	if err := commons.GenerateFile(mainCtx, "./templates/golang/server_resource_api_main_go.tmpl",
		"server_resource_api_main_go", mainFile+".go", false); err != nil {
		return err
	}

	// generate per methods file
	for _, mi := range gr.Methods {
		sm := mi.(serverMethod)
		ctx := map[string]interface{}{
			"Method":      sm,
			"APIName":     gr.Name,
			"PackageName": gr.PackageName,
		}
		filename := mainFile + "_" + sm.MethodName + ".go"
		if err := commons.GenerateFile(ctx, "./templates/golang/server_resource_api_impl_go.tmpl",
			"server_resource_api_impl_go", filename, false); err != nil {
			return err
		}
	}
	return nil
}

// generate Go representation of server's resource.
// A resource have two kind of files:
// - interface file
//		contains API interface and routing code
//		always regenerated
// - API implementation
//		implementation of the API interface.
//		Don't generate if the file already exist
func (gr *goResource) generate(r *raml.Resource, URI, dir string,
	apiFilePerMethod bool, libRootURLs []string) error {
	gr.GenerateMethods(r, "go", newServerMethod, newGoClientMethod)
	if err := gr.generateInterfaceFile(dir); err != nil {
		return err
	}
	if !apiFilePerMethod {
		return gr.generateAPIFile(dir)
	}
	return gr.generateAPIImplementations(dir)
}

// InterfaceImportPaths returns all packages imported by
// this resource interface file
func (gr goResource) InterfaceImportPaths() []string {
	ip := map[string]struct{}{
		"net/http":               struct{}{},
		"github.com/gorilla/mux": struct{}{},
	}

	for _, v := range gr.Methods {
		gm := v.(serverMethod)

		// if has middleware, we need to import middleware helper library
		if len(gm.Middlewares) > 0 {
			ip["github.com/justinas/alice"] = struct{}{}
		}
		for _, sb := range gm.SecuredBy {
			if lib := libImportPath(globRootImportPath, sb.Name, globLibRootURLs); lib != "" {
				ip[lib] = struct{}{}
			}
		}
	}
	// return sorted array for predictable order
	// we need it for unit test to always return same order
	return sortImportPaths(ip)
}

// APIImportPaths returns all packages that need to be imported
// by the API implementation
func (gr goResource) APILibImportPaths() []string {
	ip := map[string]struct{}{}

	// methods
	for _, v := range gr.Methods {
		gm := v.(serverMethod)
		for _, v := range gm.Imports() {
			ip[v] = struct{}{}
		}
	}

	// return sorted array for predictable order
	// we need it for unit test to always return same order
	return sortImportPaths(ip)
}

func sortImportPaths(ip map[string]struct{}) []string {
	libs := []string{}
	for k := range ip {
		libs = append(libs, k)
	}
	sort.Strings(libs)
	return libs
}
