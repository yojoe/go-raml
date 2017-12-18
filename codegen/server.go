package codegen

import (
	"errors"
	"path/filepath"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen/apidocs"
	"github.com/Jumpscale/go-raml/codegen/golang"
	"github.com/Jumpscale/go-raml/codegen/nim"
	"github.com/Jumpscale/go-raml/codegen/python"
	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/codegen/tarantool"
)

var (
	errInvalidLang = errors.New("invalid language")
)

// Server represents server code
type Server struct {
	RAMLFile         string
	Dir              string // Destination directory
	PackageName      string
	Lang             string
	APIDocsDir       string // API Docs directory
	RootImportPath   string
	WithMain         bool   // true if we also generate main file
	Kind             string // currently only used by python client : sanic/flask
	APIFilePerMethod bool   // true if we want to generate one API file per API method
	LibRootURLs      []string
}

// Generate generates API server files
func (s *Server) Generate() error {
	apiDef := new(raml.APIDefinition)
	// parse the raml file
	dir, fileName := filepath.Split(s.RAMLFile)
	ramlBytes, err := raml.ParseReadFile(dir, fileName, apiDef)
	if err != nil {
		return err
	}

	switch s.Lang {
	case langGo:
		gs := golang.NewServer(apiDef, s.PackageName, s.APIDocsDir, s.RootImportPath, s.WithMain,
			s.APIFilePerMethod, s.Dir, s.LibRootURLs)
		err = gs.Generate()
	case langPython:
		ps := python.NewServer(s.Kind, apiDef, s.APIDocsDir, s.WithMain, s.LibRootURLs)
		err = ps.Generate(s.Dir)
	case langNim:
		ns := nim.NewServer(apiDef, s.APIDocsDir, s.Dir)
		err = ns.Generate()
	case langTarantool:
		ts := tarantool.NewServer(apiDef, s.APIDocsDir, s.Dir)
		err = ts.Generate()
	default:
		return errInvalidLang
	}
	if err != nil {
		return err
	}

	if s.APIDocsDir == "" {
		return nil
	}

	if s.Lang == langNim {
		s.APIDocsDir = "public/" + s.APIDocsDir
	}

	log.Infof("Generating API Docs to %v", s.APIDocsDir)

	return apidocs.Generate(apiDef, s.RAMLFile, ramlBytes, filepath.Join(s.Dir, s.APIDocsDir), s.LibRootURLs)
}
