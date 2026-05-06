package security

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func proxyFromQuery(r *http.Request) *httputil.ReverseProxy {
	target := r.URL.Query().Get("target")
	upstream, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(upstream)
}

func rewriteFromHeader(r *http.Request) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			raw := r.Header.Get("X-Upstream")
			upstream, _ := url.Parse(raw)
			pr.SetURL(upstream)
		},
	}
}

func directorFromHost(r *http.Request) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(out *http.Request) {
			out.URL.Scheme = "https"
			out.URL.Host = r.Host
		},
	}
}
