package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserControllers(t *testing.T) {
	t.Run("POST /register error, email invalid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/users/register", nil)
		if err != nil {
			log.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()
		router.Handle("/api/v1/users/", http.StripPrefix("/api/v1/users", RegisterUserRoutes()))

		sqlmock
	})
}
