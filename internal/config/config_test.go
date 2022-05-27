package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_valid(t *testing.T) {
	assert := require.New(t)

	cfg, err := NewConfig()
	assert.NoError(err)

	assert.Equal("8080", cfg.Addr)
}
