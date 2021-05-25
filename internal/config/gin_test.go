package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGinConfig_PrettyString(t *testing.T) {
	setGinViperValues()
	defer ginViperCleanup()

	ginConfig := NewGinConfig()

	actual, err := ginConfig.PrettyString("  ")
	assert.NoError(t, err)

	expected, err := json.MarshalIndent(ginConfig, "", "  ")
	assert.NoError(t, err)

	assert.Equal(t, string(expected), actual)
}

func TestNewGinConfig(t *testing.T) {
	setGinViperValues()
	defer ginViperCleanup()

	actual := &GinConfig{Port: 8080}
	expected := NewGinConfig()

	assert.Equal(t, expected, actual)
}

func setGinViperValues() {
	viper.Set(ginPort, 8080)
}

func ginViperCleanup() {
	viper.Reset()
}
