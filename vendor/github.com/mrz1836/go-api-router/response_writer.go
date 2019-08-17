package apirouter

import (
	"bytes"
	"net/http"
	"time"
)

// APIResponseWriter wraps the ResponseWriter and stores the status of the request.
// It is used by the LogRequest middleware
type APIResponseWriter struct {
	http.ResponseWriter
	Buffer          bytes.Buffer  `json:"-" url:"-"`
	CacheIdentifier []string      `json:"cache_identifier" url:"cache_identifier"`
	CacheTTL        time.Duration `json:"cache_ttl" url:"cache_ttl"`
	IPAddress       string        `json:"ip_address" url:"ip_address"`
	Method          string        `json:"method" url:"method"`
	NoWrite         bool          `json:"no_write" url:"no_write"`
	RequestID       string        `json:"request_id" url:"request_id"`
	Status          int           `json:"status" url:"status"`
	URL             string        `json:"url" url:"url"`
	UserAgent       string        `json:"user_agent" url:"user_agent"`
}

// AddCacheIdentifier add cache identifier to the response writer
func (r *APIResponseWriter) AddCacheIdentifier(identifier string) {
	if r.CacheIdentifier == nil {
		r.CacheIdentifier = make([]string, 0, 2)
	}
	r.CacheIdentifier = append(r.CacheIdentifier, identifier)
}

// StatusCode give a way to get the status code
func (r *APIResponseWriter) StatusCode() int {
	return r.Status
}

// Header returns the http.Header that will be written to the response
func (r *APIResponseWriter) Header() http.Header {
	return r.ResponseWriter.Header()
}

// WriteHeader will write the header to the client, setting the status code
func (r *APIResponseWriter) WriteHeader(status int) {
	r.Status = status
	if !r.NoWrite {
		r.ResponseWriter.WriteHeader(status)
	}
}

// Write writes the data out to the client, if WriteHeader was not called, it will write status http.StatusOK (200)
func (r *APIResponseWriter) Write(data []byte) (int, error) {
	if r.Status == 0 {
		r.Status = http.StatusOK
	}

	if r.NoWrite {
		return r.Buffer.Write(data)
	}

	return r.ResponseWriter.Write(data)
}
