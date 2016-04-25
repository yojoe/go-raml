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
	rt.setMethods()
	rt.setOptionalMethods()

	// TODO : inherit from other resource type

	// TODO : apply traits
}

// set methods set all methods name
// and add it to methods slice
func (rt *ResourceType) setMethods() {
	if rt.Get != nil {
		rt.Get.Name = "GET"
		rt.methods = append(rt.methods, rt.Get)
	}
	if rt.Post != nil {
		rt.Post.Name = "POST"
		rt.methods = append(rt.methods, rt.Post)
	}
	if rt.Put != nil {
		rt.Put.Name = "PUT"
		rt.methods = append(rt.methods, rt.Put)
	}
	if rt.Patch != nil {
		rt.Patch.Name = "PATCH"
		rt.methods = append(rt.methods, rt.Patch)
	}
	if rt.Head != nil {
		rt.Head.Name = "HEAD"
		rt.methods = append(rt.methods, rt.Head)
	}
	if rt.Delete != nil {
		rt.Delete.Name = "DELETE"
		rt.methods = append(rt.methods, rt.Delete)
	}
}

// setOptionalMethods set name of all optional methods
// and add it to optionalMethods slice
func (rt *ResourceType) setOptionalMethods() {
	if rt.OptionalGet != nil {
		rt.OptionalGet.Name = "GET"
		rt.optionalMethods = append(rt.optionalMethods, rt.OptionalGet)
	}
	if rt.OptionalPost != nil {
		rt.OptionalPost.Name = "POST"
		rt.optionalMethods = append(rt.optionalMethods, rt.OptionalPost)
	}
	if rt.OptionalPut != nil {
		rt.OptionalPut.Name = "PUT"
		rt.optionalMethods = append(rt.optionalMethods, rt.OptionalPut)
	}
	if rt.OptionalPatch != nil {
		rt.OptionalPatch.Name = "PATCH"
		rt.optionalMethods = append(rt.optionalMethods, rt.OptionalPatch)
	}
	if rt.OptionalHead != nil {
		rt.OptionalHead.Name = "HEAD"
		rt.optionalMethods = append(rt.optionalMethods, rt.OptionalHead)
	}
	if rt.OptionalDelete != nil {
		rt.OptionalDelete.Name = "DELETE"
		rt.optionalMethods = append(rt.optionalMethods, rt.OptionalDelete)
	}
}
