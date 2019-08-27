// Package config provides a configuration for the API
package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/mrz1836/go-logger"
	"github.com/spf13/viper"
)

// Global configuration (config.Values)
var Values appConfig

// appConfig is the configuration values and associated env vars
type appConfig struct {
	BasicAuth         basicAuthConfig `json:"basic_auth" mapstructure:"basic_auth"`
	Cache             cacheConfig     `json:"cache" mapstructure:"cache"`
	CacheEnabled      bool            `json:"-" mapstructure:"-"`
	DatabaseDebug     bool            `json:"database_debug" mapstructure:"unable to update persons"`
	DatabaseRead      databaseConfig  `json:"database_read" mapstructure:"database_read"`
	DatabaseWrite     databaseConfig  `json:"database_write" mapstructure:"database_write"`
	Environment       string          `json:"environment" mapstructure:"environment"`
	ServerPort        string          `json:"server_port" mapstructure:"server_port"`
	UnauthorizedError string          `json:"unauthorized_error" mapstructure:"unauthorized_error"`
}

// Validate checks the configuration for specific rules
func (c appConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.BasicAuth),     // Runs validations on the child struct level
		validation.Field(&c.Cache),         // Runs validations on the child struct level
		validation.Field(&c.DatabaseRead),  // Runs validations on the child struct level
		validation.Field(&c.DatabaseWrite), // Runs validations on the child struct level
		validation.Field(&c.Environment, validation.Required, validation.In("development", "staging", "production")),
		validation.Field(&c.ServerPort, validation.Required, is.Digit, validation.Length(2, 6)),
		validation.Field(&c.UnauthorizedError, validation.Required, validation.Length(2, 0)),
	)
}

// databaseConfig is a configuration for a SQL connection
type databaseConfig struct {
	Driver             string `json:"driver" mapstructure:"driver"`                             // mysql or postgresql
	Host               string `json:"host" mapstructure:"host"`                                 // localhost
	MaxConnectionTime  int    `json:"max_connection_time" mapstructure:"max_connection_time"`   // 60
	MaxIdleConnections int    `json:"max_idle_connections" mapstructure:"max_idle_connections"` // 5
	MaxOpenConnections int    `json:"max_open_connections" mapstructure:"max_open_connections"` // 5
	Name               string `json:"name" mapstructure:"name"`                                 // database-name
	Password           string `json:"password" mapstructure:"password"`                         // user-password
	Port               string `json:"port" mapstructure:"port"`                                 // 3306
	User               string `json:"user" mapstructure:"user"`                                 // username
}

// Validate checks the configuration for specific rules
func (d databaseConfig) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Driver, validation.Required, validation.Length(3, 100)),
		validation.Field(&d.Host, validation.Required, validation.Length(3, 250)),
		validation.Field(&d.Name, validation.Required, validation.Length(2, 100)),
		validation.Field(&d.Password, validation.Required, validation.Length(3, 250)),
		validation.Field(&d.User, validation.Required, validation.Length(3, 250)),
		validation.Field(&d.Port, validation.Required, validation.Length(2, 6)),
	)
}

// cacheConfig is a configuration for a Redis connection
type cacheConfig struct {
	MaxActiveConnections  int    `json:"max_active_connections" mapstructure:"max_active_connections"`   // 0
	MaxConnectionLifetime int    `json:"max_connection_lifetime" mapstructure:"max_connection_lifetime"` // 0
	MaxIdleConnections    int    `json:"max_idle_connections" mapstructure:"max_idle_connections"`       // 10
	MaxIdleTimeout        int    `json:"max_idle_timeout" mapstructure:"max_idle_timeout"`               // 240
	URL                   string `json:"url" mapstructure:"url"`                                         // redis://localhost:6379
}

// Validate checks the configuration for specific rules
func (c cacheConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.URL, validation.Length(0, 250)), // Testing
	)
}

// basicAuthConfig is a basic HTTP auth user
type basicAuthConfig struct {
	Password string `json:"password" mapstructure:"password"` // pass876
	User     string `json:"user" mapstructure:"user"`         // john
}

// Validate checks the configuration for specific rules
func (b basicAuthConfig) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Password, validation.Required, validation.Length(3, 100)),
		validation.Field(&b.User, validation.Required, is.Alphanumeric, validation.Length(3, 100)),
	)
}

// init load all environment variables
func Load() (err error) {

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

	// Set a replacer for replacing double underscore with nested period
	replacer := strings.NewReplacer(".", "__")
	viper.SetEnvKeyReplacer(replacer)

	// Set the prefix
	viper.SetEnvPrefix("api")

	// Use env vars
	viper.AutomaticEnv()

	// Read the configuration
	if err = viper.ReadInConfig(); err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("error reading env configuration: %s", err.Error()))
		return
	} else {
		logger.Data(2, logger.INFO, environment+" configuration env file processed")
	}

	// Unmarshal into values struct
	if err = viper.Unmarshal(&Values); err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("error in unmarshal into values: %s", err.Error()))
		return
	}

	// Validate the configuration file
	if err = Values.Validate(); err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("error in configuration validation: %s", err.Error()))
	}

	return
}
