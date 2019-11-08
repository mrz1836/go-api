package apirouter

import (
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/mrz1836/go-parameters"
)

// Unix epoch time
var epoch = time.Unix(0, 0).Format(time.RFC1123)

// Taken from https://github.com/mytrile/nocache
var noCacheHeaders = map[string]string{
	"Cache-Control":   "no-cache, no-store, no-transform, must-revalidate, private, max-age=0",
	"Expires":         epoch,
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}
var etagHeaders = []string{
	"ETag",
	"If-Match",
	"If-Modified-Since",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

// camelCaseRe camel case regex
var camelCaseRe = regexp.MustCompile(`(?:^[\p{Ll}]|API|JSON|IP|URL|_?\d+|_|[\p{Lu}]+)[\p{Ll}]*`)

// SnakeCase takes a camelCaseWord and breaks it into camel_case_word
func SnakeCase(str string) string {
	words := camelCaseRe.FindAllString(str, -1)

	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(strings.Replace(words[i], "_", "", -1))
	}

	return strings.Join(words, "_")
}

// FindString returns the index of the first instance of needle in the array or -1 if it could not be found
func FindString(needle string, haystack []string) int {
	for i, straw := range haystack {
		if needle == straw {
			return i
		}
	}
	return -1
}

// GetParams gets the params from the http request (parsed once on log request)
// Helper method for the parameters package
func GetParams(req *http.Request) *parameters.Params {
	return parameters.GetParams(req)
}

// PermitParams will remove all keys that not allowed
// Helper method for the parameters package
func PermitParams(params *parameters.Params, allowedKeys []string) {
	params.Permit(allowedKeys)
}

// GetIPFromRequest gets the stored ip from the request if found
func GetIPFromRequest(req *http.Request) (ip string, ok bool) {
	ip, ok = req.Context().Value(ipAddressKey).(string)
	return
}

// GetRequestID gets the stored request id from the request
func GetRequestID(req *http.Request) (id string, ok bool) {
	id, ok = req.Context().Value(requestIDKey).(string)
	return
}

// GetClientIPAddress gets the client ip address
func GetClientIPAddress(req *http.Request) string {
	// The ip address
	var ip string

	// Do we have a load balancer
	if xForward := req.Header.Get("X-Forwarded-For"); xForward != "" {
		// Set the ip as the given forwarded ip
		ip = xForward

		// Do we have more than one?
		if strings.Contains(ip, ",") {

			// Set the first ip address (from AWS)
			ip = strings.Split(ip, ",")[0]
		}
	} else {
		// Use the client address
		ip = strings.Split(req.RemoteAddr, ":")[0]

		// Remove bracket if local host
		ip = strings.Replace(ip, "[", "", 1)

		// Hack if no ip is found
		//if len(ip) == 0 {
		//	ip = "localhost"
		//}
	}

	// Parsing will also validate if it's IPv4 or IPv6
	parsed := net.ParseIP(ip)
	if parsed == nil {
		ip = ""
	} else {
		ip = parsed.String()
	}

	//Return the ip address
	return ip
}

// NoCache is a simple piece of middleware that sets a number of HTTP headers to prevent
// a router (or subrouter) from being cached by an upstream proxy and/or client.
//
// As per http://wiki.nginx.org/HttpProxyModule - NoCache sets:
//      Expires: Thu, 01 Jan 1970 00:00:00 UTC
//      Cache-Control: no-cache, private, max-age=0
//      X-Accel-Expires: 0
//      Pragma: no-cache (for HTTP/1.0 proxies/clients)
func NoCache(w http.ResponseWriter, req *http.Request) {

	// Delete any ETag headers that may have been set
	for _, v := range etagHeaders {
		if req.Header.Get(v) != "" {
			req.Header.Del(v)
		}
	}

	// Set our NoCache headers
	for k, v := range noCacheHeaders {
		w.Header().Set(k, v)
	}
}
