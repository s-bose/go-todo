package auth

import (
	"fmt"
	"go-todo/app/config"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type ClaimsPayload struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GetToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 && parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization rheader")
	}

	authToken := strings.TrimSpace(parts[1])
	return authToken, nil
}

func ValidateJWT(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &ClaimsPayload{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Config.JWTSecret), nil
	})
}

func CreateJWT(userID *uuid.UUID) (string, error) {
	duration := time.Second * time.Duration(config.Config.JWTExpirationSeconds)

	payload := &ClaimsPayload{
		UserID: userID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(config.Config.JWTSecret))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
