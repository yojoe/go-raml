package python

import (
	"path/filepath"
	"sort"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/codegen/resource"
)

type service struct {
	resource.ClientService
	Methods            []clientMethod
	UnmarshallResponse bool
}

func newService(endpoint string, methods []resource.Method, unmarshallResponse bool) *service {
	clientMethods := make([]clientMethod, 0, len(methods))
	for _, rm := range methods {
		cm := newClientMethod(rm)
		clientMethods = append(clientMethods, cm)
	}

	return &service{
		ClientService:      resource.NewClientService(endpoint),
		Methods:            clientMethods,
		UnmarshallResponse: unmarshallResponse,
	}
}

func (s service) Imports() []string {
	if !s.UnmarshallResponse {
		return nil
	}

	ipMap := map[string]struct{}{
		"from .unmarshall_error import UnmarshallError":      struct{}{},
		"from .unhandled_api_error import UnhandledAPIError": struct{}{},
	}
	for _, m := range s.Methods {
		for _, importLine := range m.imports() {
			ipMap[importLine] = struct{}{}
		}
	}
	return sortImportPaths(ipMap)
}
func (s service) filename(dir string) string {
	return filepath.Join(dir, s.FilenameNoExt) + ".py"
}

type byEndoint []clientMethod

func (b byEndoint) Len() int      { return len(b) }
func (b byEndoint) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byEndoint) Less(i, j int) bool {
	return b[i].Endpoint+b[i].Verb() < b[j].Endpoint+b[j].Verb()
}

func (s *service) generate(tmpl clientTemplate, dir string) error {
	sort.Sort(byEndoint(s.Methods))
	return commons.GenerateFile(s, tmpl.serviceFile, tmpl.serviceName, s.filename(dir), true)
}

func sortImportPaths(ip map[string]struct{}) []string {
	libs := []string{}
	for k := range ip {
		libs = append(libs, k)
	}
	sort.Strings(libs)
	return libs
}
