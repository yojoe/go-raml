package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerGeneration(t *testing.T) {
	cmd := ServerCommand{}
	err := cmd.Execute()
	assert.NoError(t, err)
}
