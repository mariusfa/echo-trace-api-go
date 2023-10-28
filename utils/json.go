package utils

import (
	"encoding/json"
	"net/http"
)

func JsonBody[T any](r *http.Request, dto *T) error {
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}