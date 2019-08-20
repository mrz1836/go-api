// Package api is all the restful handler/router endpoint and methods
package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mrz1836/go-api-router"
	"github.com/mrz1836/go-api/actions/base"
	"github.com/mrz1836/go-api/actions/person"
)

// Handlers isolated the handlers / router for API (helps with testing)
func Handlers() *httprouter.Router {

	// Create a new router
	router := apirouter.New()

	// Create a middleware stack:
	//s := apirouter.NewStack()

	// Use your middleware:
	//s.Use(passThrough)

	// The base api routes
	base.RegisterRoutes(router)

	// The person actions
	person.RegisterRoutes(router)

	// Return the router
	return router.HTTPRouter
}
