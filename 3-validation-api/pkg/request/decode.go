package request

import (
	"encoding/json"
	"io"
)

func Decode[T interface{}](body io.ReadCloser) (*T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
