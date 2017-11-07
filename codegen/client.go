package codegen

import (
	"github.com/Jumpscale/go-raml/codegen/golang"
	"github.com/Jumpscale/go-raml/codegen/nim"
	"github.com/Jumpscale/go-raml/codegen/python"
	"github.com/Jumpscale/go-raml/raml"
)

// GenerateClient generates client library
func GenerateClient(apiDef *raml.APIDefinition, dir, packageName, lang, rootImportPath, kind string,
	libRootURLs []string, pythonUnnmarshallResponse bool) error {
	switch lang {
	case langGo:
		gc, err := golang.NewClient(apiDef, packageName, rootImportPath, dir,
			libRootURLs)
		if err != nil {
			return err
		}
		return gc.Generate()
	case langPython:
		pc := python.NewClient(apiDef, kind, pythonUnnmarshallResponse)
		return pc.Generate(dir)
	case langNim:
		nc := nim.NewClient(apiDef, dir)
		return nc.Generate()
	default:
		return errInvalidLang
	}
}
