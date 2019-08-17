package apirouter

import (
	"net/http"
	"net/url"
	"strings"
)

// GetParams gets the params from the http request (parsed once on log request)
func GetParams(req *http.Request) url.Values {
	params := req.Context().Value(paramKey).(url.Values)
	return params
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
