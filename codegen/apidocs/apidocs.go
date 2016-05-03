package apidocs

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Generate generates API docs using api-console
// https://github.com/mulesoft/api-console
func Generate(ramlBytes []byte, dir string) error {
	// extract zipped files
	if err := extract(dir); err != nil {
		return err
	}
	ramlBytes = append([]byte("#%RAML 1.0\n"), ramlBytes...)
	// write the .raml file
	return ioutil.WriteFile(filepath.Join(dir, "api.raml"), ramlBytes, 0777)
}

// Extract extracts API docs data to specified path
func extract(dir string) error {

	// create dir if needed
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0777); err != nil {
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
