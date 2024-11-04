package auth

import (
	"context"
	"go-todo/app/db"
	"go-todo/app/services"
	"go-todo/app/utils"
	"net/http"

	"github.com/google/uuid"
)

// JWT Authentication middleware
func JWTMiddleware(handlerFunc http.HandlerFunc, DB *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := GetToken(r)

		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, utils.Dict{"message": "authorization header missing or corrupt", "error": err.Error()})
			return
		}

		token, err := ValidateJWT(tokenString)

		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Dict{"message": "unauthorized", "error": err.Error()})
			return
		}

		if !token.Valid {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Dict{"message": "token invalid or expired", "error": err.Error()})
			return
		}

		claims := token.Claims.(*UserClaims)
		userID := claims.UserID

		userService := services.UserService{Db: DB.Db}

		user, err := userService.GetUserById(uuid.MustParse(userID))
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Dict{"message": "unauthorized", "error": err.Error()})
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", user.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}
