package security

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

func verifiedClient(pool *x509.CertPool) *http.Client {
	return &http.Client{
		Transport: verifiedTransport(pool),
	}
}

func verifiedTransport(pool *x509.CertPool) *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
			ServerName: "api.example.com",
			RootCAs:    pool,
		},
	}
}

func defaultTLSConfig() *tls.Config {
	return &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
}
