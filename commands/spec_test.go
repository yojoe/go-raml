package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecGeneration(t *testing.T) {
	cmd := SpecCommand{}
	err := cmd.Execute()
	assert.NoError(t, err)
}
