package python

import (
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/codegen/security"
	"github.com/Jumpscale/go-raml/raml"
	"github.com/pinzolo/casee"
)

// python server method
type serverMethod struct {
	*method
	MiddlewaresArr []middleware
	SecuredBy      []raml.DefinitionChoice
}

// setup sets all needed variables
func (sm *serverMethod) setup(apiDef *raml.APIDefinition, r *raml.Resource, rd *resource.Resource, resourceParams []string) error {
	sm.MethodName = commons.SetServerMethodName(sm.DisplayName, sm.Verb(), r)

	if commons.HasJSONBody(&(sm.Bodies)) {
		// TODO : make it to call proper func
		sm.ReqBody = casee.ToPascalCase(sm.MethodName + commons.ReqBodySuffix)
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
	sm.ReqBody = setServerReqBodyName(meth.ReqBody)

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

func setServerReqBodyName(bodyName string) string {
	if !commons.IsArrayType(bodyName) {
		return bodyName
	}
	return jsArrayName(bodyName)
}
