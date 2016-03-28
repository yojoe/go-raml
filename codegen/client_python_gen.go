package codegen

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

const (
	pythonClientTmplFile     = "./templates/python_client.tmpl"
	pythonClientUtilTmplFile = "./templates/python_client_utils.tmpl"
	pythonClientTmplName     = "python_client_template"
	pythonClientUtilTmplName = "python_client_utils"
)

// defines a python client lib method
type pythonMethod struct {
	interfaceMethod
	PRArgs string // python requests's args
}

// python client definition
type pythonClientDef struct {
	clientDef
	PythonMethods []pythonMethod
}

// generate python lib files
func (pcd pythonClientDef) generate(dir string) error {
	// generate helper
	if err := generateFile(nil, pythonClientUtilTmplFile, pythonClientUtilTmplName, dir+"/client_utils.py", false); err != nil {
		return err
	}
	// generate main client lib file
	return generateFile(pcd, pythonClientTmplFile, pythonClientTmplName, dir+"/client.py", true)
}

// generate python client lib
func (cd clientDef) generatePython(dir string) error {
	var pms []pythonMethod
	baseParams := []string{"self"}
	for _, m := range cd.Methods {
		params := baseParams
		prArgs := ""
		if m.Verb == "PUT" || m.Verb == "POST" || m.Verb == "PATCH" {
			params = append(params, "data")
			prArgs = ", data"
		}

		params = append(params, getResourceParams(m.Resource)...)

		pm := pythonMethod{
			interfaceMethod: m,
			PRArgs:          prArgs,
		}
		pm.Params = strings.Join(append(params, "headers=None, query_params=None"), ", ")

		if len(pm.DisplayName) > 0 {
			pm.MethodName = strings.Replace(pm.DisplayName, " ", "", -1)
		} else {
			pm.MethodName = snakeCaseResourceURI(m.Resource) + "_" + strings.ToLower(m.Verb)
		}

		pms = append(pms, pm)
	}

	pcd := pythonClientDef{
		clientDef:     cd,
		PythonMethods: pms,
	}
	return pcd.generate(dir)
}

// create snake case function name from a resource URI
func snakeCaseResourceURI(r *raml.Resource) string {
	return _snakeCaseResourceURI(r, "")
}

func _snakeCaseResourceURI(r *raml.Resource, completeURI string) string {
	if r == nil {
		return completeURI
	}
	var snake string
	if len(r.URI) > 0 {
		uri := normalizeURI(r.URI)
		if r.Parent != nil { // not root resource, need to add "_"
			snake = "_"
		}

		if strings.HasPrefix(r.URI, "/{") {
			snake += "by" + strings.ToUpper(uri[:1])
		} else {
			snake += strings.ToLower(uri[:1])
		}

		if len(uri) > 1 { // append with the rest of uri
			snake += uri[1:]
		}
	}
	return _snakeCaseResourceURI(r.Parent, snake+completeURI)
}
