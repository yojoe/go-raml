package nim

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	cr "github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

type resource struct {
	cr.Resource
}

func newResource(name string, apiDef *raml.APIDefinition) resource {
	rd := cr.New(apiDef, name, "")
	rd.IsServer = true
	r := resource{
		Resource: rd,
	}
	res := apiDef.Resources[name]
	r.GenerateMethods(&res, "nim", newServerMethod, nil)
	return r
}

func (r *resource) Imports() []string {
	ip := map[string]struct{}{}

	for _, mi := range r.Methods {
		m := mi.(method)
		if m.ReqBody != "" {
			ip[m.ReqBody] = struct{}{}
		}
		if m.RespBody != "" {
			ip[m.RespBody] = struct{}{}
		}
	}
	// filter it
	imports := []string{}
	for k := range ip {
		if !inGeneratedObjs(k) {
			continue
		}
		imports = append(imports, k)
	}
	sort.Strings(imports)
	return imports
}

func (r *resource) generate(dir string) error {
	filename := filepath.Join(dir, r.apiName()+".nim")
	return commons.GenerateFile(r, "./templates/server_resources_api_nim.tmpl", "server_resources_api_nim", filename, true)
}

func (r *resource) apiName() string {
	return strings.ToLower(r.Name) + "_api"
}

func getAllResources(apiDef *raml.APIDefinition) []resource {
	rs := []resource{}

	// sort the keys, so we have resource sorted by keys.
	// the generated code actually don't need it to be sorted.
	// but test fixture need it
	var keys []string
	for k := range apiDef.Resources {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		rs = append(rs, newResource(k, apiDef))
	}
	return rs
}
