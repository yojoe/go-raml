package python

import (
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
)

type sanicRouteView struct {
	Name     string
	Endpoint string
	Methods  []serverMethod
}

func newSanicRouteView(endpoint string) sanicRouteView {
	// TODO : need better way to generate name
	name := strings.TrimSuffix(endpoint, "/")
	name = strings.TrimPrefix(name, "/")
	name = strings.TrimSuffix(name, ">")
	name = strings.Replace(name, "/<", "_by", -1)
	name = strings.Replace(name, ">/", "_", -1)
	name = strings.Replace(name, "/", "_", -1)
	name = strings.Replace(name, ">", "", -1)

	return sanicRouteView{
		Name:     name,
		Endpoint: endpoint,
	}
}
func (s SanicServer) generateResources(dir string) error {
	if err := s.generateResourcesAPIImpl(dir); err != nil {
		return err
	}
	if err := s.generateResourcesAPIIface(dir); err != nil {
		return err
	}
	return nil
}

func (s SanicServer) generateResourcesAPIIface(dir string) error {
	// routing & middlewares
	for _, pr := range s.ResourcesDef {
		// create sanicRouteView objects
		srvMap := map[string]sanicRouteView{}
		for _, m := range pr.Methods {
			pm := m.(serverMethod)

			srv, ok := srvMap[pm.Endpoint]
			if !ok {
				srv = newSanicRouteView(pm.Endpoint)
			}

			srv.Methods = append(srv.Methods, pm)
			srvMap[pm.Endpoint] = srv
		}

		// generate
		filename := filepath.Join(dir, strings.ToLower(pr.Name)+"_if.py")
		ctx := map[string]interface{}{
			"Name":      strings.ToLower(pr.Name),
			"RouteView": srvMap,
			"PR":        pr,
		}
		if err := commons.GenerateFile(ctx, "./templates/server_resources_if_python_sanic.tmpl", "server_resources_if_python_sanic", filename, true); err != nil {
			return err
		}
	}

	return nil
}

func (s SanicServer) generateResourcesAPIImpl(dir string) error {
	for _, pr := range s.ResourcesDef {
		filename := filepath.Join(dir, strings.ToLower(pr.Name)+"_api.py")
		if err := pr.generate(filename, "./templates/server_resources_api_python_sanic.tmpl", "server_resources_api_python_sanic", dir); err != nil {
			return err
		}
	}
	return nil
}
