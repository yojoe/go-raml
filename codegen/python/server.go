package python

import (
	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/raml"
)

var (
	globLibRootURLs []string
)

const (
	serverKindSanic       = "sanic"
	serverKindFlask       = "flask"
	serverKindGeventFlask = "gevent-flask"
)

// Server represents a python server
type Server interface {
	Generate(dir string) error
}

// NewServer creates a new python server
func NewServer(kind string, apiDef *raml.APIDefinition, apiDocsDir string,
	withMain bool, libRootURLs []string) Server {
	switch kind {
	case "", serverKindFlask:
		return NewFlaskServer(apiDef, apiDocsDir, withMain, libRootURLs, false)
	case serverKindGeventFlask:
		return NewFlaskServer(apiDef, apiDocsDir, true, libRootURLs, true)
	case serverKindSanic:
		return NewSanicServer(apiDef, apiDocsDir, withMain, libRootURLs)
	default:
		log.Fatalf("Invalid kind of python server : %v", kind)
		return nil
	}
}
