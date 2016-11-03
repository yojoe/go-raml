package codegen

import (
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

// library defines an RAML library
// it is implemented as package in Go.
// Library defined in RAML using `uses` keyword.
// - key become library's package name
// - if value contain directory, the directories become root directory
//   of the generated package
// example :
// files: libraries/files.raml -> generated as `files` package in `libraries` directory
// types-lib: lib-types.raml  -> generated as `types_lib` package in current directory
type goLibrary struct {
	*raml.Library
	PackageName string
	baseDir     string // root directory
	dir         string // library directory
}

// create new library instance
func newGoLibrary(name string, lib *raml.Library, baseDir string) *goLibrary {
	return &goLibrary{
		Library:     lib,
		baseDir:     baseDir,
		PackageName: commons.NormalizePkgName(name),
		dir:         commons.NormalizePkgName(filepath.Join(baseDir, goLibPackageDir(name, lib.Filename))),
	}
}

// generate code of all libraries
func generateLibraries(libraries map[string]*raml.Library, baseDir string) error {
	for name, ramlLib := range libraries {
		l := newGoLibrary(name, ramlLib, baseDir)
		if err := l.generate(); err != nil {
			return err
		}
	}
	return nil
}

// generate code of this library
func (l *goLibrary) generate() error {
	if err := commons.CheckCreateDir(l.dir); err != nil {
		return err
	}

	// generate all Type structs
	if err := generateStructs(l.Types, l.dir, l.PackageName, langGo); err != nil {
		return err
	}

	// security schemes
	if err := generateSecurity(l.SecuritySchemes, l.dir, l.PackageName); err != nil {
		return err
	}

	// included libraries
	for name, ramlLib := range l.Libraries {
		childLib := newGoLibrary(name, ramlLib, l.baseDir)
		if err := childLib.generate(); err != nil {
			return err
		}
	}
	return nil
}

// get library import path from a type
func libImportPath(rootImportPath, typ string) string {
	// library use '.', return nothing if it is not a library
	if strings.Index(typ, ".") < 0 {
		return ""
	}

	// library name in the current document
	libName := strings.Split(typ, ".")[0]

	if libName == "goraml" { // special package name, reserved for goraml
		return filepath.Join(rootImportPath, "goraml")
	}

	// raml file of this lib
	libRAMLFile := globAPIDef.FindLibFile(commons.DenormalizePkgName(libName))

	if libRAMLFile == "" {
		log.Fatalf("can't find library : %v", libName)
	}

	return filepath.Join(rootImportPath, goLibPackageDir(libName, libRAMLFile))
}

// returns Go package directory of a library
// name is library name. filename is library file name.
// for the rule, see comment of `type goLibrary struct`
func goLibPackageDir(name, filename string) string {
	return commons.NormalizePkgName(filepath.Join(filepath.Dir(filename), name))
}
