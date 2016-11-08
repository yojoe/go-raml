package capnp

import (
	"fmt"
)

func (s *Struct) goImports() []string {
	imports := []string{`using Go = import "/go.capnp"`}

	// import enum
	for _, f := range s.Fields {
		if f.Enum != nil {
			imports = append(imports, fmt.Sprintf(`using import "%v.capnp".%v`, f.Enum.Name, f.Enum.Name))
		}
	}

	// import non buitin types
	for _, f := range s.Fields {
		if typesRegistered(f.Type) {
			imports = append(imports, fmt.Sprintf(`using import "%v.capnp".%v`, f.Type, f.Type))
		}
	}
	return imports
}

func (s *Struct) goAnnotations() []string {
	annos := []string{fmt.Sprintf(`$Go.package("%v")`, s.pkg)}
	// import enums
	for _, f := range s.Fields {
		if f.Enum != nil {
			annos = append(annos, fmt.Sprintf(`$Go.import("%v")`, f.Enum.pkg))
		}
	}
	// import our own package
	annos = append(annos, fmt.Sprintf(`$Go.import("%v")`, s.pkg))
	return annos
}
