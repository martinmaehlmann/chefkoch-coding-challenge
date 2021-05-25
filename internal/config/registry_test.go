package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRegistry(t *testing.T) {
	setGinViperValues()
	setPostgresViperValues()

	defer ginViperCleanup()

	actual := NewRegistry(NewGinConfig(), NewPostgresConfig(), nil)

	expected := &Registry{
		logger:    nil,
		GinConfig: &GinConfig{Port: 8080},
		PostgresConfig: &PostgresConfig{
			Hostname: "test",
			Database: "test",
			SSLMode:  "test",
			TimeZone: "test",
			Username: "test",
			Password: "test",
			Port:     8080,
		},
	}

	assert.Equal(t, expected, actual)
}

func TestRegistry_PrettyString(t *testing.T) {
	setGinViperValues()
	setPostgresViperValues()

	defer ginViperCleanup()

	registry := NewRegistry(NewGinConfig(), NewPostgresConfig(), nil)

	actual, err := registry.PrettyString("  ")
	assert.NoError(t, err)

	expected, err := json.MarshalIndent(registry, "", "  ")
	assert.NoError(t, err)

	assert.Equal(t, string(expected), actual)
}
