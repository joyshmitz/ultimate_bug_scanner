package security

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var allowedProxyHosts = map[string]bool{
	"api.example.com": true,
	"cdn.example.com": true,
}

func safeProxyURL(raw string) (*url.URL, error) {
	upstream, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	if upstream.Scheme != "https" || !allowedProxyHosts[upstream.Hostname()] {
		return nil, errors.New("blocked proxy target")
	}
	return upstream, nil
}

func proxyWithSafeTarget(r *http.Request) (*httputil.ReverseProxy, error) {
	upstream, err := safeProxyURL(r.URL.Query().Get("target"))
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(upstream), nil
}

func rewriteWithSafeTarget(r *http.Request) (*httputil.ReverseProxy, error) {
	upstream, err := safeProxyURL(r.Header.Get("X-Upstream"))
	if err != nil {
		return nil, err
	}
	return &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(upstream)
		},
	}, nil
}

func directorWithSafeTarget(r *http.Request) (*httputil.ReverseProxy, error) {
	upstream, err := safeProxyURL("https://" + r.Host)
	if err != nil {
		return nil, err
	}
	return &httputil.ReverseProxy{
		Director: func(out *http.Request) {
			out.URL.Scheme = upstream.Scheme
			out.URL.Host = upstream.Host
		},
	}, nil
}
