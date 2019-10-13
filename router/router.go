// Package router is all the restful handler/router endpoint and methods for the application
package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mrz1836/go-api-router"
	"github.com/mrz1836/go-api/actions/base_api"
	"github.com/mrz1836/go-api/actions/persons"
	"github.com/mrz1836/go-api/config"
)

// Handlers isolated the handlers / router for API (helps with testing)
func Handlers() *httprouter.Router {

	// Create a new router
	r := apirouter.New()

	// Based on application mode
	if config.Values.ApplicationMode == config.ApplicationModeAPI {
		r.CrossOriginAllowOriginAll = false
		r.CrossOriginAllowOrigin = "*"

		// Create a middleware stack:
		//s := apirouter.NewStack()

		// Use your middleware:
		//s.Use(passThrough)

		baseApi.RegisterRoutes(r)
		persons.RegisterRoutes(r)

	} else {

		// Another service
		//baseAnotherService.RegisterRoutes(r)
	}

	// Return the router
	return r.HTTPRouter
}
