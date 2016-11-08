package capnp

import (
	log "github.com/Sirupsen/logrus"
	"os/exec"
	"strings"
)

func getID() string {
	b, err := exec.Command("capnp", []string{"id"}...).CombinedOutput()
	if err != nil {
		log.Fatalf("failed to generate capnp ID. please make sure you've installed `capnp`\nerr:%v\n", err)
	}
	return strings.TrimSpace(strings.Trim(string(b), "\n"))
}
