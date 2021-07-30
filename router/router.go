// Package router is all the restful handler/router endpoint and methods for the application
package router

import (
	"github.com/julienschmidt/httprouter"
	apirouter "github.com/mrz1836/go-api-router"
	"github.com/mrz1836/go-api/actions/api"
	"github.com/mrz1836/go-api/actions/persons"
	"github.com/mrz1836/go-api/config"
)

// Handlers isolated the handlers / router for API (helps with testing)
func Handlers() *httprouter.Router {

	// Create a new router
	r := apirouter.New()

	// Based on service mode
	if config.Values.ServiceMode == config.ServiceModeAPI {
		// r.CrossOriginAllowOriginAll = false
		// r.CrossOriginAllowOrigin = "*"

		// This is used for the "Origin" to be returned as the origin
		r.CrossOriginAllowOriginAll = true

		// Create a middleware stack:
		// s := apirouter.NewStack()

		// Use your middleware:
		// s.Use(passThrough)

		api.RegisterRoutes(r)
		persons.RegisterRoutes(r)

	} // else (another service mode?)

	// Return the router
	return r.HTTPRouter.Router
}
