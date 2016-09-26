package codegen

import (
	"path/filepath"
	"strings"
)

// ClientService represents a root endpoint of an API
type ClientService struct {
	lang         string
	rootEndpoint string
	PackageName  string
	Methods      []methodInterface
}

// Name returns it's struct name
func (cs ClientService) Name() string {
	return strings.Title(cs.rootEndpoint[1:]) + "Service"
}

// EndpointName returns root endpoint name
func (cs ClientService) EndpointName() string {
	name := cs.rootEndpoint[1:]
	if cs.lang == langGo {
		name = strings.Title(name)
	}
	return name
}

// FilenameNoExt return filename without extension
func (cs ClientService) FilenameNoExt() string {
	return cs.rootEndpoint[1:] + "_service"
}
func (cs ClientService) filename(dir string) string {
	name := filepath.Join(dir, cs.FilenameNoExt())
	if cs.lang == langGo {
		return name + ".go"
	}
	return name + ".py"
}

// LibImportPaths returns all imported lib
func (cs ClientService) LibImportPaths() map[string]struct{} {
	ip := map[string]struct{}{}

	// methods
	for _, v := range cs.Methods {
		gm := v.(goClientMethod)
		for lib := range gm.libImported(globRootImportPath) {
			ip[lib] = struct{}{}
		}
	}
	return ip
}
