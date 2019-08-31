package parameters

import (
	"compress/gzip"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strings"
)

//CORSHeaders adds cross origin resource sharing headers to a response
func CORSHeaders(fn http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		fn(w, r)
	}
}

//SendCORS sends a cross origin resource sharing header only
func SendCORS(w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
}

//JSONResp will set the content-type to application/json
func JSONResp(fn httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		fn(rw, req, p)
	}
}

//GeneralResponse calls the default wrappers: EnableGZIP, MakeHTTPRouterParsedReq, CORSHeaders
func GeneralResponse(fn http.HandlerFunc) httprouter.Handle {
	return EnableGZIP(MakeHTTPRouterParsedReq(CORSHeaders(fn)))
}

//GeneralJSONResponse calls the default wrappers for a json response: EnableGZIP, JSONResp, MakeHTTPRouterParsedReq, CORSHeaders
func GeneralJSONResponse(fn http.HandlerFunc) httprouter.Handle {
	return EnableGZIP(JSONResp(MakeHTTPRouterParsedReq(CORSHeaders(fn))))
}

var filterReplace = [...]string{"FILTERED"}

//FilteredKeys is a lower case array of keys to filter when logging
var FilteredKeys []string

//filterMap will filter the parameters and not log parameters with sensitive data. To add more parameters - see the if in the loop
func filterMap(params *Params) *Params {
	var filtered Params
	filtered.Values = make(map[string]interface{}, len(params.Values))

	for k, v := range params.Values {
		if contains(FilteredKeys, k) {
			filtered.Values[k] = filterReplace[:]
		} else if b, ok := v.([]byte); ok {
			filtered.Values[k] = string(b)
		} else {
			filtered.Values[k] = v
		}
	}
	return &filtered
}

//gzipResponseWriter gzip response writer
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

//Write write the content
func (w gzipResponseWriter) Write(b []byte) (int, error) {
	if "" == w.Header().Get("Content-Type") {
		// If no content type, apply sniffing algorithm to un-gzipped body.
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.Writer.Write(b)
}

//EnableGZIP will attempt to compress the response if the client has passed a header value for Accept-Encoding which allows gzip
func EnableGZIP(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r, p)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		fn(gzr, r, p)
		_ = gz.Close()
	}
}
