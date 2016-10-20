package nim

import (
	"sort"

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
