/*
Package apirouter is a lightweight API router middleware for CORS, logging, and standardized error handling.

This package is intended to be used with Julien Schmidt's httprouter and uses MrZ's go-logger package.
*/
package apirouter

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mrz1836/go-logger"
	"github.com/mrz1836/go-parameters"
	"github.com/satori/go.uuid"
)

// Log formats for the request
const (
	defaultHeaders  string = "Accept, Content-Type, Content-Length, Cache-Control, Pragma, Accept-Encoding, X-CSRF-Token, Authorization, X-Auth-Cookie"
	defaultMethods  string = "POST, GET, OPTIONS, PUT, DELETE, LINK, HEAD"
	logParamsFormat string = "request_id=\"%s\" method=%s path=\"%s\" ip_address=\"%s\" user_agent=\"%s\" params=%v\n"
	logTimeFormat   string = "request_id=\"%s\" method=%s path=\"%s\" ip_address=\"%s\" user_agent=\"%s\" service=%dms status=%d\n"
)

// Package variables
var (
	ipAddressKey paramRequestKey = "ip_address"
	requestIDKey paramRequestKey = "request_id"
)

// paramRequestKey for context key
type paramRequestKey string

// Router is the configuration for the middleware service
type Router struct {
	CrossOriginAllowCredentials bool               `json:"cross_origin_allow_credentials" url:"cross_origin_allow_credentials"` // Allow credentials for BasicAuth()
	CrossOriginAllowHeaders     string             `json:"cross_origin_allow_headers" url:"cross_origin_allow_headers"`         // Allowed headers
	CrossOriginAllowMethods     string             `json:"cross_origin_allow_methods" url:"cross_origin_allow_methods"`         // Allowed methods
	CrossOriginAllowOrigin      string             `json:"cross_origin_allow_origin" url:"cross_origin_allow_origin"`           // Custom value for allow origin
	CrossOriginAllowOriginAll   bool               `json:"cross_origin_allow_origin_all" url:"cross_origin_allow_origin_all"`   // Allow all origins
	CrossOriginEnabled          bool               `json:"cross_origin_enabled" url:"cross_origin_enabled"`                     // Enable or Disable CrossOrigin
	HTTPRouter                  *httprouter.Router `json:"-" url:"-"`                                                           // J Schmidt's httprouter
	SkipLoggingPaths            []string           `json:"skip_logging_paths" url:"skip_logging_paths"`                         // Skip logging on these paths  (IE: /health)
}

// New returns a router middleware configuration to use for all future requests
func New() *Router {

	// Create new configuration
	config := new(Router)

	// Default is cross_origin = enabled
	config.CrossOriginEnabled = true

	// Default is to allow credentials for BasicAuth()
	config.CrossOriginAllowCredentials = true

	// Default is to allow all (easier to get started)
	config.CrossOriginAllowOriginAll = true

	// Default is for CrossOrigin to be enabled and these are common headers
	config.CrossOriginAllowHeaders = defaultHeaders

	// Default is for the common request methods
	config.CrossOriginAllowMethods = defaultMethods

	// Create the router
	config.HTTPRouter = new(httprouter.Router)

	// Turn on trailing slash redirect
	config.HTTPRouter.RedirectTrailingSlash = true
	config.HTTPRouter.RedirectFixedPath = true

	// return the default configuration
	return config
}

// Request will write the request to the logs before and after calling the handler
func (r *Router) Request(h httprouter.Handle) httprouter.Handle {
	return parameters.MakeHTTPRouterParsedReq(func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

		// Get the params from MakeHTTPRouterParsedReq()
		params := GetParams(req)

		// Start the custom response writer
		var writer *APIResponseWriter
		writer = &APIResponseWriter{
			IPAddress:      GetClientIPAddress(req),
			Method:         fmt.Sprintf("%s", req.Method),
			RequestID:      uuid.NewV4().String(),
			ResponseWriter: w,
			Status:         0, // future use with E-tags
			URL:            fmt.Sprintf("%s", req.URL),
			UserAgent:      req.UserAgent(),
		}

		// Store key information into the request that can be used by other methods
		req = req.WithContext(context.WithValue(req.Context(), ipAddressKey, writer.IPAddress))
		req = req.WithContext(context.WithValue(req.Context(), requestIDKey, writer.RequestID))

		// Set cross origin on each request that goes through logging
		r.SetCrossOriginHeaders(writer, req, ps)

		// Do we have paths to skip?
		// todo: this was added because some requests are confidential or health checks and they
		//  can't be split apart from the router
		var skipLogging bool
		if len(r.SkipLoggingPaths) > 0 {
			for _, path := range r.SkipLoggingPaths {
				if path == req.URL.Path {
					skipLogging = true
					break
				}
			}
		}

		// Skip logging this specific request
		if !skipLogging {

			// Start the log (timer)
			logger.Printf(logParamsFormat, writer.RequestID, writer.Method, writer.URL, writer.IPAddress, writer.UserAgent, params)
			start := time.Now()

			// Fire the request
			h(writer, req, ps)

			// Complete the timer and final log
			elapsed := time.Since(start)
			logger.Printf(logTimeFormat, writer.RequestID, writer.Method, writer.URL, writer.IPAddress, writer.UserAgent, int64(elapsed/time.Millisecond), writer.Status)

		} else {
			// Fire the request (no logging)
			h(writer, req, ps)
		}
	})
}

// RequestNoLogging will just call the handler without any logging
// Used for API calls that do not require any logging overhead
func (r *Router) RequestNoLogging(h httprouter.Handle) httprouter.Handle {
	return parameters.MakeHTTPRouterParsedReq(func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

		// Get the params from MakeHTTPRouterParsedReq()
		//params := GetParams(req)

		// Start the custom response writer
		var writer *APIResponseWriter
		writer = &APIResponseWriter{
			IPAddress:      GetClientIPAddress(req),
			Method:         fmt.Sprintf("%s", req.Method),
			RequestID:      uuid.NewV4().String(),
			ResponseWriter: w,
			Status:         0, // future use with E-tags
			URL:            fmt.Sprintf("%s", req.URL),
			UserAgent:      req.UserAgent(),
		}

		// Set cross origin on each request that goes through logging
		r.SetCrossOriginHeaders(writer, req, ps)

		// Fire the request
		h(writer, req, ps)
	})
}

// BasicAuth wraps a request for Basic Authentication (RFC 2617)
func (r *Router) BasicAuth(h httprouter.Handle, requiredUser, requiredPassword string, errorResponse interface{}) httprouter.Handle {

	// Return the function up the chain
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := req.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, req, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			ReturnResponse(w, req, http.StatusUnauthorized, errorResponse)
		}
	}
}

// SetCrossOriginHeaders sets the cross-origin headers if enabled
func (r *Router) SetCrossOriginHeaders(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Turned cross_origin off? Just return
	if !r.CrossOriginEnabled {
		return
	}

	// On for all origins?
	if r.CrossOriginAllowOriginAll {
		w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
		w.Header().Set("Vary", "Origin")
	} else { //Only the origin set by config
		w.Header().Set("Access-Control-Allow-Origin", r.CrossOriginAllowOrigin)
	}

	// Allow credentials (used for BasicAuth)
	if r.CrossOriginAllowCredentials {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	// Allowed methods to accept
	w.Header().Set("Access-Control-Allow-Methods", r.CrossOriginAllowMethods)

	// Allowed headers to accept
	w.Header().Set("Access-Control-Allow-Headers", r.CrossOriginAllowHeaders)
}
