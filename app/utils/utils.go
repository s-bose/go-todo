package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ParseJSON(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(&v)
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type Dict map[string]any

func ValidateJSON(r *http.Request, schema interface{}) error {
	validator := validator.New()

	err := ParseJSON(r, schema)
	if err != nil {
		return err
	}
	return validator.Struct(schema)
}
