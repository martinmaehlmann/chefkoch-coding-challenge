package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRegistry(t *testing.T) {
	setGinViperValues(t)
	setPostgresViperValues(t)

	defer ginViperCleanup(t)

	actual := NewRegistry(NewGinConfig(), NewPostgresConfig(), nil)

	expected := &Registry{
		logger:    nil,
		GinConfig: &GinConfig{Port: 8080},
		PostgresConfig: &PostgresConfig{
			hostname: "test",
			database: "test",
			sslMode:  "test",
			timeZone: "test",
			username: "test",
			password: "test",
			port:     8080,
		},
	}

	assert.Equal(t, expected, actual)
}

func TestRegistry_PrettyString(t *testing.T) {
	setGinViperValues(t)
	setPostgresViperValues(t)

	defer ginViperCleanup(t)

	registry := NewRegistry(NewGinConfig(), NewPostgresConfig(), nil)

	actual, err := registry.PrettyString("  ")
	assert.NoError(t, err)

	expected, err := json.MarshalIndent(registry, "", "  ")
	assert.NoError(t, err)

	assert.Equal(t, string(expected), actual)
}
