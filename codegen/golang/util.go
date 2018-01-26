package golang

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
)

// try to set root import path based on target dir
// - do nothing if user specify import path
// - if target dir is under GOPATH, set it to the target dir
func setRootImportPath(importPath, targetDir string) string {
	// use import path if not empty
	if len(importPath) > 0 {
		return importPath
	}

	// get GOPATH dir
	gopath, err := filepath.Abs(os.Getenv("GOPATH"))
	if err != nil {
		return ""
	}
	gopath = filepath.Join(gopath, "src")

	// get absolute target dir
	absTargetDir, err := filepath.Abs(targetDir)
	if err != nil {
		panic("invalid targetDir:" + err.Error())
	}

	// panic if user doesn't specify import path and target dir
	// not under GOPATH.
	if !strings.HasPrefix(absTargetDir, gopath) {
		panic("please specify '--import-path' or set '--dir' under your GOPATH")
	}

	// set import path
	newImportPath, err := filepath.Rel(gopath, absTargetDir)
	if err != nil {
		panic("failed to set import path automatically:" + err.Error())
	}

	// re-join because otherwise windows will use `\`
	return path.Join(strings.Split(newImportPath, string(filepath.Separator))...)
}

// check that the API definition has duplicated types
// of the title case version of the types.
// title case = first letter become uppercase
// example of non duplicate:
//  - One & Two
//  - One & oNe = One & ONe -> N is uppercase
// example of duplicate
// - One & one = One & One
func checkDuplicatedTitleTypes(apiDef *raml.APIDefinition) error {
	var title string

	for name := range apiDef.Types {
		title = strings.Title(name)
		if title == name {
			continue
		}
		if _, duplicate := apiDef.Types[title]; duplicate {
			return fmt.Errorf("types conflict: %s with %v", name, title)
		}
	}
	return nil
}
