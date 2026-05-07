package security

import (
	"crypto/tls"
	"net/http"
)

var insecureClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func insecureTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: insecureTLSConfig(),
	}
}

func insecureTLSConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}
