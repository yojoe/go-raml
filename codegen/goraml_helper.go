package codegen

import (
	"path/filepath"

	log "github.com/Sirupsen/logrus"
)

const (
	goramlHelperPkgName = "goraml"
)

// goramlHelper represents helper package
// that is not described in .raml file but needed by
// generated code.
type goramlHelper struct {
	rootImportPath string
}

func (gh goramlHelper) generate(dir string) error {
	pkgDir := filepath.Join(dir, goramlHelperPkgName)

	// create directory if needed
	if err := checkCreateDir(pkgDir); err != nil {
		return err
	}

	/// dates
	d := dateGen{PackageName: goramlHelperPkgName}
	if err := d.generate(pkgDir); err != nil {
		log.Errorf("generate() failed to generate date files:%v", err)
		return err
	}

	// generate struct validator
	if err := generateInputValidator(goramlHelperPkgName, pkgDir); err != nil {
		return err
	}
	return nil
}
