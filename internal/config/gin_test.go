package config

import (
	"encoding/json"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGinConfig_PrettyString(t *testing.T) {
	defer ginViperCleanup(t)
	setGinViperValues(t)

	ginConfig := NewGinConfig()

	actual, err := ginConfig.PrettyString("  ")
	assert.NoError(t, err)

	expected, err := json.MarshalIndent(ginConfig, "", "  ")
	assert.NoError(t, err)

	assert.Equal(t, string(expected), actual)
}

func TestNewGinConfig(t *testing.T) {
	defer ginViperCleanup(t)
	setGinViperValues(t)

	actual := &GinConfig{Port: 8080}
	expected := NewGinConfig()

	assert.Equal(t, expected, actual)
}

func setGinViperValues(t *testing.T) {
	t.Helper()
	viper.Set(ginPort, 8080)
}

func ginViperCleanup(t *testing.T) {
	t.Helper()
	viper.Reset()
}
