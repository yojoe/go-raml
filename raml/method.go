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
	m.Bodies.inherit(rtm.Bodies, r, rt)

	// inherit headers
	if len(m.Headers) == 0 {
		m.Headers = map[HTTPHeader]Header{}
	}
	for name, ph := range rtm.Headers {
		h, ok := m.Headers[name]
		if !ok {
			h = Header{}
		}
		np := NamedParameter(h)
		np.inherit(NamedParameter(ph), r)
		m.Headers[name] = Header(np)
	}

	// inherit query params
	if len(m.QueryParameters) == 0 {
		m.QueryParameters = map[string]NamedParameter{}
	}
	for name, qp := range rtm.QueryParameters {
		nQp := newQueryParam(r, rt, rtm, name, qp)
		m.QueryParameters[nQp.Name] = nQp
	}

	// inherit response
	if len(m.Responses) == 0 {
		m.Responses = map[HTTPCode]Response{}
	}
	for code, rParent := range rtm.Responses {
		resp, ok := m.Responses[code]
		if !ok {
			resp = Response{}
		}
		resp.inherit(r, rParent, rt)
		m.Responses[code] = resp
	}

	// inherit protocols
	for _, p := range rtm.Protocols {
		m.Protocols = appendStrNotExist(p, m.Protocols)
	}
}

// inherit from all traits
func (m *Method) inheritAllTraits() error {
	for _, tDef := range m.Is {
		// acquire traits object
		t, ok := traitsMap[tDef.Name]
		if !ok {
			return fmt.Errorf("invalid traits name:%v", tDef.Name)
		}

		if err := m.inheritTraits(&t, tDef.Parameters); err != nil {
			return err
		}
	}
	return nil
}

// inherit from a trait
func (m *Method) inheritTraits(t *Trait, params map[string]interface{}) error {
	return nil
}

// inherit from resource type
func (resp *Response) inherit(r *Resource, parent Response, rt *ResourceType) {
	resp.Bodies.Type = r.substituteParams(resp.Bodies.Type, parent.Bodies.Type)
}

// create new query params from another query params owned by resource type
func newQueryParam(r *Resource, rt *ResourceType, rtm *ResourceTypeMethod, name string, params NamedParameter) NamedParameter {
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

func (b *Bodies) inherit(parent Bodies, r *Resource, rt *ResourceType) {
	b.DefaultSchema = r.substituteParams(b.DefaultSchema, parent.DefaultSchema)
	b.DefaultDescription = r.substituteParams(b.DefaultDescription, parent.DefaultDescription)
	b.DefaultExample = r.substituteParams(b.DefaultExample, parent.DefaultExample)

	b.Type = r.substituteParams(b.Type, parent.Type)

	if parent.ApplicationJson != nil {
		if b.ApplicationJson == nil {
			b.ApplicationJson = &BodiesProperty{Properties: map[string]interface{}{}}
		}
		b.ApplicationJson.Type = r.substituteParams(b.ApplicationJson.Type, parent.ApplicationJson.Type)
		for k, p := range parent.ApplicationJson.Properties {
			b.ApplicationJson.Properties[k] = p
		}
	}
}
