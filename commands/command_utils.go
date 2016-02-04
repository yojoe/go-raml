package commands

import (
	"errors"
	"os/exec"
)

const (
	cmdGenerateImport = "goimports"
	cmdFormatCode     = "gofmt"
)

var (
	availableGoTools = []string{cmdFormatCode}
)

var (
	errCmdGoToolsNotAvalilable = errors.New("Your Go tools command is not available")
)

func runGoTools(commandType, filePath string) error {
	valid := false
	for _, val := range availableGoTools {
		if val == commandType {
			valid = true
			break
		}
	}

	if !valid {
		return errCmdGoToolsNotAvalilable
	}

	args := []string{"-w", filePath}

	if err := exec.Command(commandType, args...).Run(); err != nil {
		return err
	}

	return nil
}
