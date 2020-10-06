// Package config provides a configuration for the API
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/OrlovEvgeny/go-mcache"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/mrz1836/go-logger"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

// Values global configuration (config.Values)
var Values appConfig

// SchedulerConfig is our cron task wrapper
type SchedulerConfig struct {
	CronApp *cron.Cron
}

// AddJob adds a new cron job and returns the entry ID
func (s SchedulerConfig) AddJob(name, spec string, cmd func()) (entryID cron.EntryID, err error) {

	// No name or spec?
	if spec == "" || name == "" {
		err = fmt.Errorf("failed adding cron job %s, spec or name was empty", name)
		logger.Data(2, logger.ERROR, err.Error())
		return
	}

	// Add the cron job
	entryID, err = s.CronApp.AddFunc(spec, cmd)
	if err != nil {
		err = fmt.Errorf("error creating cron job %s spec: %s error: %s", name, spec, err.Error())
		logger.Data(2, logger.ERROR, err.Error())
	} else {
		logger.Data(2, logger.DEBUG, fmt.Sprintf("%s cron job added successfully, spec: [%s] entryID: [%d]", name, spec, entryID))
	}

	return
}

// RemoveJob will remove a cron job by entryID (int)
func (s SchedulerConfig) RemoveJob(entryID cron.EntryID) (err error) {
	s.CronApp.Remove(entryID)
	return
}

// Config constants used for optimization and value testing
const (
	DatabaseDefaultTxTimeout = 15 * time.Second
	EnvironmentDevelopment   = "development"
	EnvironmentKey           = "API_ENVIRONMENT"
	EnvironmentProduction    = "production"
	EnvironmentStaging       = "staging"
	HealthRequestPath        = "health"
	HTTPRequestReadTimeout   = 15 * time.Second
	HTTPRequestWriteTimeout  = 15 * time.Second
	ServiceModeAPI           = "api"
)

// appConfig is the configuration values and associated env vars
type appConfig struct {
	BasicAuth         basicAuthConfig `json:"basic_auth" mapstructure:"basic_auth"`
	Cache             cacheConfig     `json:"cache" mapstructure:"cache"`
	CacheEnabled      bool            `json:"-" mapstructure:"-"`
	DatabaseDebug     bool            `json:"database_debug" mapstructure:"database_debug"`
	DatabaseRead      databaseConfig  `json:"database_read" mapstructure:"database_read"`
	DatabaseWrite     databaseConfig  `json:"database_write" mapstructure:"database_write"`
	Email             emailConfig     `json:"email" mapstructure:"email"`
	Environment       string          `json:"environment" mapstructure:"environment"`
	Scheduler         SchedulerConfig `json:"-" mapstructure:"-"`
	ServerPort        string          `json:"server_port" mapstructure:"server_port"`
	ServiceMode       string          `json:"service_mode" mapstructure:"service_mode"`
	UnauthorizedError string          `json:"unauthorized_error" mapstructure:"unauthorized_error"`
}

// Validate checks the configuration for specific rules
func (a appConfig) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.BasicAuth),     // Runs validations on the child struct level
		validation.Field(&a.Cache),         // Runs validations on the child struct level
		validation.Field(&a.DatabaseRead),  // Runs validations on the child struct level
		validation.Field(&a.DatabaseWrite), // Runs validations on the child struct level
		validation.Field(&a.Environment, validation.Required, validation.In(EnvironmentDevelopment, EnvironmentStaging, EnvironmentProduction)),
		validation.Field(&a.ServerPort, validation.Required, is.Digit, validation.Length(2, 6)),
		validation.Field(&a.ServiceMode, validation.Required, validation.In(ServiceModeAPI)),
		validation.Field(&a.UnauthorizedError, validation.Required, validation.Length(2, 0)),
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
// Most of theses variables are for "redis" configuration
// MemStore is an internal storage mechanism for the local instance only
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
type cacheConfig struct {
	URL                   string              `json:"url" mapstructure:"url"`                                         // redis://localhost:6379
	MaxActiveConnections  int                 `json:"max_active_connections" mapstructure:"max_active_connections"`   // 0
	MaxConnectionLifetime int                 `json:"max_connection_lifetime" mapstructure:"max_connection_lifetime"` // 0
	MaxIdleConnections    int                 `json:"max_idle_connections" mapstructure:"max_idle_connections"`       // 10
	MaxIdleTimeout        int                 `json:"max_idle_timeout" mapstructure:"max_idle_timeout"`               // 240
	MemStore              *mcache.CacheDriver `json:"-" mapstructure:"-"`                                             // In-memory store (local box only)
	DependencyMode        bool                `json:"dependency_mode" mapstructure:"dependency_mode"`                 // false for digital ocean (not supported)
	UseTLS                bool                `json:"use_tls" mapstructure:"use_tls"`                                 // true for digital ocean (required)
}

// Validate checks the configuration for specific rules
func (c cacheConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.URL, validation.Length(0, 250)), // Testing
	)
}

// emailConfig is a configuration for a email services
type emailConfig struct {
	AwsSesAccessID      string `json:"aws_ses_access_id" mapstructure:"aws_ses_access_id"`         // 12345
	AwsSesSecretKey     string `json:"aws_ses_secret_key" mapstructure:"aws_ses_secret_key"`       // 12345
	FromDomain          string `json:"from_domain" mapstructure:"from_domain"`                     // example.com
	FromName            string `json:"from_name" mapstructure:"from_name"`                         // Test User
	FromUsername        string `json:"from_username" mapstructure:"from_username"`                 // testuser
	MandrillAPIKey      string `json:"mandrill_api_key" mapstructure:"mandrill_api_key"`           // 12345
	PostmarkServerToken string `json:"postmark_server_token" mapstructure:"postmark_server_token"` // 12345
	SMTPHost            string `json:"smtp_host" mapstructure:"smtp_host"`                         // example.com
	SMTPPassword        string `json:"smtp_password" mapstructure:"smtp_password"`                 // secret123
	SMTPPort            int    `json:"smtp_port" mapstructure:"smtp_port"`                         // 25
	SMTPUsername        string `json:"smtp_username" mapstructure:"smtp_username"`                 // testuser
}

// Validate checks the configuration for specific rules
func (e emailConfig) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.AwsSesAccessID, validation.Length(0, 100)),
		validation.Field(&e.AwsSesSecretKey, validation.Length(0, 100)),
		validation.Field(&e.FromDomain, validation.Required, validation.Length(1, 100)),
		validation.Field(&e.FromName, validation.Required, validation.Length(1, 100)),
		validation.Field(&e.FromUsername, validation.Required, validation.Length(1, 100)),
		validation.Field(&e.MandrillAPIKey, validation.Length(0, 100)),
		validation.Field(&e.PostmarkServerToken, validation.Length(0, 100)),
		validation.Field(&e.SMTPHost, validation.Length(0, 255)),
		validation.Field(&e.SMTPPassword, validation.Length(0, 255)),
		validation.Field(&e.SMTPUsername, validation.Length(0, 255)),
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

// Load all environment variables
func Load() (err error) {

	// Check the environment we are running
	environment := os.Getenv(EnvironmentKey)
	if len(environment) == 0 {
		logger.Data(2, logger.ERROR, "missing required environment var: "+EnvironmentKey)
		logger.Fatalln("exiting...")
	} else if environment != EnvironmentStaging && environment != EnvironmentProduction && environment != EnvironmentDevelopment {
		logger.Data(2, logger.ERROR, "invalid environment var: "+EnvironmentKey+" value: "+environment)
		logger.Fatalln("exiting...")
	}

	// Load configuration from json based on the environment
	viper.SetConfigFile(GetCurrentDir() + "/" + environment + ".json")

	// Set a replacer for replacing double underscore with nested period
	replacer := strings.NewReplacer(".", "__")
	viper.SetEnvKeyReplacer(replacer)

	// Set the prefix
	viper.SetEnvPrefix(ServiceModeAPI)

	// Use env vars
	viper.AutomaticEnv()

	// Read the configuration
	if err = viper.ReadInConfig(); err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("error reading %s configuration: %s", environment, err.Error()))
		return
	}

	logger.Data(2, logger.INFO, environment+" configuration env file processed")

	// Unmarshal into values struct
	if err = viper.Unmarshal(&Values); err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("error in unmarshal into values: %s", err.Error()))
		return
	}

	// Validate the configuration file
	if err = Values.Validate(); err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("error in %s configuration validation: %s", environment, err.Error()))
	}

	// Check service mode
	if Values.ServiceMode != ServiceModeAPI {
		logger.Data(2, logger.ERROR, "invalid value for service mode")
		logger.Fatalln("exiting...")
	}

	// Load the scheduler and start
	Values.Scheduler.CronApp = cron.New()
	Values.Scheduler.CronApp.Start()

	// Load the in-memory cache store
	// Used for storing values in local memory (with TTL)
	Values.Cache.MemStore = mcache.New()

	return
}

// GetCurrentDir gets the current directory for all operating systems
func GetCurrentDir() string {
	// Get the current path
	_, path, _, ok := runtime.Caller(1)

	// Return the file path
	if ok {
		return filepath.Dir(path)
	}
	return ""
}
