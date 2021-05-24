package config

import (
	"encoding/json"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

// PrettyStringer interface that shall return a pretty formatted representation of a struct. Equivalent to Stringer.
type PrettyStringer interface {
	PrettyString(indentation string) (string, error)
}

// Dialector interface to return a gorm.Dialector.
type Dialector interface {
	DSN() string
	Dialector() gorm.Dialector
}

// Registry a registry containing all different configuration structs.
type Registry struct {
	logger         *zap.Logger
	GinConfig      *GinConfig
	PostgresConfig *PostgresConfig
}

// PrettyString returns a pretty formatted JSON representation of the struct.
func (r *Registry) PrettyString(indentation string) (string, error) {
	out, err := json.MarshalIndent(r, "", indentation)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

// NewRegistry returns a new Registry.
func NewRegistry(ginConfig *GinConfig, postgresConfig *PostgresConfig, logger *zap.Logger) *Registry {
	return &Registry{
		logger:         logger,
		GinConfig:      ginConfig,
		PostgresConfig: postgresConfig,
	}
}
