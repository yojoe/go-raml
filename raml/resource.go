package raml

import (
	"log"
	"strings"
)

// postProcess doing post processing of a resource after being constructed by the parser.
// some of the workds:
// - assign all properties that can't be obtained from RAML document
// - inherit from resource type
// - inherit from traits
func (r *Resource) postProcess(uri string, parent *Resource, resourceTypes []map[string]ResourceType) error {
	r.URI = uri
	r.Parent = parent

	r.setMethods()

	if err := r.inheritResourceType(resourceTypes); err != nil {
		return err
	}

	// process nested/child resources
	for k := range r.Nested {
		n := r.Nested[k]
		if n == nil {
			continue
		}
		if err := n.postProcess(k, r, resourceTypes); err != nil {
			return err
		}
		r.Nested[k] = n
	}
	return nil
}

// inherit from a resource type
func (r *Resource) inheritResourceType(resourceTypes []map[string]ResourceType) error {
	rt, err := r.getResourceType(resourceTypes)
	if err != nil || rt == nil {
		return err
	}
	r.Description = r.substituteParams(rt.Description)
	r.inheritMethods(rt)
	return nil
}

// inherit methods inherits all methods based on it's resource type
func (r *Resource) inheritMethods(rt *ResourceType) {
	// inherit all methods from resource type
	// if it doesn't have the methods, we create it
	for _, rtm := range rt.methods {
		m := r.MethodByName(rtm.Name)
		if m == nil {
			m = newMethod(rtm.Name)
			r.assignMethod(m, m.Name)
		}
		m.inherit(r, rtm, rt)
	}

	// inherit optional methods if only the resource also has the method
	for _, rtm := range rt.optionalMethods {
		m := r.MethodByName(rtm.Name)
		if m == nil {
			continue
		}
		m.inherit(r, rtm, rt)
	}

}

// get resource type from which this resource will inherit
func (r *Resource) getResourceType(resourceTypes []map[string]ResourceType) (*ResourceType, error) {
	if r.Type == nil || r.Type.Name == "" {
		return nil, nil
	}
	for _, rts := range resourceTypes {
		for k, rt := range rts {
			if k == r.Type.Name {
				return &rt, nil
			}
		}
	}
	return nil, nil
}

// set methods set all methods name
// and add it to Methods slice
func (r *Resource) setMethods() {
	if r.Get != nil {
		r.Get.Name = "GET"
		r.Methods = append(r.Methods, r.Get)
	}
	if r.Post != nil {
		r.Post.Name = "POST"
		r.Methods = append(r.Methods, r.Post)
	}
	if r.Put != nil {
		r.Put.Name = "PUT"
		r.Methods = append(r.Methods, r.Put)
	}
	if r.Patch != nil {
		r.Patch.Name = "PATCH"
		r.Methods = append(r.Methods, r.Patch)
	}
	if r.Head != nil {
		r.Head.Name = "HEAD"
		r.Methods = append(r.Methods, r.Head)
	}
	if r.Delete != nil {
		r.Delete.Name = "DELETE"
		r.Methods = append(r.Methods, r.Delete)
	}
}

// MethodByName return resource's method by it's name
func (r *Resource) MethodByName(name string) *Method {
	switch name {
	case "GET":
		return r.Get
	case "POST":
		return r.Post
	case "PUT":
		return r.Put
	case "PATCH":
		return r.Patch
	case "HEAD":
		return r.Head
	case "DELETE":
		return r.Delete
	default:
		return nil
	}
}

func (r *Resource) assignMethod(m *Method, name string) {
	switch name {
	case "GET":
		r.Get = m
	case "POST":
		r.Post = m
	case "PUT":
		r.Put = m
	case "PATCH":
		r.Patch = m
	case "HEAD":
		r.Head = m
	case "DELETE":
		r.Delete = m
	default:
		log.Fatalf("assignMethod fatal error, invalid method name:%v", name)
	}
}

// substituteParams substitute all params inside double chevron to the correct value
func (r *Resource) substituteParams(words string) string {
	removeParamBracket := func(param string) string {
		return param[2 : len(param)-2]
	}

	if words == "" {
		return words
	}

	// search params
	params := dcRe.FindAllString(words, -1)

	// substitute the params
	for _, p := range params {
		pVal := r.getResourceTypeParamValue(removeParamBracket(p))
		words = strings.Replace(words, p, pVal, -1)
	}
	return words
}

// get value of a resource type param
func (r *Resource) getResourceTypeParamValue(param string) string {
	// split between inflector and real param
	// real param and inflector is seperated by `|`
	cleanParam, inflector := func() (string, string) {
		arr := strings.Split(param, "|")
		if len(arr) != 2 {
			return param, ""
		}
		return strings.TrimSpace(arr[0]), strings.TrimSpace(arr[1])
	}()

	// get the value
	val := func() string {
		// check reserved params
		switch cleanParam {
		case "resourcePathName":
			return r.CleanURI()
		}
		return ""
	}()
	if inflector != "" {
		var ok bool
		val, ok = doInflect(val, inflector)
		if !ok {
			panic("invalid inflector " + inflector)
		}
	}
	return val
}

// CleanURI returns URI without `/`, `\`', `{`, and `}`
func (r *Resource) CleanURI() string {
	s := strings.Replace(r.URI, "/", "", -1)
	s = strings.Replace(s, `\`, "", -1)
	s = strings.Replace(s, "{", "", -1)
	s = strings.Replace(s, "}", "", -1)
	return strings.TrimSpace(s)
}
