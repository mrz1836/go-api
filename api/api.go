// Package api is all the restful handler/router endpoint and methods
package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mrz1836/go-api-router"
	"github.com/mrz1836/go-api/actions/base"
)

// Handlers isolated the handlers / router for API (helps with testing)
func Handlers() *httprouter.Router {

	// Create a new router
	router := apirouter.New()

	// Create a middleware stack:
	//s := apirouter.NewStack()

	// Use your middleware:
	//s.Use(passThrough)

	base.RegisterRoutes(router)
	return router.HTTPRouter
}
