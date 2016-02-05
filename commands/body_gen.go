package commands

import "github.com/Jumpscale/go-raml/raml"

const (
	reqBodySuffix  = "ReqBody"
	respBodySuffix = "RespBody"
)

//GenerateBodyStruct generate body
func GenerateBodyStruct(dir string, apiDef *raml.APIDefinition) error {
	//check create dir
	if err := checkCreateDir(dir); err != nil {
		return err
	}

	//generate
	for _, v := range apiDef.Resources {
		if err := generateBodyFromResources("", dir, &v); err != nil {
			return err
		}
	}

	return nil
}

func generateBodyFromResources(resourcePath, dir string, r *raml.Resource) error {
	if r == nil {
		return nil
	}

	//build struct name
	structName := normalizeURITitle(resourcePath + r.URI)

	//build
	if err := buildBodyFromMethod(structName, "Get", dir, r.Get); err != nil {
		return err
	}

	if err := buildBodyFromMethod(structName, "Post", dir, r.Post); err != nil {
		return err
	}

	if err := buildBodyFromMethod(structName, "Head", dir, r.Head); err != nil {
		return err
	}

	if err := buildBodyFromMethod(structName, "Put", dir, r.Put); err != nil {
		return err
	}

	if err := buildBodyFromMethod(structName, "Delete", dir, r.Delete); err != nil {
		return err
	}

	if err := buildBodyFromMethod(structName, "Patch", dir, r.Patch); err != nil {
		return err
	}

	//build child
	for _, v := range r.Nested {
		if err := generateBodyFromResources(resourcePath+r.URI, dir, v); err != nil {
			return err
		}
	}

	return nil
}

func buildBodyFromMethod(structName, methodName, dir string, method *raml.Method) error {
	if method == nil {
		return nil
	}

	//generate struct for body node below method
	if err := generateStructFromBody(structName+methodName, dir, &method.Bodies, false); err != nil {
		return err
	}

	//generate struct for body node below response
	for _, val := range method.Responses {
		if err := generateStructFromBody(structName+methodName, dir, &val.Bodies, true); err != nil {
			return err
		}

	}

	return nil
}

func generateStructFromBody(structNamePrefix, dir string, body *raml.Bodies, isPartial bool) error {
	if body == nil || len(body.ApplicationJson.Properties) == 0 || len(body.ApplicationJson.Type) > 0 {
		return nil
	}

	//construct struct from body
	reqStructDef, respStructDef := newStructDefFromBody(body, structNamePrefix, isPartial)
	if !isPartial {
		if err := reqStructDef.generate(dir); err != nil {
			return err
		}
	}

	//generate
	return respStructDef.generate(dir)
}
