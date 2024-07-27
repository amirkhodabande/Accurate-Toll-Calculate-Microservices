package handlers

import (
	"encoding/json"
	"net/http"
)

type ApiFunc func(w http.ResponseWriter, r *http.Request) error

func MakeApiFunc(fn ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			// TODO: create a custom err struct and get the status from it
			writeJson(w, http.StatusInternalServerError, map[string]any{
				"success": false, "message": err.Error(),
			})
		}
	}
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
