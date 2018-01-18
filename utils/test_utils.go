package utils

import (
	"io/ioutil"
	"strings"
)

func TestLoadFile(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// remove capnp ID from a file, we need it for the test.
// because capnp ID will always produce different value
// this func is not elegant
func removeID(s string) string {
	splt := strings.Split(s, "\n")
	clean := []string{}
	for i, v := range splt {
		if strings.HasPrefix(v, "@0x") {
			clean = append(clean, splt[i+1:]...)
			break
		}
		clean = append(clean, v)
	}
	return strings.Join(clean, "\n")
}

func TestLoadFileRemoveID(filename string) (string, error) {
	file, err := TestLoadFile(filename)
	if err != nil {
		return "", err
	}
	return removeID(file), nil
}
