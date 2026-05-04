package security

import "net/http"

type cleanCORSOptions struct {
	AllowedOrigins   []string
	AllowCredentials bool
}

var TrustedOriginOptions = cleanCORSOptions{
	AllowedOrigins: []string{
		"https://app.example.com",
		"https://admin.example.com",
	},
	AllowCredentials: true,
}

func publicWildcardCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Vary", "Origin")
}

func trustedCredentialedCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "https://app.example.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Vary", "Origin")
}
