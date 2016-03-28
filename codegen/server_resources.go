package codegen

import (
	"sort"

	"github.com/Jumpscale/go-raml/raml"
)

// generate Server's Go representation of RAML resources
func generateServerResources(apiDef *raml.APIDefinition, directory, packageName, lang string) ([]resourceInterface, error) {
	var rds []resourceInterface

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
		rd := newResourceDef(apiDef, k, packageName)
		rd.IsServer = true
		switch lang {
		case langGo:
			gr := goResource{resourceDef: &rd}
			err = gr.generate(&r, k, directory)
		case langPython:
			pr := pythonResource{resourceDef: &rd}
			err = pr.generate(&r, k, directory)
		}
		if err != nil {
			return rds, err
		}
		rds = append(rds, rd)
	}
	return rds, nil
}
