package codegen

const (
	pythonClientTmplFile     = "./templates/python_client.tmpl"
	pythonClientUtilTmplFile = "./templates/python_client_utils.tmpl"
	pythonClientTmplName     = "python_client_template"
	pythonClientUtilTmplName = "python_client_utils"
)

// python client definition
type pythonClient struct {
	clientDef
}

// generate python lib files
func (pc pythonClient) generate(dir string) error {
	// generate helper
	if err := generateFile(nil, pythonClientUtilTmplFile, pythonClientUtilTmplName, dir+"/client_utils.py", false); err != nil {
		return err
	}
	// generate main client lib file
	return generateFile(pc, pythonClientTmplFile, pythonClientTmplName, dir+"/client.py", true)
}
