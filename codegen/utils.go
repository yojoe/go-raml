package codegen

import (
	"os"
	"regexp"
	"strings"
)

var (
	regNonAlphanum = regexp.MustCompile("[^A-Za-z0-9]+")
)

// create directory if not exist
func checkCreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0777)
	}
	return nil
}

// replace non alphanumerics with "_"
func replaceNonAlphanumerics(s string) string {
	return strings.Trim(regNonAlphanum.ReplaceAllString(s, "_"), "_")
}

func displayNameToFuncName(str string) string {
	str = strings.Replace(str, " ", "", -1) // remove the space
	return replaceNonAlphanumerics(str)     // change the other to _
}
