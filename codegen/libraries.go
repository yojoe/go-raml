package codegen

import (
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
)

// library defines an RAML library
// it is implemented as package in Go
type library struct {
	*raml.Library
	PackageName string
	baseDir     string // root directory
	dir         string // library directory
}

func newLibrary(lib *raml.Library, baseDir string) *library {
	l := library{
		Library: lib,
		baseDir: baseDir,
	}

	// package name  : base of filename without the extension
	l.PackageName = strings.TrimSuffix(filepath.Base(l.Filename), filepath.Ext(l.Filename))
	l.PackageName = normalizePkgName(l.PackageName)

	// package directory : filename without the extension
	relDir := libRelDir(l.Filename) //strings.TrimSuffix(l.Filename, filepath.Ext(l.Filename))
	l.dir = filepath.Join(l.baseDir, relDir)

	return &l
}

func generateLibraries(libraries map[string]*raml.Library, baseDir string) error {
	for _, ramlLib := range libraries {
		l := newLibrary(ramlLib, baseDir)
		if err := l.generate(); err != nil {
			return err
		}
	}
	return nil
}

func (l *library) generate() error {
	if err := checkCreateDir(l.dir); err != nil {
		return err
	}

	// generate dates
	dg := dateGen{PackageName: l.PackageName}

	if err := dg.generate(l.dir); err != nil {
		log.Errorf("library.generate() failed to generate date files:%v", err)
		return err
	}

	// generate input validator
	if err := generateInputValidator(l.PackageName, l.dir); err != nil {
		return err
	}

	// generate all Type structs
	if err := generateStructs(l.Types, l.dir, l.PackageName, langGo); err != nil {
		return err
	}

	// security schemes
	if err := generateSecurity(l.SecuritySchemes, l.dir, l.PackageName, langGo); err != nil {
		return err
	}

	// included libraries
	for _, ramlLib := range l.Libraries {
		childLib := newLibrary(ramlLib, l.baseDir)
		if err := childLib.generate(); err != nil {
			return err
		}
	}
	return nil
}

// get relative lib directory from library filename
// this relative directory will be used as:
// - lib package name
// - lib files directory
func libRelDir(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

// get lib import path from a type
func libImportPath(rootPath, typ string) string {
	if strings.Index(typ, ".") < 0 {
		return ""
	}
	// library name in the current document
	libName := strings.Split(typ, ".")[0]

	// raml file of this lib
	libRAMLFile := globAPIDef.FindLibFile(denormalizePkgName(libName))

	if libRAMLFile == "" {
		log.Fatalf("can't find library : %v", libName)
	}

	// relative lib package
	libPkg := libRelDir(libRAMLFile)

	return filepath.Join(rootImportPath, normalizePkgName(libPkg))
}

// normalize package name because not all characters can be used as package name
func normalizePkgName(name string) string {
	return strings.Replace(name, "-", "_", -1)
}

// inverse of normalizePkgName
func denormalizePkgName(name string) string {
	return strings.Replace(name, "_", "-", -1)
}
