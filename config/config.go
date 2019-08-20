// Package config provides a configuration for the API
package config

import (
	"fmt"
	"os"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/mrz1836/go-api/database"
	"github.com/mrz1836/go-cache"
	"github.com/mrz1836/go-logger"
	"github.com/spf13/viper"
)

// Global configuration (config.Values)
var Values appConfig

// appConfig are the configuration values and env vars
type appConfig struct {
	CacheEnabled                    bool   `env:"-" json:"-" mapstructure:"-"`
	CacheMaxActiveConnections       int    `env:"API_CACHE_MAX_ACTIVE_CONNECTIONS" json:"cache_max_active_connections" mapstructure:"cache_max_active_connections"`    // (0 is unlimited)
	CacheMaxConnectionLifetime      int    `env:"API_CACHE_MAX_CONNECTION_LIFETIME" json:"cache_max_connection_lifetime" mapstructure:"cache_max_connection_lifetime"` // (0 is unlimited)
	CacheMaxIdleConnections         int    `env:"API_CACHE_MAX_IDLE_CONNECTIONS" json:"cache_max_idle_connections" mapstructure:"cache_max_idle_connections"`          // (0 is unlimited)
	CacheMaxIdleTimeout             int    `env:"API_CACHE_MAX_IDLE_TIMEOUT" json:"cache_max_idle_timeout" mapstructure:"cache_max_idle_timeout"`                      // (0 is unlimited)
	CacheURL                        string `env:"API_CACHE_URL" json:"cache_url" mapstructure:"cache_url"`
	DatabaseMaxConnTime             int    `env:"API_DATABASE_MAX_CONN_TIME" json:"database_max_conn_time" mapstructure:"database_max_conn_time"` // (0 is unlimited)
	DatabaseReadHost                string `env:"API_DATABASE_READ_HOST" json:"database_read_host" mapstructure:"database_read_host"`
	DatabaseReadMaxIdleConnections  int    `env:"API_DATABASE_READ_MAX_IDLE_CONNECTIONS" json:"database_read_max_idle_connections" mapstructure:"database_read_max_idle_connections"` // (0 is unlimited)
	DatabaseReadMaxOpenConnections  int    `env:"API_DATABASE_READ_MAX_OPEN_CONNECTIONS" json:"database_read_max_open_connections" mapstructure:"database_read_max_open_connections"` // (0 is unlimited)
	DatabaseReadName                string `env:"API_DATABASE_READ_NAME" json:"database_read_name" mapstructure:"database_read_name"`
	DatabaseReadPassword            string `env:"API_DATABASE_READ_PASSWORD" json:"database_read_password" mapstructure:"database_read_password"`
	DatabaseReadPort                string `env:"API_DATABASE_READ_PORT" json:"database_read_port" mapstructure:"database_read_port"`
	DatabaseReadUser                string `env:"API_DATABASE_READ_USER" json:"database_read_user" mapstructure:"database_read_user"`
	DatabaseWriteHost               string `env:"API_DATABASE_WRITE_HOST" json:"database_write_host" mapstructure:"database_write_host"`
	DatabaseWriteMaxIdleConnections int    `env:"API_DATABASE_WRITE_MAX_IDLE_CONNECTIONS" json:"database_write_max_idle_connections" mapstructure:"database_write_max_idle_connections"` // (0 is unlimited)
	DatabaseWriteMaxOpenConnections int    `env:"API_DATABASE_WRITE_MAX_OPEN_CONNECTIONS" json:"database_write_max_open_connections" mapstructure:"database_write_max_open_connections"` // (0 is unlimited)
	DatabaseWriteName               string `env:"API_DATABASE_WRITE_NAME" json:"database_write_name" mapstructure:"database_write_name"`
	DatabaseWritePassword           string `env:"API_DATABASE_WRITE_PASSWORD" json:"database_write_password" mapstructure:"database_write_password"`
	DatabaseWritePort               string `env:"API_DATABASE_WRITE_PORT" json:"database_write_port" mapstructure:"database_write_port"`
	DatabaseWriteUser               string `env:"API_DATABASE_WRITE_USER" json:"database_write_user" mapstructure:"database_write_user"`
	Environment                     string `env:"API_ENVIRONMENT" json:"environment" mapstructure:"environment"`
	ServerPort                      string `env:"API_SERVER_PORT" json:"server_port" mapstructure:"server_port"`
}

// Validate checks the configuration for specific rules
func (c appConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ServerPort, validation.Required, validation.Length(2, 6)),
		validation.Field(&c.DatabaseReadHost, validation.Required),
		validation.Field(&c.DatabaseReadName, validation.Required),
		validation.Field(&c.DatabaseReadPassword, validation.Required),
		validation.Field(&c.DatabaseReadUser, validation.Required),
		validation.Field(&c.DatabaseReadPort, validation.Required, validation.Length(2, 6)),
		validation.Field(&c.DatabaseWriteHost, validation.Required),
		validation.Field(&c.DatabaseWriteName, validation.Required),
		validation.Field(&c.DatabaseWritePassword, validation.Required),
		validation.Field(&c.DatabaseWriteUser, validation.Required),
		validation.Field(&c.DatabaseWritePort, validation.Required, validation.Length(2, 6)),
		validation.Field(&c.Environment, validation.Required, validation.In("development", "staging", "production")),
	)
}

// init load all environment variables
func Load() {

	// Check the environment we are running
	environment := os.Getenv("API_ENVIRONMENT")
	if len(environment) == 0 {
		logger.Data(2, logger.ERROR, "missing required environment var: API_ENVIRONMENT")
		logger.Fatalln("exiting...")
	}

	// Load configuration from json based on the environment
	if environment == "production" {
		viper.SetConfigFile("./config/production.json")
	} else if environment == "staging" {
		viper.SetConfigFile("./config/staging.json")
	} else {
		viper.SetConfigFile("./config/development.json")
	}

	// Set the prefix
	viper.SetEnvPrefix("api")

	// Use env vars
	viper.AutomaticEnv()

	// Read the configuration
	if err := viper.ReadInConfig(); err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("error reading env configuration: %s", err.Error()))
	} else {
		logger.Data(2, logger.INFO, environment+" configuration env file processed")
	}

	// Unmarshal into values struct
	if err := viper.Unmarshal(&Values); err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("error in unmarshal into values: %s", err.Error()))
	}

	// Validate the configuration file
	if err := Values.Validate(); err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("error in configuration validation: %s", err.Error()))
	}

	// Check the environment and use caching if set
	if len(Values.CacheURL) > 0 {

		// Attempt to connect to the cache (redis)
		err := cache.Connect(Values.CacheURL, Values.CacheMaxActiveConnections, Values.CacheMaxIdleConnections, Values.CacheMaxConnectionLifetime, Values.CacheMaxIdleTimeout)
		if err != nil {
			logger.Data(2, logger.ERROR, "failed to enable cache: "+err.Error())
		} else {
			Values.CacheEnabled = true
			logger.Data(2, logger.INFO, "cache enabled at: "+Values.CacheURL)
		}
	} else {
		logger.Data(2, logger.INFO, "caching: disabled")
	}

	// Load the database configuration
	dbConfig := database.Configuration{
		DatabaseDriver:                  database.MySQLDriver,
		DatabaseMaxConnLifetime:         Values.DatabaseMaxConnTime,
		DatabaseReadHost:                Values.DatabaseReadHost,
		DatabaseReadMaxIdleConnections:  Values.DatabaseReadMaxIdleConnections,
		DatabaseReadMaxOpenConnections:  Values.DatabaseReadMaxOpenConnections,
		DatabaseReadName:                Values.DatabaseReadName,
		DatabaseReadPort:                Values.DatabaseReadPort,
		DatabaseReadUser:                Values.DatabaseReadUser,
		DatabaseReadPassword:            Values.DatabaseReadPassword,
		DatabaseWriteHost:               Values.DatabaseWriteHost,
		DatabaseWriteMaxIdleConnections: Values.DatabaseWriteMaxIdleConnections,
		DatabaseWriteMaxOpenConnections: Values.DatabaseWriteMaxOpenConnections,
		DatabaseWriteName:               Values.DatabaseWriteName,
		DatabaseWritePort:               Values.DatabaseWritePort,
		DatabaseWriteUser:               Values.DatabaseWriteUser,
		DatabaseWritePassword:           Values.DatabaseWritePassword,
	}
	database.SetConfiguration(dbConfig)

	// Open the connections
	database.OpenConnection()
}
