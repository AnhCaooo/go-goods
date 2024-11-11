// AnhCao 2024
package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// receive token and JWT secret key to do verify the token
func VerifyToken(tokenString, secretKey string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return fmt.Errorf("failed to parse token: %s", err.Error())
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
