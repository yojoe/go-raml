package commands

import (
	"strings"
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
	Params string // methods params
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
			Params:          strings.Join(append(params, "headers=None,queryParams=None"), ","),
			PRArgs:          prArgs,
		}
		pms = append(pms, pm)
	}

	pcd := pythonClientDef{
		clientDef:     cd,
		PythonMethods: pms,
	}
	return pcd.generate(dir)
}
