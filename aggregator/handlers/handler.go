package handlers

import (
	"aggregator/handlers/custom_errors"
	"encoding/json"
	"fmt"
	"net/http"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func MakeApiFunc(method string, fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			writeJson(w, http.StatusInternalServerError, map[string]any{
				"success": false, "message": fmt.Sprintf("method not allowed! %s", method),
			})
			return
		}

		if err := fn(w, r); err != nil {
			if customError, ok := err.(custom_errors.Error); ok {
				writeJson(w, customError.Status(), customError.MsgAndErr())
				return
			}

			writeJson(w, http.StatusInternalServerError, map[string]any{
				"success": false, "message": "fucked up :))",
			})
			return
		}
	}
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
