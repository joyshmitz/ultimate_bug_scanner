package security

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
)

func safeRedirectTarget(raw string) string {
	parsed, err := url.Parse(raw)
	if err != nil {
		return "/"
	}
	if parsed.IsAbs() {
		if parsed.Scheme == "https" && parsed.Hostname() == "app.example.com" {
			return parsed.String()
		}
		return "/"
	}
	if !strings.HasPrefix(parsed.Path, "/") || strings.HasPrefix(parsed.Path, "//") {
		return "/"
	}
	return parsed.String()
}

func redirectWithSafeHelper(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")
	http.Redirect(w, r, safeRedirectTarget(next), http.StatusFound)
}

func locationWithSafeHelper(w http.ResponseWriter, r *http.Request) {
	target := safeRedirectTarget(r.Header.Get("X-Return-To"))
	w.Header().Set("Location", target)
	w.WriteHeader(http.StatusFound)
}

func redirectRouteParamWithSafeHelper(w http.ResponseWriter, r *http.Request) {
	target := safeRedirectTarget(chi.URLParam(r, "next"))
	http.Redirect(w, r, target, http.StatusFound)
}
