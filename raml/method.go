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
func (m *Method) inheritFromResourceType(r *Resource, rtm *ResourceTypeMethod, rt *ResourceType) {
	if rtm == nil {
		return
	}
	dicts := initResourceTypeDicts(r, r.Type.Parameters)

	// inherit description
	m.Description = substituteParams(m.Description, rtm.Description, dicts)

	// inherit bodies
	m.Bodies.inherit(rtm.Bodies, r, dicts)

	// inherit headers
	m.inheritHeaders(r, rtm.Headers, dicts)

	// inherit query params
	m.inheritQueryParams(r, rtm.QueryParameters, dicts)

	// inherit response
	m.inheritResponses(r, rtm.Responses, dicts)

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
// dicts is map of trait parameters values
func (m *Method) inheritFromATrait(r *Resource, t *Trait, dicts map[string]interface{}) error {
	// initialize dicts if it still empty
	dicts = initTraitDicts(r, m, dicts)

	m.Description = substituteParams(m.Description, t.Description, dicts)

	m.Bodies.inherit(t.Bodies, r, dicts)

	m.inheritHeaders(r, t.Headers, dicts)

	m.inheritResponses(r, t.Responses, dicts)

	m.inheritQueryParams(r, t.QueryParameters, dicts)

	m.inheritProtocols(t.Protocols)

	// optional bodies
	// optional headers
	// optional responses
	// optional query parameters
	return nil
}

// inheritHeaders inherit method's headers from parent headers.
// parent headers could be from resource type or a trait
func (m *Method) inheritHeaders(r *Resource, parent map[HTTPHeader]Header, dicts map[string]interface{}) {
	if len(m.Headers) == 0 {
		m.Headers = map[HTTPHeader]Header{}
	}
	for name, ph := range parent {
		h, ok := m.Headers[name]
		if !ok {
			h = Header{}
		}
		np := NamedParameter(h)
		np.inherit(NamedParameter(ph), dicts)
		m.Headers[name] = Header(np)
	}
}

// inheritQueryParams inherit method's query params from parent query params.
// parent query params could be from resource type or a trait
func (m *Method) inheritQueryParams(r *Resource, parent map[string]NamedParameter, dicts map[string]interface{}) {
	if len(m.QueryParameters) == 0 {
		m.QueryParameters = map[string]NamedParameter{}
	}
	for name, qp := range parent {
		nQp := newQueryParam(r, name, qp, dicts)
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
		resp.inherit(rParent, dicts)
		m.Responses[code] = resp
	}

}

// inherit from parent response
func (resp *Response) inherit(parent Response, dicts map[string]interface{}) {
	resp.Bodies.Type = substituteParams(resp.Bodies.Type, parent.Bodies.Type, dicts)
}

// create new query params from another query params owned by resource type
func newQueryParam(r *Resource, name string, params NamedParameter, dicts map[string]interface{}) NamedParameter {
	return NamedParameter{
		Name:        substituteParams("", name, dicts),
		DisplayName: substituteParams("", params.DisplayName, dicts),
		Description: substituteParams("", params.Description, dicts),
		Type:        substituteParams("", params.Type, dicts),
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
	b.DefaultSchema = substituteParams(b.DefaultSchema, parent.DefaultSchema, dicts)
	b.DefaultDescription = substituteParams(b.DefaultDescription, parent.DefaultDescription, dicts)
	b.DefaultExample = substituteParams(b.DefaultExample, parent.DefaultExample, dicts)

	b.Type = substituteParams(b.Type, parent.Type, dicts)

	if parent.ApplicationJson != nil {
		if b.ApplicationJson == nil {
			b.ApplicationJson = &BodiesProperty{Properties: map[string]interface{}{}}
		}

		b.ApplicationJson.Type = substituteParams(b.ApplicationJson.Type, parent.ApplicationJson.Type, dicts)

		for k, p := range parent.ApplicationJson.Properties {
			if _, ok := b.ApplicationJson.Properties[k]; !ok { // only inherits property that not exist
				b.ApplicationJson.Properties[k] = p
			}
		}
	}
}
