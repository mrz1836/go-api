package apirouter

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/mrz1836/go-parameters"
)

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
	//The ip address
	var ip string

	//Do we have a load balancer
	if xForward := req.Header.Get("X-Forwarded-For"); xForward != "" {
		//Set the ip as the given forwarded ip
		ip = xForward

		//Do we have more than one?
		if strings.Contains(ip, ",") {

			//Set the first ip address (from AWS)
			ip = strings.Split(ip, ",")[0]
		}
	} else {
		//Use the client address
		ip = strings.Split(req.RemoteAddr, ":")[0]

		//Remove bracket if local host
		ip = strings.Replace(ip, "[", "", 1)

		//Hack if no ip is found
		if len(ip) == 0 {
			ip = "localhost"
		}
	}

	//Return the ip address
	return ip
}
