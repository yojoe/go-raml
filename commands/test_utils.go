package commands

import (
	"os"
)

func cleanTestingDir() {
	os.RemoveAll("./test")
}
