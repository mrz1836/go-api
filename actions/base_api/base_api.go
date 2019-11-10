// Package baseapi is all the base requests and router configuration
package baseapi

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrz1836/go-api-router"
	"github.com/mrz1836/go-api/config"
	"github.com/mrz1836/go-api/jobs"
	"github.com/mrz1836/go-logger"
)

// RegisterRoutes register all the package specific routes
func RegisterRoutes(router *apirouter.Router) {

	// Load the service dependencies
	err := loadService()
	if err != nil {
		logger.Fatalf("failed to load required service dependencies - error: %s", err.Error())
	}

	// Set the main index page (navigating to slash)
	router.HTTPRouter.GET("/", router.Request(index))
	router.HTTPRouter.OPTIONS("/", router.SetCrossOriginHeaders)

	// Set the health request (used for load balancers)
	router.HTTPRouter.GET("/"+config.HealthRequestPath, router.Request(health))
	router.HTTPRouter.OPTIONS("/"+config.HealthRequestPath, router.SetCrossOriginHeaders)
	router.HTTPRouter.HEAD("/"+config.HealthRequestPath, router.SetCrossOriginHeaders)

	// Set the 404 handler (any request not detected)
	router.HTTPRouter.NotFound = http.HandlerFunc(notFound)

	// Set the method not allowed
	router.HTTPRouter.MethodNotAllowed = http.HandlerFunc(notAllowed)
}

// loadService will load all dependencies for the service
func loadService() (err error) {

	// Load jobs or services
	jobs.RunExampleJob(true, 5)

	// Done!
	logger.Data(2, logger.DEBUG, config.ApplicationModeAPI+" dependencies loaded!")
	return
}

// index basic request to /
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var returnResponse = map[string]interface{}{"message": "Welcome to the TonicPow API!"}
	apirouter.ReturnResponse(w, req, http.StatusOK, returnResponse)
}

// health basic request to return a health response
func health(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

// notFound handles all 404 requests
func notFound(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("404 occurred: %s", req.RequestURI), "Whoops - this request is not recognized", http.StatusNotFound, "")
	apirouter.ReturnResponse(w, req, apiError.Code, apiError)
	return
}

// notAllowed handles all 405 requests
func notAllowed(w http.ResponseWriter, req *http.Request) {
	apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("405 occurred: %s method: %s", req.RequestURI, req.Method), "Whoops - this method is not allowed", http.StatusMethodNotAllowed, "")
	apirouter.ReturnResponse(w, req, apiError.Code, apiError)
	return
}
