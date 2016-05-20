package codegen

import (
	"path"
	"strings"

	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
)

// go representation of a security scheme
type pythonSecurity struct {
	*security
}

func (ps *pythonSecurity) generate(dir string) error {
	fileName := path.Join(dir, "oauth2_"+ps.Name+".py")
	return generateFile(ps, "./templates/oauth2_middleware_python.tmpl", "oauth2_middleware_python", fileName, false)
}

type pythonMiddleware struct {
	ImportPath string
	Name       string
	Args       string
}

func newPythonOauth2Middleware(ss raml.DefinitionChoice) (pythonMiddleware, error) {
	quotedScopes, err := getQuotedSecurityScopes(ss)
	if err != nil {
		return pythonMiddleware{}, err
	}

	importPath, name := pythonOauth2libImportPath(ss.Name)
	return pythonMiddleware{
		ImportPath: importPath,
		Name:       name,
		Args:       strings.Join(quotedScopes, ", "),
	}, nil
}

// get library import path from a type
func pythonOauth2libImportPath(typ string) (string, string) {
	typ = securitySchemeName(typ)
	// library use '.'
	if strings.Index(typ, ".") < 0 {
		return "oauth2_" + typ, "oauth2_" + typ
	}

	splitted := strings.Split(typ, ".")
	if len(splitted) != 2 {
		log.Fatalf("pythonOauth2libImportPath invalid security:" + typ)
	}
	// library name in the current document
	libName := splitted[0]

	// raml file of this lib
	libRAMLFile := globAPIDef.FindLibFile(denormalizePkgName(libName))

	if libRAMLFile == "" {
		log.Fatalf("pythonOauth2libImportPath() can't find library : %v", libName)
	}

	// relative lib package
	libPkg := libRelDir(libRAMLFile)

	return strings.Replace(normalizePkgName(libPkg), "/", ".", -1) + ".oauth2_" + splitted[1], "oauth2_" + splitted[1]
}
