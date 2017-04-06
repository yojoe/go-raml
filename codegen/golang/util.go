package golang

import (
	"os"
	"path/filepath"
	"strings"
)

func setRootImportPath(importPath, targetDir string) string {
	// use import path if not empty
	if len(importPath) > 0 {
		return importPath
	}

	gopath, err := filepath.Abs(os.Getenv("GOPATH"))
	if err != nil {
		return ""
	}
	gopath = filepath.Join(gopath, "src")

	absTargetDir, err := filepath.Abs(targetDir)
	if err != nil {
		panic("invalid targetDir:" + err.Error())
	}
	// return empty  path if targetDir not under GOPATH
	if !strings.HasPrefix(absTargetDir, gopath) {
		return ""
	}
	newImportPath, err := filepath.Rel(gopath, absTargetDir)
	if err != nil {
		panic("failed to set import path automatically:" + err.Error())
	}
	return newImportPath
}
