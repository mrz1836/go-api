package apirouter

import (
	"encoding/json"
	"net/http"

	"github.com/mrz1836/go-logger"
)

const (
	// ErrCodeUnknown unknown error code (example)
	ErrCodeUnknown int = 600
)

// APIError is the enriched error message for API related errors
type APIError struct {
	Code            int         `json:"code" url:"code"`                 // Associated error code
	Data            interface{} `json:"data" url:"data"`                 // Arbitrary data that is relevant
	InternalMessage string      `json:"-" url:"-"`                       // Internal message for engineers
	IPAddress       string      `json:"ip_address" url:"ip_address"`     // Current IP of user
	Method          string      `json:"method" url:"method"`             // Method requested (IE: POST)
	PublicMessage   string      `json:"message" url:"message"`           // Public error message
	RequestGUID     string      `json:"request_guid" url:"request_guid"` // Unique Request ID for tracking
	URL             string      `json:"url" url:"url"`                   // Requesting URL
}

// ErrorFromResponse generates a new error struct using CustomResponseWriter from LogRequest()
func ErrorFromResponse(w *APIResponseWriter, internalMessage string, publicMessage string, errorCode int, data interface{}) *APIError {

	// Log the error
	logError(errorCode, internalMessage)

	return &APIError{
		Code:            errorCode,
		Data:            data,
		InternalMessage: internalMessage,
		IPAddress:       w.IPAddress,
		Method:          w.Method,
		PublicMessage:   publicMessage,
		RequestGUID:     w.RequestID,
		URL:             w.URL,
	}
}

// ErrorFromRequest gives an error without a response writer using the request
func ErrorFromRequest(req *http.Request, internalMessage string, publicMessage string, errorCode int, data interface{}) *APIError {

	// Get values from req if available
	ip, _ := GetIPFromRequest(req)
	id, _ := GetRequestID(req)

	// Log the error
	logError(errorCode, internalMessage)

	// Return an error
	return &APIError{
		Code:            errorCode,
		Data:            data,
		InternalMessage: internalMessage,
		IPAddress:       ip,
		Method:          req.Method,
		PublicMessage:   publicMessage,
		RequestGUID:     id,
		URL:             req.URL.String(),
	}
}

// logError will log the internal message and code for diagnosing
func logError(errorCode int, internalMessage string) {

	// Skip non-error codes
	if errorCode < 400 || errorCode == 404 {
		return
	}

	// Switch based on severity
	var logLevel logger.LogLevel
	if errorCode == 400 || errorCode > 422 {
		logLevel = logger.ERROR
	} else {
		logLevel = logger.WARN
	}

	logger.Data(2, logLevel, "internal error message: "+internalMessage, logger.MakeParameter("code", errorCode))
}

// Error returns the string error message (only public message)
func (e *APIError) Error() string {
	return e.PublicMessage
}

// ErrorCode returns the error code
func (e *APIError) ErrorCode() int {
	return e.Code
}

// JSON returns the entire public version of the error message
func (e *APIError) JSON() (string, error) {
	m, err := json.Marshal(e)
	return string(m), err
}

// Internal returns the string error message (only internal message)
func (e *APIError) Internal() string {
	return e.InternalMessage
}
