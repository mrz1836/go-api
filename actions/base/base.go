// Package base is all the base requests and router configuration
package base

import (
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
	//router.HTTPRouter.NotFound = http.HandlerFunc(notFound) // todo: logging?
	router.HTTPRouter.NotFound = http.HandlerFunc(notFound) // todo: logging?

	// Set the method not allowed
	router.HTTPRouter.MethodNotAllowed = http.HandlerFunc(notAllowed) // todo: logging?
}

// index basic request to /
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var returnResponse = map[string]interface{}{"message": "Welcome to the Go API!"}
	apirouter.ReturnResponse(w, req, http.StatusOK, returnResponse)
}

// health basic request to return a health response
func health(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

// notFound handles all 404 requests
func notFound(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	apiError := apirouter.ErrorFromRequest(req, "request is not recognized", "Whoops - this request is not recognized", http.StatusNotFound, "")
	apirouter.ReturnResponse(w, req, http.StatusNotFound, apiError)
	return
}

// notAllowed handles all 405 requests
func notAllowed(w http.ResponseWriter, req *http.Request) {
	apiError := apirouter.ErrorFromRequest(req, "request is not allowed", "Whoops - this method is not allowed", http.StatusMethodNotAllowed, "")
	apirouter.ReturnResponse(w, req, http.StatusMethodNotAllowed, apiError)
	return
}
