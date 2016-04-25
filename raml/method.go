package raml

func newMethod(name string) *Method {
	return &Method{
		Name: name,
	}
}

// inherit from resource type
// fields need to be inherited:
// - description
// - response
func (m *Method) inherit(r *Resource, rtm *ResourceTypeMethod, rt *ResourceType) {
	if rtm == nil {
		return
	}

	// inherit description
	m.Description = r.substituteParams(rtm.Description, rt)

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
}

// inherit from resource type
func (resp *Response) inherit(r *Resource, parent Response, rt *ResourceType) {
	resp.Bodies.Type = r.substituteParams(parent.Bodies.Type, rt)
}

// create new query params from another query params owned by resource type
func newQueryParam(r *Resource, rt *ResourceType, rtm *ResourceTypeMethod, name string, params NamedParameter) NamedParameter {
	return NamedParameter{
		Name:        r.substituteParams(name, rt),
		DisplayName: r.substituteParams(params.DisplayName, rt),
		Description: r.substituteParams(params.Description, rt),
		Type:        r.substituteParams(params.Type, rt),
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
