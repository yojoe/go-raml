package golang

import (
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/resource"
)

// ClientService represents a root endpoint of an API
type ClientService struct {
	resource.ClientService
	PackageName string
	Methods     []clientMethod
}

func newClientService(rootEndpoint, packageName string, resMethods []resource.Method) *ClientService {
	var methods []clientMethod

	for _, rm := range resMethods {
		clientMeth := newClientMethod(rm)
		methods = append(methods, clientMeth)
	}

	cs := resource.NewClientService(rootEndpoint)
	cs.EndpointName = strings.Title(cs.EndpointName)
	return &ClientService{
		ClientService: cs,
		PackageName:   packageName,
		Methods:       methods,
	}
}

func (cs ClientService) filename(dir string) string {
	name := filepath.Join(dir, cs.FilenameNoExt)
	return name + ".go"
}

// NeedImportJSON returns true if this service need
// to import encoding/json
func (cs ClientService) NeedImportJSON() bool {
	for _, cm := range cs.Methods {
		if cm.needImportEncodingJSON() {
			return true
		}
	}
	return false
}

// LibImportPaths returns all imported lib
func (cs ClientService) LibImportPaths() map[string]struct{} {
	ip := map[string]struct{}{}

	var needImportGoraml bool
	// methods
	for _, gm := range cs.Methods {
		//gm := v.(clientMethod)
		for lib := range gm.libImported(globRootImportPath) {
			ip[lib] = struct{}{}
		}
		if !needImportGoraml {
			needImportGoraml = gm.needImportGoraml()
		}
	}
	if needImportGoraml {
		ip[joinImportPath(globRootImportPath, "goraml")] = struct{}{}
	}
	return ip
}
