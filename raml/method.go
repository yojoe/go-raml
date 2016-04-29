package raml

import (
	"fmt"
	"strings"
)

func newMethod(name string) *Method {
	return &Method{
		Name: name,
	}
}

// inherit from resource type
// fields need to be inherited:
// - description
// - response
func (m *Method) inheritFromResourceType(r *Resource, rtm *ResourceTypeMethod, rt *ResourceType) {
	if rtm == nil {
		return
	}
	dicts := initResourceTypeDicts(r, r.Type.Parameters)

	// inherit description
	m.Description = substituteParams(m.Description, rtm.Description, dicts)

	// inherit bodies
	m.Bodies.inherit(rtm.Bodies, dicts)

	// inherit headers
	m.inheritHeaders(rtm.Headers, dicts)

	// inherit query params
	m.inheritQueryParams(rtm.QueryParameters, dicts)

	// inherit response
	m.inheritResponses(rtm.Responses, dicts)

	// inherit protocols
	m.inheritProtocols(rtm.Protocols)
}

// inherit from all traits, inherited traits are:
// - resource level trait
// - method trait
func (m *Method) inheritFromAllTraits(r *Resource) error {
	for _, tDef := range append(r.Is, m.Is...) {
		// acquire traits object
		t, ok := traitsMap[tDef.Name]
		if !ok {
			return fmt.Errorf("invalid traits name:%v", tDef.Name)
		}

		if err := m.inheritFromATrait(r, &t, tDef.Parameters); err != nil {
			return err
		}
	}
	return nil
}

// inherit from a trait
// dicts is map of trait parameters values
func (m *Method) inheritFromATrait(r *Resource, t *Trait, dicts map[string]interface{}) error {
	dicts = initTraitDicts(r, m, dicts)

	m.Description = substituteParams(m.Description, t.Description, dicts)

	m.Bodies.inherit(t.Bodies, dicts)

	m.inheritHeaders(t.Headers, dicts)

	m.inheritResponses(t.Responses, dicts)

	m.inheritQueryParams(t.QueryParameters, dicts)

	m.inheritProtocols(t.Protocols)

	// optional bodies
	// optional headers
	// optional responses
	// optional query parameters
	return nil
}

// inheritHeaders inherit method's headers from parent headers.
// parent headers could be from resource type or a trait
func (m *Method) inheritHeaders(parents map[HTTPHeader]Header, dicts map[string]interface{}) {
	if len(m.Headers) == 0 {
		m.Headers = map[HTTPHeader]Header{}
	}

	for name, parent := range parents {
		h, ok := m.Headers[name]
		if !ok {
			if optionalTraitProperty(string(name)) { // don't inherit optional property if not exist
				continue
			}
			h = Header{}
		}
		parent.Name = string(name)
		np := NamedParameter(h)
		np.inherit(NamedParameter(parent), dicts)
		m.Headers[name] = Header(np)
	}
}

// inheritQueryParams inherit method's query params from parent query params.
// parent query params could be from resource type or a trait
func (m *Method) inheritQueryParams(parents map[string]NamedParameter, dicts map[string]interface{}) {
	if len(m.QueryParameters) == 0 {
		m.QueryParameters = map[string]NamedParameter{}
	}
	for name, parent := range parents {
		qp, ok := m.QueryParameters[name]
		if !ok {
			if optionalTraitProperty(string(name)) { // don't inherit optional property if not exist
				continue
			}
			qp = NamedParameter{Name: name}
		}
		parent.Name = name // parent name is not initialized by the parser
		qp.inherit(parent, dicts)
		m.QueryParameters[qp.Name] = qp
	}

}

// inheritProtocols inherit method's protocols from parent protocols
// parent protocols could be from resource type or a trait
func (m *Method) inheritProtocols(parent []string) {
	for _, p := range parent {
		m.Protocols = appendStrNotExist(p, m.Protocols)
	}
}

// inheritResponses inherit method's responses from parent responses
// parent responses could be from resource type or a trait
func (m *Method) inheritResponses(parent map[HTTPCode]Response, dicts map[string]interface{}) {
	if len(m.Responses) == 0 { // allocate if needed
		m.Responses = map[HTTPCode]Response{}
	}
	for code, rParent := range parent {
		resp, ok := m.Responses[code]
		if !ok {
			resp = Response{}
		}
		resp.inherit(rParent, dicts)
		m.Responses[code] = resp
	}

}

// inherit from parent response
func (resp *Response) inherit(parent Response, dicts map[string]interface{}) {
	resp.Bodies.Type = substituteParams(resp.Bodies.Type, parent.Bodies.Type, dicts)
}

// inherit inherits bodies properties from a parent bodies
// parent object could be from trait or response type
func (b *Bodies) inherit(parent Bodies, dicts map[string]interface{}) {
	b.DefaultSchema = substituteParams(b.DefaultSchema, parent.DefaultSchema, dicts)
	b.DefaultDescription = substituteParams(b.DefaultDescription, parent.DefaultDescription, dicts)
	b.DefaultExample = substituteParams(b.DefaultExample, parent.DefaultExample, dicts)

	b.Type = substituteParams(b.Type, parent.Type, dicts)

	// request body
	if parent.ApplicationJson != nil {
		if b.ApplicationJson == nil { // allocate if needed
			b.ApplicationJson = &BodiesProperty{Properties: map[string]interface{}{}}
		}

		b.ApplicationJson.Type = substituteParams(b.ApplicationJson.Type, parent.ApplicationJson.Type, dicts)

		for k, p := range parent.ApplicationJson.Properties {
			if _, ok := b.ApplicationJson.Properties[k]; !ok {

				// handle optional properties as described in
				// https://github.com/raml-org/raml-spec/blob/raml-10/versions/raml-10/raml-10.md#optional-properties
				switch {
				case strings.HasSuffix(k, `\?`): // if ended with `\?` we make it optional property
					k = k[:len(k)-2] + "?"
				case strings.HasSuffix(k, "?"): // if only ended with `?`, we can ignore it
					continue
				}
				b.ApplicationJson.Properties[k] = p
			}
		}
	}
}
