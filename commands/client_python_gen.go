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

type pythonMethod struct {
	interfaceMethod
	Params  string
	URIArgs string
	PRArgs  string // python requests's args
}

type pythonClientDef struct {
	clientDef
	PythonMethods []pythonMethod
}

func (pcd pythonClientDef) generate(dir string) error {
	if err := generateFile(nil, pythonClientUtilTmplFile, pythonClientUtilTmplName, dir+"/client_utils.py", false); err != nil {
		return err
	}
	return generateFile(pcd, pythonClientTmplFile, pythonClientTmplName, dir+"/client.py", true)
}

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
			URIArgs:         paramizingURI(`"` + completeResourceURI(m.Resource) + `"`),
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
