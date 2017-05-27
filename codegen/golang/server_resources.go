package golang

import (
	"sort"

	"github.com/Jumpscale/go-raml/codegen/resource"
)

// generate Server's Go representation of RAML resources
func (s *Server) generateServerResources(dir string) ([]resource.ResourceInterface, error) {
	var rds []resource.ResourceInterface

	rs := s.apiDef.Resources

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
		rd := resource.New(s.apiDef, k, s.PackageName)
		rd.IsServer = true
		gr := goResource{Resource: &rd}
		err = gr.generate(&r, k, dir, s.APIFilePerMethod, s.libsRootURLs)
		if err != nil {
			return rds, err
		}
		rds = append(rds, rd)
	}
	return rds, nil
}
