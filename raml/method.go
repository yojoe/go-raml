package raml

import (
	"regexp"
	"strings"
)

// inherit from resource type
func (m *Method) inherit(r *Resource, rtm *ResourceTypeMethod, rt *ResourceType) {
	if rtm == nil {
		return
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
	removeParamBracket := func(param string) string {
		return param[2 : len(param)-2]
	}
	// inherit type
	re, err := regexp.Compile(`\<<([^]]+)\>>`)
	if err != nil {
		panic(err)
	}
	params := re.FindAllString(parent.Bodies.Type, -1)
	for _, p := range params {
		pVal := r.getResourceTypeParamValue(removeParamBracket(p), rt)
		resp.Bodies.Type = strings.Replace(parent.Bodies.Type, p, pVal, -1)
	}
}
