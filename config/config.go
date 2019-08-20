// Package config provides a configuration for the API
package config

import (
	"os"

	"github.com/mrz1836/go-api/database"
	"github.com/mrz1836/go-cache"
	"github.com/mrz1836/go-logger"
)

// Set the global configuration
var (
	// ServerPort is for the API server
	ServerPort = "3000"

	// CacheURL is the redis url connection string
	CacheURL = os.Getenv("CACHE_URL")

	// CacheEnabled is a flag for caching (uses redis) (env: REDIS_URL)
	CacheEnabled = false

	// CacheMaxActiveConnections is the max active connections (0 is unlimited)
	CacheMaxActiveConnections = 0

	// CacheMaxIdleConnections is the max idle connections (0 is unlimited)
	CacheMaxIdleConnections = 10

	// CacheMaxConnectionLifetime is the max time in a connection (0 is unlimited)
	CacheMaxConnectionLifetime = 0

	// CacheMaxIdleTimeout is the max idle time (0 is unlimited)
	CacheMaxIdleTimeout = 240

	// DatabaseMaxConnLifetime max connection time
	DatabaseMaxConnLifetime = 60

	// DatabaseReadHost is the read hostname
	DatabaseReadHost = "localhost"

	// DatabaseReadPort is the read port
	DatabaseReadPort = "3306"

	// DatabaseReadName is the name of the database
	DatabaseReadName = "api_example"

	// DatabaseReadUser is the username
	DatabaseReadUser = "apiDbTestUser"

	// DatabaseReadPassword is the password for the user
	DatabaseReadPassword = "ThisIsSecureEnough123"

	// DatabaseReadMaxIdleConnections max idle connections
	DatabaseReadMaxIdleConnections = 5

	// DatabaseReadMaxOpenConnections max open connection
	DatabaseReadMaxOpenConnections = 5

	// DatabaseWriteHost is the read hostname
	DatabaseWriteHost = "localhost"

	// DatabaseWritePort is the read port
	DatabaseWritePort = "3306"

	// DatabaseWriteName is the name of the database
	DatabaseWriteName = "api_example"

	// DatabaseWriteUser is the username
	DatabaseWriteUser = "apiDbTestUser"

	// DatabaseWritePassword is the password for the user
	DatabaseWritePassword = "ThisIsSecureEnough123"

	// DatabaseWriteMaxIdleConnections max idle connections
	DatabaseWriteMaxIdleConnections = 5

	// DatabaseWriteMaxOpenConnections max open connections
	DatabaseWriteMaxOpenConnections = 5
)

// init load all environment variables
func Load() {

	// Check the environment and use caching if set
	if len(CacheURL) > 0 {

		// Attempt to connect to the cache (redis)
		err := cache.Connect(CacheURL, CacheMaxActiveConnections, CacheMaxIdleConnections, CacheMaxConnectionLifetime, CacheMaxIdleTimeout)
		if err != nil {
			logger.Data(2, logger.ERROR, "failed to enable cache: "+err.Error())
		} else {
			CacheEnabled = true
			logger.Data(2, logger.INFO, "cache enabled at: "+CacheURL)
		}
	} else {
		logger.Data(2, logger.INFO, "caching: disabled")
	}

	// Load the database configuration
	dbConfig := database.Configuration{
		DatabaseDriver:                  "mysql",
		DatabaseMaxConnLifetime:         DatabaseMaxConnLifetime,
		DatabaseReadHost:                DatabaseReadHost,
		DatabaseReadMaxIdleConnections:  DatabaseReadMaxIdleConnections,
		DatabaseReadMaxOpenConnections:  DatabaseReadMaxOpenConnections,
		DatabaseReadName:                DatabaseReadName,
		DatabaseReadPort:                DatabaseReadPort,
		DatabaseReadUser:                DatabaseReadUser,
		DatabaseReadPassword:            DatabaseReadPassword,
		DatabaseWriteHost:               DatabaseWriteHost,
		DatabaseWriteMaxIdleConnections: DatabaseWriteMaxIdleConnections,
		DatabaseWriteMaxOpenConnections: DatabaseWriteMaxOpenConnections,
		DatabaseWriteName:               DatabaseWriteName,
		DatabaseWritePort:               DatabaseWritePort,
		DatabaseWriteUser:               DatabaseWriteUser,
		DatabaseWritePassword:           DatabaseWritePassword,
	}
	database.SetConfiguration(dbConfig)

	// Open the connections
	database.OpenConnection()
}
