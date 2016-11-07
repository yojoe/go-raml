package capnp

import (
	"os/exec"
	"strings"
)

func getID() (string, error) {
	b, err := exec.Command("capnp", []string{"id"}...).CombinedOutput()
	return strings.TrimSpace(strings.Trim(string(b), "\n")), err
}
