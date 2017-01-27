package python

import (
	"path"

	"github.com/Jumpscale/go-raml/raml"
)

func (fs FlaskServer) generateOauth2(schemes map[string]raml.SecurityScheme, dir string) error {
	securityDefs := getOauth2Defs(schemes)

	for _, sd := range securityDefs {
		fileName := path.Join(dir, "oauth2_"+sd.Name+".py")
		if err := sd.generate(fileName, "./templates/oauth2_middleware_python.tmpl", "oauth2_middleware_python"); err != nil {
			return err
		}
	}
	return nil
}
