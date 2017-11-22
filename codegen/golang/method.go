package golang

import (
	"fmt"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/codegen/security"
	"github.com/Jumpscale/go-raml/raml"
)

// defines go method base object
type method struct {
	resource.Method
	ResourcePath string
	ReqBody      string
	PackageName  string
	resps        []respBody
}

// TODO : move it to codegen/resource
type respBody struct {
	Code int
	Type string
}

func newMethod(resMeth resource.Method) *method {
	var resps []respBody
	// creates response body
	for code, resp := range resMeth.Responses {
		resp := respBody{
			Code: commons.AtoiOrPanic(string(code)),
			Type: setBodyName(resp.Bodies, resMeth.Endpoint+resMeth.VerbTitle(), commons.RespBodySuffix),
		}
		if resp.Type != "" {
			resps = append(resps, resp)
		}
	}

	// normalized endpoint
	normalizedEndpoint := commons.NormalizeURITitle(resMeth.Endpoint)

	return &method{
		Method:       resMeth,
		ResourcePath: commons.ParamizingURI(resMeth.Endpoint, "+"),
		ReqBody:      setBodyName(resMeth.Bodies, normalizedEndpoint+resMeth.VerbTitle(), commons.ReqBodySuffix),
		resps:        resps,
	}
}

func (m method) HasRespBody() bool {
	return len(m.RespBodyTypes()) > 0
}

// RespBodyTypes returns all possible type of response body
func (m method) RespBodyTypes() []respBody {
	return m.resps
}

// FailedRespBodyTypes return all response body that considered a failed response
// i.e. non 2xx status code
func (m method) FailedRespBodyTypes() (resps []respBody) {
	for _, resp := range m.RespBodyTypes() {
		if resp.Code < 200 || resp.Code >= 300 {
			resps = append(resps, resp)
		}
	}
	return
}

// SuccessRespBodyTypes returns all response body that considered as success
// i.e. 2xx status code
func (m method) SuccessRespBodyTypes() (resps []respBody) {
	for _, resp := range m.RespBodyTypes() {
		if resp.Code >= 200 && resp.Code < 300 {
			resps = append(resps, resp)
		}
	}
	return
}

func (m method) firstSuccessRespBodyType() string {
	resps := m.SuccessRespBodyTypes()
	if len(resps) == 0 {
		return ""
	}
	return resps[0].Type
}

// get oauth2 middleware handler from a security scheme
func getOauth2MwrHandler(ss raml.DefinitionChoice) (string, error) {
	// construct security scopes
	quotedScopes, err := security.GetQuotedScopes(ss)
	if err != nil {
		return "", err
	}
	scopesArgs := strings.Join(quotedScopes, ", ")

	// middleware name
	// need to handle case where it reside in different package
	var packageName string
	name := ss.Name

	if splitted := strings.Split(name, "."); len(splitted) == 2 {
		packageName = splitted[0]
		name = splitted[1]
	}
	mwr := fmt.Sprintf(`NewOauth2%vMiddleware([]string{%v}).Handler`, name, scopesArgs)
	if packageName != "" {
		mwr = packageName + "." + mwr
	}
	return mwr, nil
}

// setBodyName set name of method's request/response body.
//
// Rules:
//	- use bodies.Type if not empty and not `object`
//	- use bodies.ApplicationJSON.Type if not empty and not `object`
//	- use prefix+suffix if:
//		- not meet previous rules
//		- previous rules produces JSON string
func setBodyName(bodies raml.Bodies, prefix, suffix string) string {
	var tipe string
	prefix = commons.NormalizeURITitle(prefix)

	if len(bodies.Type) > 0 && bodies.Type != "object" {
		tipe = convertToGoType(bodies.Type, "")
	} else if bodies.ApplicationJSON != nil {
		if bodies.ApplicationJSON.TypeString() != "" && bodies.ApplicationJSON.TypeString() != "object" {
			tipe = convertToGoType(bodies.ApplicationJSON.TypeString(), "")
		} else {
			tipe = prefix + suffix
		}
	}

	if commons.IsJSONString(tipe) {
		tipe = prefix + suffix
	}

	return tipe
}
