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

// PostgresConfig a struct containing all the configuration needed to connect to a postgres database.
type PostgresConfig struct {
	hostname string
	database string
	sslMode  string
	timeZone string
	username string
	password string
	port     int
}

// DSN returns the dsn to connect against the database.
func (p *PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		p.hostname, p.username, p.password, p.database, p.port, p.sslMode, p.timeZone)
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
		hostname: viper.GetString(postgresHostname),
		database: viper.GetString(postgresDatabase),
		sslMode:  viper.GetString(postgresSslMode),
		timeZone: viper.GetString(postgresTimeZone),
		username: viper.GetString(postgresUsername),
		password: viper.GetString(postgresPassword),
		port:     viper.GetInt(postgresPort),
	}
}
