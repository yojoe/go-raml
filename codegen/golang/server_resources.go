package golang

import (
	"sort"

	"github.com/Jumpscale/go-raml/codegen/resource"
)

// generate Server's Go representation of RAML resources
func (s *Server) generateServerResources(dir string) ([]resource.Resource, error) {
	var rds []resource.Resource

	resources := s.apiDef.Resources

	// sort the keys, so we have resource sorted by keys.
	// the generated code actually don't need it to be sorted.
	// but test fixture need it
	var keys []string
	for k := range resources {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// create resource def
	var err error
	for _, endpoint := range keys {
		r := resources[endpoint]
		rd := resource.New(s.apiDef, &r, endpoint, true)
		gr := newGoResource(&rd, s.PackageName)
		err = gr.generate(&r, endpoint, dir, s.APIFilePerMethod, s.libsRootURLs)
		if err != nil {
			return rds, err
		}
		rds = append(rds, rd)
	}
	return rds, nil
}
