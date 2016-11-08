package capnp

import (
	"fmt"
	"sort"
)

// schema for Go need additional data needed by Go compiler
// go compilation : capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std   -ogo *.capnp

func (s *Struct) goImports() []string {
	return []string{`using Go = import "/go.capnp"`}
}

func (s *Struct) goAnnotations() []string {
	pkg := []string{fmt.Sprintf(`$Go.package("%v")`, s.pkg)}
	annos := []string{}
	// import enums
	for _, f := range s.Fields {
		if f.Enum != nil {
			annos = append(annos, fmt.Sprintf(`$Go.import("%v")`, f.Enum.pkg))
		}
	}
	// import our own package
	annos = append(annos, fmt.Sprintf(`$Go.import("%v")`, s.pkg))
	sort.Strings(annos)
	return append(pkg, annos...)
}
