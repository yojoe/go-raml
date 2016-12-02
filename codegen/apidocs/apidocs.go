package apidocs

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Jumpscale/go-raml/raml"
)

// Generate generates API docs using api-console
// https://github.com/mulesoft/api-console
func Generate(apiDef *raml.APIDefinition, ramlFile string, ramlBytes []byte, dir string) error {
	// extract zipped files
	if err := extract(dir); err != nil {
		return err
	}
	ramlBytes = append([]byte("#%RAML 1.0\n"), ramlBytes...)

	// write the main .raml file
	if err := ioutil.WriteFile(filepath.Join(dir, "api.raml"), ramlBytes, 0777); err != nil {
		return err
	}

	// copy all libraries files
	return copyLibrariesFiles(apiDef.Uses, apiDef.Libraries, ramlFile, dir)
}

// copy all library files to apidocs directory
func copyLibrariesFiles(uses map[string]string, libraries map[string]*raml.Library, ramlFile, dir string) error {
	baseDir := filepath.Dir(ramlFile)
	// copy library files
	for _, path := range uses {
		if err := copyFile(filepath.Join(baseDir, path), filepath.Join(dir, path)); err != nil {
			return err
		}
	}

	// do it recursively
	for _, l := range libraries {
		if err := copyLibrariesFiles(l.Uses, l.Libraries, filepath.Join(baseDir, l.Filename), filepath.Join(dir, filepath.Dir(l.Filename))); err != nil {
			return err
		}
	}
	return nil
}

// copy file from source to dest
func copyFile(source, dest string) error {
	// source file
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	// create target dir if needed
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}

	// creaate dest file
	destFile, err := os.OpenFile(dest, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// Extract extracts API docs data to specified path
func extract(dir string) error {

	// create dir if needed
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	return unzip("apidocs_html.zip", dir)
}

func unzip(zipFile, dir string) error {
	// get zip file from go-bindata asset
	b, err := Asset(zipFile)
	if err != nil {
		return err
	}

	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return err
	}

	// Iterate through the files in the archive
	// and write it to specified directory
	for _, f := range r.File {
		path := filepath.Join(dir, f.Name)

		// if file is a dir, simply create it
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
			continue
		}

		// open/create result file
		cpFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, f.Mode())
		if err != nil {
			return err
		}
		defer cpFile.Close()

		// open zipped file
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		if _, err := io.Copy(cpFile, rc); err != nil {
			return err
		}
	}
	return nil
}
