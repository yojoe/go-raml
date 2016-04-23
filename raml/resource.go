package raml

import (
	"strings"
)

// postProcess assign all properties that can't be obtained from RAML document
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
	r.inheritMethods(rt)
	return nil
}

// inherit methods inherits all methods based on it's resource type
func (r *Resource) inheritMethods(rt *ResourceType) {
	// only inherit methods that also defined in the resource => TODO make sure this assumption is true
	for _, m := range r.Methods {
		switch m.Name {
		case "GET":
			m.inherit(r, rt.Get, rt)
		case "POST":
			m.inherit(r, rt.Post, rt)
		case "PUT":
			m.inherit(r, rt.Put, rt)
		}
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

// get value of a resource type param
func (r *Resource) getResourceTypeParamValue(param string, rt *ResourceType) string {
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
