package commands

import (
	"strings"
)

const (
	pythonClientTmplFile = "./templates/python_client.tmpl"
	pythonClientTmplName = "python_client_template"
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
		pm := pythonMethod{
			interfaceMethod: m,
			Params:          strings.Join(append(params, getResourceParams(m.Resource)...), ","),
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
