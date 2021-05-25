package config

import (
	"encoding/json"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
)

func TestNewPostgresConfig(t *testing.T) {
	defer postgresVipertCleanup()
	setPostgresViperValues()

	actual := NewPostgresConfig()
	expected := &PostgresConfig{
		Hostname: "test",
		Database: "test",
		SSLMode:  "test",
		TimeZone: "test",
		Username: "test",
		Password: "test",
		Port:     8080,
	}

	assert.Equal(t, expected, actual)
}

func TestPostgresConfig_DSN(t *testing.T) {
	defer postgresVipertCleanup()
	setPostgresViperValues()

	actual := NewPostgresConfig()
	expected := "host=test user=test password=test dbname=test port=8080 sslmode=test TimeZone=test"

	assert.Equal(t, expected, actual.DSN())
}

func TestPostgresConfig_Dialector(t *testing.T) {
	defer postgresVipertCleanup()
	setPostgresViperValues()

	actual := NewPostgresConfig()
	expected := postgres.Open("host=test user=test password=test dbname=test port=8080 sslmode=test TimeZone=test")

	assert.Equal(t, expected, actual.Dialector())
}

func TestPostgresConfig_PrettyString(t *testing.T) {
	defer postgresVipertCleanup()
	setPostgresViperValues()

	postgresConfig := NewPostgresConfig()

	actual, err := postgresConfig.PrettyString("  ")
	assert.NoError(t, err)

	expected, err := json.MarshalIndent(postgresConfig, "", "  ")
	assert.NoError(t, err)

	assert.Equal(t, string(expected), actual)
}

func setPostgresViperValues() {
	viper.Set(postgresHostname, "test")
	viper.Set(postgresDatabase, "test")
	viper.Set(postgresSslMode, "test")
	viper.Set(postgresTimeZone, "test")
	viper.Set(postgresUsername, "test")
	viper.Set(postgresPassword, "test")
	viper.Set(postgresPort, 8080)
}

func postgresVipertCleanup() {
	viper.Reset()
}
