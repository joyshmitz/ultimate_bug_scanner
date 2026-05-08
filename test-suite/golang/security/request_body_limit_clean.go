package security

import (
	"encoding/json"
	"io"
	"net/http"
)

const maxBodyBytes int64 = 1 << 20

type uploadPayload struct {
	Name string `json:"name"`
}

func importPayload(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	return io.ReadAll(r.Body)
}

func readLimitedUpload(r *http.Request) ([]byte, error) {
	limited := io.LimitReader(r.Body, maxBodyBytes)
	return io.ReadAll(limited)
}

func decodePayload(w http.ResponseWriter, r *http.Request) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	var payload uploadPayload
	return json.NewDecoder(r.Body).Decode(&payload)
}

func decodeLimitedPayload(r *http.Request) error {
	limited := io.LimitReader(r.Body, maxBodyBytes)
	var payload uploadPayload
	return json.NewDecoder(limited).Decode(&payload)
}
