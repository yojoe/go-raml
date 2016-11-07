package python

import (
	"sort"

	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
)

// generate Server's Go representation of RAML resources
func generateServerResources(apiDef *raml.APIDefinition, directory string) ([]resource.ResourceInterface, error) {
	var rds []resource.ResourceInterface

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
	var err error
	for _, k := range keys {
		r := rs[k]
		rd := resource.New(apiDef, k, "")
		rd.IsServer = true
		pr := pythonResource{Resource: &rd}
		err = pr.generate(&r, k, directory)
		if err != nil {
			return rds, err
		}
		rds = append(rds, rd)
	}
	return rds, nil
}
