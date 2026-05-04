package security

import "net/http"

type corsContext interface {
	GetHeader(string) string
	Header(string, string)
}

type CORSOptions struct {
	AllowedOrigins   []string
	AllowCredentials bool
}

var UnsafeWildcardOptions = CORSOptions{
	AllowedOrigins:   []string{"*"},
	AllowCredentials: true,
}

func wildcardCredentials(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func reflectedOriginCredentials(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func directReflectionCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func frameworkReflectionCredentials(c corsContext) {
	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	c.Header("Access-Control-Allow-Credentials", "true")
}
