package raml

import (
	"regexp"
)

var (
	dcRe = regexp.MustCompile(`\<<([^]]+)\>>`) // double chevron regex
)

// postProcess doing post processing of a resource type after being constructed
// by the .raml parser, some of the works:
// - assign all properties that can't be obtained from RAML document
// - inherit from other resource type
// - apply traits
func (rt *ResourceType) postProcess(name string) {
	rt.Name = name

	// TODO : inherit from other resource type

	// TODO : apply traits
}
