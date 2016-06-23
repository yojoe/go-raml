package codegen

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
)

type pythonLibrary struct {
	*raml.Library
	baseDir string
	dir     string
}

func newPythonLibrary(lib *raml.Library, baseDir string) *pythonLibrary {
	pl := pythonLibrary{
		Library: lib,
		baseDir: baseDir,
	}

	// package directory : filename without the extension
	relDir := libRelDir(pl.Filename)
	pl.dir = normalizePkgName(filepath.Join(pl.baseDir, relDir))

	return &pl
}

// generate code of all libraries
func generatePythonLibraries(libraries map[string]*raml.Library, baseDir string) error {
	for _, ramlLib := range libraries {
		pl := newPythonLibrary(ramlLib, baseDir)

		if err := pl.generate(); err != nil {
			return err
		}
	}
	return nil
}

// generate code of this library
func (l *pythonLibrary) generate() error {
	// create directory if needed
	if err := checkCreateDir(l.dir); err != nil {
		return err
	}

	// write empty __init__.py in each dir
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		return generateEmptyInitPy(path)
	}

	if err := filepath.Walk(l.baseDir, walkFn); err != nil {
		return err
	}

	// python classes
	if err := generatePythonClasses(l.Types, l.dir); err != nil {
		log.Errorf("failed to generate python clased:%v", err)
		return err
	}

	// security schemes
	if err := generateSecurity(l.SecuritySchemes, l.dir, "", langPython); err != nil {
		return err
	}

	// included libraries
	for _, ramlLib := range l.Libraries {
		childLib := newPythonLibrary(ramlLib, l.baseDir)
		if err := childLib.generate(); err != nil {
			return err
		}
	}
	return nil
}

func pythonLibImportPath(typ, prefix string) (string, string) {
	// library use '.'
	if strings.Index(typ, ".") < 0 {
		return prefix + typ, prefix + typ
	}

	splitted := strings.Split(typ, ".")
	if len(splitted) != 2 {
		log.Fatalf("pythonLibImportPath invalid typ:" + typ)
	}
	// library name in the current document
	libName := splitted[0]

	// raml file of this lib
	libRAMLFile := globAPIDef.FindLibFile(denormalizePkgName(libName))

	if libRAMLFile == "" {
		log.Fatalf("pythonLibImportPath() can't find library : %v", libName)
	}

	// relative lib package
	libPkg := libRelDir(libRAMLFile)

	return strings.Replace(normalizePkgName(libPkg), "/", ".", -1) + "." + prefix + splitted[1], prefix + splitted[1]
}
