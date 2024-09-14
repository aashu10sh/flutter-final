package auth_utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JWT_SECRET string = "Barcelona"
)

func SignJWTToken(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jid": username,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(JWT_SECRET))

	return tokenString, err
}
