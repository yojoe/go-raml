package codegen

import (
	"path"
)

type goSecurity struct {
	*security
}

// generate Go representation of a security scheme
// it implemented as struct based middleware
func (gs *goSecurity) generate(dir string) error {
	fileName := path.Join(dir, "oauth2_"+gs.Name+"_middleware.go")
	return generateFile(gs, "./templates/oauth2_middleware.tmpl", "oauth2_middleware", fileName, false)
}
