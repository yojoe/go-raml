package codegen

import (
	"path/filepath"
)

// python client definition
type pythonClient struct {
	clientDef
}

// generate empty __init__.py without overwrite it
func generateEmptyInitPy(dir string) error {
	return generateFile(nil, "./templates/init_py.tmpl", "init_py", filepath.Join(dir, "__init__.py"), false)
}

// generate python lib files
func (pc pythonClient) generate(dir string) error {
	// generate empty __init__.py
	if err := generateEmptyInitPy(dir); err != nil {
		return err
	}

	// generate helper
	if err := generateFile(nil, "./templates/client_utils_python.tmpl", "client_utils_python", filepath.Join(dir, "client_utils.py"), false); err != nil {
		return err
	}
	// generate main client lib file
	return generateFile(pc, "./templates/client_python.tmpl", "client_python", filepath.Join(dir, "client.py"), true)
}
