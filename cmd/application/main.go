// Package main is the main application for the People Finder Service
package main

import (
	"net/http"

	"github.com/mrz1836/go-api/api"
	"github.com/mrz1836/go-api/config"
	"github.com/mrz1836/go-api/database"
	"github.com/mrz1836/go-cache"
	"github.com/mrz1836/go-logger"
)

// main application method
func main() {

	// Load the configuration and services
	config.Load()

	// Defer any connections
	defer func() {

		// Close the cache connection on exit
		if config.CacheEnabled {
			cache.Disconnect()
		}

		// Close the database on exit
		database.CloseAllConnections()
	}()

	// Load the server
	logger.Data(2, logger.DEBUG, "starting Go API server...", logger.MakeParameter("port", config.ServerPort))
	logger.Fatalln(http.ListenAndServe(":"+config.ServerPort, api.Handlers()))
}
