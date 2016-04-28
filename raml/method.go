package raml

import (
	"fmt"
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
func (m *Method) inheritResourceType(r *Resource, rtm *ResourceTypeMethod, rt *ResourceType) {
	if rtm == nil {
		return
	}

	// inherit description
	m.Description = r.substituteParams(m.Description, rtm.Description)

	// inherit bodies
	m.Bodies.inherit(rtm.Bodies, r, r.Type.Parameters)

	// inherit headers
	m.inheritHeaders(r, rtm.Headers)

	// inherit query params
	m.inheritQueryParams(r, rtm.QueryParameters)

	// inherit response
	m.inheritResponses(r, rtm.Responses, r.Type.Parameters)

	// inherit protocols
	m.inheritProtocols(rtm.Protocols)
}

// inherit from all traits
func (m *Method) inheritFromAllTraits(r *Resource) error {
	for _, tDef := range m.Is {
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
func (m *Method) inheritFromATrait(r *Resource, t *Trait, dicts map[string]interface{}) error {

	m.Description = substituteParams(r, m.Description, t.Description, dicts)

	m.Bodies.inherit(t.Bodies, r, dicts)

	m.inheritHeaders(r, t.Headers)

	m.inheritResponses(r, t.Responses, dicts)

	m.inheritQueryParams(r, t.QueryParameters)

	m.inheritProtocols(t.Protocols)

	// optional bodies
	// optional headers
	// optional responses
	// optional query parameters
	return nil
}

// inheritHeaders inherit method's headers from parent headers.
// parent headers could be from resource type or a trait
func (m *Method) inheritHeaders(r *Resource, parent map[HTTPHeader]Header) {
	if len(m.Headers) == 0 {
		m.Headers = map[HTTPHeader]Header{}
	}
	for name, ph := range parent {
		h, ok := m.Headers[name]
		if !ok {
			h = Header{}
		}
		np := NamedParameter(h)
		np.inherit(NamedParameter(ph), r)
		m.Headers[name] = Header(np)
	}
}

// inheritQueryParams inherit method's query params from parent query params.
// parent query params could be from resource type or a trait
func (m *Method) inheritQueryParams(r *Resource, parent map[string]NamedParameter) {
	if len(m.QueryParameters) == 0 {
		m.QueryParameters = map[string]NamedParameter{}
	}
	for name, qp := range parent {
		nQp := newQueryParam(r, name, qp)
		m.QueryParameters[nQp.Name] = nQp
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
func (m *Method) inheritResponses(r *Resource, parent map[HTTPCode]Response, dicts map[string]interface{}) {
	if len(m.Responses) == 0 { // allocate if needed
		m.Responses = map[HTTPCode]Response{}
	}
	for code, rParent := range parent {
		resp, ok := m.Responses[code]
		if !ok {
			resp = Response{}
		}
		resp.inherit(r, rParent, dicts)
		m.Responses[code] = resp
	}

}

// inherit from parent response
func (resp *Response) inherit(r *Resource, parent Response, dicts map[string]interface{}) {
	resp.Bodies.Type = substituteParams(r, resp.Bodies.Type, parent.Bodies.Type, dicts)
}

// create new query params from another query params owned by resource type
func newQueryParam(r *Resource, name string, params NamedParameter) NamedParameter {
	return NamedParameter{
		Name:        r.substituteParams("", name),
		DisplayName: r.substituteParams("", params.DisplayName),
		Description: r.substituteParams("", params.Description),
		Type:        r.substituteParams("", params.Type),
		Enum:        params.Enum,
		Pattern:     params.Pattern,
		MinLength:   params.MinLength,
		MaxLength:   params.MaxLength,
		Minimum:     params.Minimum,
		Maximum:     params.Maximum,
		Example:     params.Example,
		Repeat:      params.Repeat,
		Required:    params.Required,
		Default:     params.Default,
		format:      params.format,
	}
}

// inherit inherits bodies properties from a parent bodies
// parent object could be from trait or response type
func (b *Bodies) inherit(parent Bodies, r *Resource, dicts map[string]interface{}) {
	b.DefaultSchema = substituteParams(r, b.DefaultSchema, parent.DefaultSchema, dicts)
	b.DefaultDescription = substituteParams(r, b.DefaultDescription, parent.DefaultDescription, dicts)
	b.DefaultExample = substituteParams(r, b.DefaultExample, parent.DefaultExample, dicts)

	b.Type = substituteParams(r, b.Type, parent.Type, dicts)

	if parent.ApplicationJson != nil {
		if b.ApplicationJson == nil {
			b.ApplicationJson = &BodiesProperty{Properties: map[string]interface{}{}}
		}

		b.ApplicationJson.Type = substituteParams(r, b.ApplicationJson.Type, parent.ApplicationJson.Type, dicts)

		for k, p := range parent.ApplicationJson.Properties {
			if _, ok := b.ApplicationJson.Properties[k]; !ok { // only inherits property that not exist
				b.ApplicationJson.Properties[k] = p
			}
		}
	}
}
