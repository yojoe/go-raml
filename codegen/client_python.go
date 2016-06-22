package codegen

import (
	"path/filepath"
)

// python client definition
type pythonClient struct {
	clientDef
}

// generate python lib files
func (pc pythonClient) generate(dir string) error {
	// generate helper
	if err := generateFile(nil, "./templates/client_utils_python.tmpl", "client_utils_python", filepath.Join(dir, "client_utils.py"), false); err != nil {
		return err
	}
	// generate main client lib file
	return generateFile(pc, "./templates/client_python.tmpl", "client_python", filepath.Join(dir, "client.py"), true)
}
