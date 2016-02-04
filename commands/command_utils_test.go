package commands

import (
	"fmt"
	"os"
	"testing"
)

func TestGoToolsCommand(t *testing.T) {
	fmt.Printf("Test: Go Tools Command ... \n")

	fmt.Printf("Setup test ...\n")
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current dir, err = %v \n", err)
	}
	fmt.Printf("Setup passed ...\n")

	fmt.Printf("Phase 1 : Testing gofmt ...\n")
	if err := runGoTools(cmdGenerateImport, dir+"/command_utils.go"); err != nil {
		t.Fatalf("Failed when run gofmt, err = %v \n", err)
	}
	fmt.Printf("Phase 1 passed ... \n")

	fmt.Printf("Phase 2 : Testing goimports ... \n")
	if err := runGoTools(cmdFormatCode, dir+"/command_utils.go"); err != nil {
		t.Fatalf("Failed when run gofmt, err = %v \n", err)
	}
	fmt.Printf("Phase 2 passed ... \n")

	fmt.Printf("Test: Go Tools Command Passed \n")
}
