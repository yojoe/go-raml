package codegen

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

type goServerMethod struct {
	*method
	Middlewares string
}

// setup go server method, initializes all needed variables
func (gm *goServerMethod) setup(apiDef *raml.APIDefinition, r *raml.Resource, rd *resourceDef, methodName string) error {
	// set method name
	name := normalizeURI(gm.Endpoint)
	if len(gm.DisplayName) > 0 {
		gm.MethodName = strings.Replace(gm.DisplayName, " ", "", -1)
	} else {
		gm.MethodName = name[len(rd.Name):] + methodName
	}

	// setting middlewares
	middlewares := []string{}

	// security middlewares
	for _, v := range gm.SecuredBy {
		if !validateSecurityScheme(v.Name, apiDef) {
			continue
		}
		// oauth2 middleware
		m, err := getOauth2MwrHandler(v)
		if err != nil {
			return err
		}
		middlewares = append(middlewares, m)
	}

	gm.Middlewares = strings.Join(middlewares, ", ")

	return nil
}

type goClientMethod struct {
	*method
}

func (gcm *goClientMethod) setup(methodName string) error {
	// build func/method params
	buildParams := func(r *raml.Resource, bodyType string) (string, error) {
		paramsStr := strings.Join(getResourceParams(r), ",")
		if len(paramsStr) > 0 {
			paramsStr += " string"
		}

		// append request body type
		if len(bodyType) > 0 {
			if len(paramsStr) > 0 {
				paramsStr += ", "
			}
			paramsStr += strings.ToLower(bodyType) + " " + bodyType
		}

		// append header
		if len(paramsStr) > 0 {
			paramsStr += ","
		}
		paramsStr += "headers,queryParams map[string]interface{}"

		return paramsStr, nil
	}

	// method name
	name := normalizeURITitle(gcm.Endpoint)

	if len(gcm.DisplayName) > 0 {
		gcm.MethodName = strings.Replace(gcm.DisplayName, " ", "", -1)
	} else {
		gcm.MethodName = strings.Title(name + methodName)
	}

	// method param
	methodParam, err := buildParams(gcm.resource, gcm.ReqBody)
	if err != nil {
		return err
	}
	gcm.Params = methodParam

	return nil
}
