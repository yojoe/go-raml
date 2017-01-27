package python

import (
	"path"

	"github.com/Jumpscale/go-raml/raml"
)

func (s SanicServer) generateOauth2(schemes map[string]raml.SecurityScheme, dir string) error {
	for _, sd := range getOauth2Defs(schemes) {
		fileName := path.Join(dir, "oauth2_"+sd.Name+".py")
		if err := sd.generate(fileName, "./templates/oauth2_middleware_python_sanic.tmpl", "oauth2_middleware_python_sanic"); err != nil {
			return err
		}
	}
	return nil
}
