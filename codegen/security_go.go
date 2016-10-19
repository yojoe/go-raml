package codegen

import (
	"path"

	"github.com/Jumpscale/go-raml/codegen/commons"
)

type goSecurity struct {
	*security
}

// generate Go representation of a security scheme
// it implemented as struct based middleware
func (gs *goSecurity) generate(dir string) error {
	fileName := path.Join(dir, "oauth2_"+gs.Name+"_middleware.go")
	return commons.GenerateFile(gs, "./templates/oauth2_middleware.tmpl", "oauth2_middleware", fileName, false)
}
