// Package main is the main application
package main

import (
	"net/http"

	"github.com/mrz1836/go-api/api"
	"github.com/mrz1836/go-api/config"
	"github.com/mrz1836/go-api/database"
	"github.com/mrz1836/go-cache"
	"github.com/mrz1836/go-logger"
	"github.com/volatiletech/sqlboiler/boil"
)

// main application method
func main() {

	var err error

	// Load the configuration and services
	if err = config.Load(); err != nil {
		logger.Fatalln("fatal error loading config:", err.Error())
	}

	// Load the services and connections
	if err = loadAPI(); err != nil {
		logger.Fatalln("fatal error loading api:", err.Error())
	}

	// Defer any connections
	defer func() {

		// Close the cache connection on exit
		if config.Values.CacheEnabled {
			cache.Disconnect()
		}

		// Close the database on exit
		database.CloseAllConnections()
	}()

	// Load the server
	logger.Data(2, logger.DEBUG, "starting Go API server...", logger.MakeParameter("port", config.Values.ServerPort))
	logger.Fatalln(http.ListenAndServe(":"+config.Values.ServerPort, api.Handlers()))
}

// loadAPI loads all the required services and connections
func loadAPI() (err error) {

	// Check the environment and use caching if set
	if len(config.Values.Cache.URL) > 0 {

		// Attempt to connect to the cache (redis)
		err = cache.Connect(config.Values.Cache.URL, config.Values.Cache.MaxActiveConnections, config.Values.Cache.MaxIdleConnections, config.Values.Cache.MaxConnectionLifetime, config.Values.Cache.MaxIdleTimeout)
		if err != nil {
			logger.Data(2, logger.ERROR, "failed to enable cache: "+err.Error())
			return
		} else {
			config.Values.CacheEnabled = true
			logger.Data(2, logger.INFO, "cache enabled at: "+config.Values.Cache.URL)
		}
	} else {
		logger.Data(2, logger.INFO, "caching: disabled")
	}

	// Turn on database debugging
	if config.Values.DatabaseDebug {
		boil.DebugMode = config.Values.DatabaseDebug
	}

	// Load the database configuration
	database.SetConfiguration(database.Configuration{
		DatabaseRead:  database.ConnectionConfig(config.Values.DatabaseRead),
		DatabaseWrite: database.ConnectionConfig(config.Values.DatabaseWrite),
	})

	// Open the connections
	err = database.OpenConnection()

	return
}
