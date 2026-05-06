package security

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func fetchByChiRouteHost(r *http.Request) (*http.Response, error) {
	host := chi.URLParam(r, "host")
	target := fmt.Sprintf("http://%s/internal/status", host)
	return http.Get(target)
}

func proxyRawURL(r *http.Request) (*http.Response, error) {
	target := r.URL.Query().Get("url")
	return http.Get(target)
}

func postToCallback(r *http.Request) (*http.Response, error) {
	callback := r.FormValue("callback")
	return http.Post(callback, "application/json", nil)
}

func fetchBuiltRequest(r *http.Request, client *http.Client) (*http.Response, error) {
	target := r.Header.Get("X-Webhook-Url")
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func fetchByHostPath(r *http.Request) (*http.Response, error) {
	host := r.PathValue("host")
	target := fmt.Sprintf("http://%s/internal/status", host)
	return http.Get(target)
}

func fetchInboundHost(r *http.Request) (*http.Response, error) {
	target := "https://" + r.Host + "/internal/status"
	return http.Get(target)
}
