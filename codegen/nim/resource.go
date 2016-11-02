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

func newResource(name string, apiDef *raml.APIDefinition, isServer bool) resource {
	rd := cr.New(apiDef, name, "")
	rd.IsServer = isServer
	r := resource{
		Resource: rd,
	}
	res := apiDef.Resources[name]
	r.GenerateMethods(&res, "nim", newServerMethod, newMethod)
	return r
}

// get array of all imported modules
func (r *resource) Imports() []string {
	ip := map[string]struct{}{}

	for _, mi := range r.Methods {
		m := mi.(method)
		if m.ReqBody != "" && objectRegistered(m.ReqBody) {
			ip[m.ReqBody] = struct{}{}
		}
		if m.RespBody != "" && objectRegistered(m.RespBody) {
			ip[m.RespBody] = struct{}{}
		}
	}
	return commons.MapToSortedStrings(ip)
}

func (r *resource) generate(dir string) error {
	filename := filepath.Join(dir, r.apiName()+".nim")
	return commons.GenerateFile(r, "./templates/server_resources_api_nim.tmpl", "server_resources_api_nim", filename, true)
}

func (r *resource) apiName() string {
	return strings.ToLower(r.Name) + "_api"
}

func getAllResources(apiDef *raml.APIDefinition, isServer bool) []resource {
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
		rs = append(rs, newResource(k, apiDef, isServer))
	}
	return rs
}
