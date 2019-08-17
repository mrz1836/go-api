// Package config provides a configuration for the People Finder Service
package config

import (
	"os"

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
)

// init load all environment variables
func init() {

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
}
