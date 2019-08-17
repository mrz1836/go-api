// Package base is all the base requests and router configuration
package base

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrz1836/go-api-router"
)

// RegisterRoutes register all the package specific routes
func RegisterRoutes(router *apirouter.Router) {
	// Set the main index page (navigating to slash)
	router.HTTPRouter.GET("/", router.Request(index))
	router.HTTPRouter.OPTIONS("/", router.SetCrossOriginHeaders)

	// Set the health request (used for load balancers)
	router.HTTPRouter.GET("/health", router.Request(health))
	router.HTTPRouter.OPTIONS("/health", router.SetCrossOriginHeaders)
	router.HTTPRouter.HEAD("/health", router.SetCrossOriginHeaders)

	// Set the 404 handler (any request not detected)
	router.HTTPRouter.NotFound = http.HandlerFunc(notFound) // todo: logging?
}

// index basic request to /
func index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, _ = fmt.Fprint(w, "Welcome to the Go API!\n")
}

// health basic request to return a health response
func health(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

// notFound handles all 404 requests
func notFound(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	apirouter.ReturnResponse(w, http.StatusNotFound, "404 - Request not recognized :-(", false)
}
