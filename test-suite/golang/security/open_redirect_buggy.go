package security

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type redirectContext interface {
	QueryParam(string) string
	Redirect(int, string) error
}

func redirectFromRouteParam(w http.ResponseWriter, r *http.Request) {
	target := chi.URLParam(r, "next")
	http.Redirect(w, r, target, http.StatusFound)
}

func redirectFromQuery(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")
	http.Redirect(w, r, next, http.StatusFound)
}

func redirectFromHeader(w http.ResponseWriter, r *http.Request) {
	target := r.Header.Get("X-Return-To")
	w.Header().Set("Location", target)
	w.WriteHeader(http.StatusFound)
}

func redirectFromFramework(c redirectContext) error {
	return c.Redirect(http.StatusFound, c.QueryParam("redirect"))
}
