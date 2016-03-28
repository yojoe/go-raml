package commands

import "github.com/Jumpscale/go-raml/raml"

const (
	reqBodySuffix  = "ReqBody"
	respBodySuffix = "RespBody"
)

// generate Go struct from RAML definition
func generateBodyStructs(apiDef *raml.APIDefinition, dir, packageName string) error {
	// generate
	for _, v := range apiDef.Resources {
		if err := generateStructsFromResourceBody("", dir, packageName, &v); err != nil {
			return err
		}
	}

	return nil
}

// generate all structs from a resource
func generateStructsFromResourceBody(resourcePath, dir, packageName string, r *raml.Resource) error {
	if r == nil {
		return nil
	}

	//build struct name
	structName := normalizeURITitle(resourcePath + r.URI)

	//build
	if err := buildBodyFromMethod(structName, "Get", dir, packageName, r.Get); err != nil {
		return err
	}

	if err := buildBodyFromMethod(structName, "Post", dir, packageName, r.Post); err != nil {
		return err
	}

	if err := buildBodyFromMethod(structName, "Head", dir, packageName, r.Head); err != nil {
		return err
	}

	if err := buildBodyFromMethod(structName, "Put", dir, packageName, r.Put); err != nil {
		return err
	}

	if err := buildBodyFromMethod(structName, "Delete", dir, packageName, r.Delete); err != nil {
		return err
	}

	if err := buildBodyFromMethod(structName, "Patch", dir, packageName, r.Patch); err != nil {
		return err
	}

	//build child
	for _, v := range r.Nested {
		if err := generateStructsFromResourceBody(resourcePath+r.URI, dir, packageName, v); err != nil {
			return err
		}
	}

	return nil
}

func buildBodyFromMethod(structName, methodName, dir, packageName string, method *raml.Method) error {
	if method == nil {
		return nil
	}

	//generate struct for body node below method
	if err := generateStructFromBody(structName+methodName, dir, packageName, &method.Bodies, true); err != nil {
		return err
	}

	//generate struct for body node below response
	for _, val := range method.Responses {
		if err := generateStructFromBody(structName+methodName, dir, packageName, &val.Bodies, false); err != nil {
			return err
		}

	}

	return nil
}

// generate a struct from an RAML request/response body
func generateStructFromBody(structNamePrefix, dir, packageName string, body *raml.Bodies, isGenerateRequest bool) error {
	// check if this method has JSON body that need to be generated
	hasJSONBody := func() bool {
		return body.ApplicationJson != nil && len(body.ApplicationJson.Properties) > 0
	}
	if body == nil || !hasJSONBody() {
		return nil
	}

	// construct struct from body
	structDef := newStructDefFromBody(body, structNamePrefix, packageName, isGenerateRequest)

	// generate
	return structDef.generate(dir)
}
