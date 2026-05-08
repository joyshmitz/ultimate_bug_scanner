package security

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type uploadPayload struct {
	Name string `json:"name"`
}

func importPayload(r *http.Request) ([]byte, error) {
	return io.ReadAll(r.Body)
}

func legacyUpload(req *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func decodePayload(r *http.Request) error {
	var payload uploadPayload
	return json.NewDecoder(r.Body).Decode(&payload)
}

func decodePayloadViaDecoder(req *http.Request) error {
	var payload uploadPayload
	decoder := json.NewDecoder(req.Body)
	return decoder.Decode(&payload)
}
