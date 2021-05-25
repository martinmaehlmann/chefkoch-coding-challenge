package config

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
)

const (
	postgresHostname = "POSTGRES_HOSTNAME"
	postgresDatabase = "POSTGRES_DATABASE"
	postgresSslMode  = "POSTGRES_SSL_MODE"
	postgresTimeZone = "POSTGRES_TIME_ZONE"
	postgresUsername = "POSTGRES_USERNAME"
	// nolint:gosec // not a hardcoded password
	postgresPassword = "POSTGRES_PASSWORD"
	postgresPort     = "POSTGRES_PORT"
)

// PostgresConfig a struct conaining all the configuration needed to connect to a postgres database.
type PostgresConfig struct {
	Hostname string
	Database string
	SSLMode  string
	TimeZone string
	Username string
	Password string
	Port     int
}

// DSN returns the dsn to connecto against the database.
func (p *PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		p.Hostname, p.Username, p.Password, p.Database, p.Port, p.SSLMode, p.TimeZone)
}

// Dialector returns a gorm postgres dialector.
func (p *PostgresConfig) Dialector() gorm.Dialector {
	return postgres.Open(p.DSN())
}

// PrettyString returns a pretty formatted JSON representation of the struct.
func (p *PostgresConfig) PrettyString(indentation string) (string, error) {
	jsonString, err := json.MarshalIndent(p, "", indentation)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}

// NewPostgresConfig returns a new PostgresConfig.
func NewPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		Hostname: viper.GetString(postgresHostname),
		Database: viper.GetString(postgresDatabase),
		SSLMode:  viper.GetString(postgresSslMode),
		TimeZone: viper.GetString(postgresTimeZone),
		Username: viper.GetString(postgresUsername),
		Password: viper.GetString(postgresPassword),
		Port:     viper.GetInt(postgresPort),
	}
}
