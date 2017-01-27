package python

import (
	"sort"

	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

func getServerResourcesDefs(apiDef *raml.APIDefinition) []resource.Resource {
	var rds []resource.Resource

	rs := apiDef.Resources

	// sort the keys, so we have resource sorted by keys.
	// the generated code actually don't need it to be sorted.
	// but test fixture need it
	var keys []string
	for k := range rs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// create resource def
	for _, k := range keys {
		rd := resource.New(apiDef, k, "")
		rd.IsServer = true
		rds = append(rds, rd)
	}
	return rds
}
