package python

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/chuckpreslar/inflect"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/codegen/security"
	"github.com/Jumpscale/go-raml/raml"
)

// python server method
type serverMethod struct {
	*method
	MiddlewaresArr []middleware
	SecuredBy      []raml.DefinitionChoice
}

func setServerMethodName(displayName, verb string, resource *raml.Resource) string {
	if len(displayName) > 0 {
		return commons.DisplayNameToFuncName(displayName)
	}
	return snakeCaseResourceURI(resource) + "_" + strings.ToLower(verb)
}

func setReqBodyName(methodName string) string {
	return inflect.UpperCamelCase(methodName + "ReqBody")
}

// setup sets all needed variables
func (sm *serverMethod) setup(apiDef *raml.APIDefinition, r *raml.Resource, rd *resource.Resource, resourceParams []string) error {
	sm.MethodName = setServerMethodName(sm.DisplayName, sm.Verb(), r)

	if commons.HasJSONBody(&(sm.Bodies)) {
		sm.ReqBody = setReqBodyName(sm.MethodName)
	}

	sm.Params = strings.Join(resourceParams, ", ")
	sm.Endpoint = strings.Replace(sm.Endpoint, "{", "<", -1)
	sm.Endpoint = strings.Replace(sm.Endpoint, "}", ">", -1)

	// security middlewares
	for _, v := range sm.SecuredBy {
		if !security.ValidateScheme(v.Name, apiDef) {
			continue
		}
		// oauth2 middleware
		m, err := newPythonOauth2Middleware(v)
		if err != nil {
			log.Errorf("error creating middleware for method.err = %v", err)
			return err
		}
		sm.MiddlewaresArr = append(sm.MiddlewaresArr, m)
	}
	return nil
}

func newServerMethod(apiDef *raml.APIDefinition, rd *resource.Resource, rm resource.Method, kind string) serverMethod {
	meth := newMethod(rm)
	sm := serverMethod{
		method:    meth,
		SecuredBy: security.GetMethodSecuredBy(apiDef, rm.Resource, rm.Method),
	}
	params := resource.GetResourceParams(rm.Resource)
	if kind == serverKindSanic {
		params = append([]string{"request"}, params...)
	}
	sm.setup(apiDef, rm.Resource, rd, params)
	return sm

}

func (sm *serverMethod) RespBody() string {
	return sm.firstSuccessRespBodyType()
}
