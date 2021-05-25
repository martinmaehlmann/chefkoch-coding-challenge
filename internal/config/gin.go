package config

import (
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	ginPort = "GIN_PORT"
)

// GinConfig a struct containing all the configuration needed to run a gin-gonic server.
type GinConfig struct {
	Port int
}

// PrettyString returns a pretty formatted JSON representation of the struct.
func (g *GinConfig) PrettyString(indentation string) (string, error) {
	jsonString, err := json.MarshalIndent(g, "", indentation)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}

// NewGinConfig returns a new GinConfig.
func NewGinConfig() *GinConfig {
	return &GinConfig{Port: viper.GetInt(ginPort)}
}
