// AnhCao 2024
package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// VerifyToken verifies the authenticity of a JWT token using the provided secret key.
//
// This function parses the provided `tokenString`, verifies its signature using the provided
// `secretKey`, and checks the token's validity. If the token is invalid or the verification
// fails, an error is returned.
//
// EXAMPLE USAGE:
//
//	if err := auth.VerifyToken(tokenString, secretKey); err != nil {
//	    return err
//	}
//
// PARAMETERS:
//   - tokenString: The JWT token string to be verified.
//   - secretKey: The SECRET KEY used to verify the token's signature.
//
// RETURNS:
//   - error: An ERROR if the token cannot be parsed or is invalid; nil otherwise.
func VerifyToken(tokenString, secretKey string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return fmt.Errorf("failed to parse token: %s", err.Error())
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
