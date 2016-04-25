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
	m.Description = r.substituteParams(rtm.Description)

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
	resp.Bodies.Type = r.substituteParams(parent.Bodies.Type)
}
