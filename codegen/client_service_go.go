package codegen

import (
	"path/filepath"
	"strings"
)

// ClientService represents a root endpoint of an API
type ClientService struct {
	rootEndpoint string
	PackageName  string
	Methods      []methodInterface
}

// Name returns it's struct name
func (cs ClientService) Name() string {
	return strings.Title(cs.rootEndpoint[1:]) + "Service"
}

func (cs ClientService) EndpointName() string {
	return strings.Title(cs.rootEndpoint[1:])
}

func (cs ClientService) filename(dir string) string {
	return filepath.Join(dir, cs.rootEndpoint[1:]+"_service.go")
}
