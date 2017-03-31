package types

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

const (
	HTTPRequest = iota
	HTTPResponse
)

type Endpoint struct {
	Addr     string // complete endpoint address
	Method   *raml.Method
	Resource *raml.Resource
	Verb     string
}

func (ep Endpoint) ResourceName() string {
	name := ep.Addr
	splt := strings.Split(name, "/")
	if len(splt) > 0 {
		name = splt[0]
	}
	return strings.TrimSuffix(strings.TrimPrefix(name, "/"), "/")
}

func getAllEndpoints(apiDef *raml.APIDefinition) map[string][]Endpoint {
	endpoints := map[string][]Endpoint{}
	for _, r := range apiDef.Resources {
		getEndpointsOfResource("", &r, endpoints)
	}
	return endpoints
}

func getEndpointsOfResource(parentPath string, r *raml.Resource, endpoints map[string][]Endpoint) {
	var methods = []struct {
		Name   string
		Method *raml.Method
	}{
		{Name: "Get", Method: r.Get},
		{"Post", r.Post},
		{"Head", r.Head},
		{"Put", r.Put},
		{"Delete", r.Delete},
		{"Patch", r.Patch},
		{"Options", r.Options},
	}

	for _, m := range methods {
		if m.Method == nil {
			continue
		}
		endp := Endpoint{
			Addr:     parentPath + r.URI,
			Method:   m.Method,
			Verb:     strings.ToUpper(m.Name),
			Resource: r,
		}
		if _, exists := endpoints[endp.Addr]; exists {
			endpoints[endp.Addr] = append(endpoints[endp.Addr], endp)
		} else {
			endpoints[endp.Addr] = []Endpoint{endp}
		}
	}

	for _, v := range r.Nested {
		getEndpointsOfResource(parentPath+r.URI, v, endpoints)
	}
}
