package auth

import (
	"go-todo/app/utils"
	"net/http"
)

func JWTMiddleware(handlerFunc *http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := GetToken(r)

		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, err.Error())
			return
		}

		ValidateJWT(token)

	}
}
