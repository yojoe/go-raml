package libraries

import (
	"path/filepath"
	"strings"
)

// IsRemote returns true if this library is on the remote url
func IsRemote(path string) bool {
	return strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://")
}

// StripLibRootURL strips root URL from library path
func StripLibRootURL(libFile string, libRootURLs []string) string {
	if !IsRemote(libFile) {
		return libFile
	}
	for _, root := range libRootURLs {
		if strings.HasPrefix(libFile, root) {
			libFile = strings.TrimPrefix(libFile, root)
			libFile = strings.TrimPrefix(libFile, "/")
			break
		}
	}
	return libFile
}

func JoinPath(libDir, libFile string, libRootURLs []string) string {
	if IsRemote(libFile) {
		return StripLibRootURL(libFile, libRootURLs)
	}
	return filepath.Join(libDir, libFile)
}
