package libraries

import (
	"strings"
)

// StripLibRootURL strips root URL from library path
func StripLibRootURL(libFileName string, libRootURLs []string) string {
	if !strings.HasPrefix(libFileName, "http://") && !strings.HasPrefix(libFileName, "https://") {
		return libFileName
	}
	for _, root := range libRootURLs {
		if strings.HasPrefix(libFileName, root) {
			libFileName = strings.TrimPrefix(libFileName, root)
			return strings.TrimPrefix(libFileName, "/")
		}
	}
	return libFileName
}
