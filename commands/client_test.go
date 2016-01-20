package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientGeneration(t *testing.T) {
	cmd := ClientCommand{}
	err := cmd.Execute()
	assert.NoError(t, err)
}
