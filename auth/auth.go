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
//   - token: the JWT token that can be verified and used for authorization purposes
//   - error: An ERROR if the token cannot be parsed or is invalid; nil otherwise.
func VerifyToken(tokenString, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %s", err.Error())
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

// ExtractUserIdFromTokenClaim extracts the user ID from the claims in a JWT token.
// It assumes the token uses the `jwt.MapClaims` format and that the user ID is stored
// under the key `"userId"`.
//
// Parameters:
//   - token (*jwt.Token): The JWT token containing the claims.
//
// Returns:
//   - string: The extracted user ID if found.
//   - error: An error if the claims are invalid or if the user ID is not present.
//
// Usage Example:
//
//	token, err := auth.VerifyToken(tokenString, secretKey)
//	if err != nil {
//	    log.Fatalf("Failed to parse token: %v", err)
//	}
//
//	userID, err := ExtractUserIdFromTokenClaim(token)
//	if err != nil {
//	    log.Printf("Failed to extract user ID: %v", err)
//	} else {
//	    log.Printf("Extracted user ID: %s", userID)
//	}
//
// Errors:
//   - Returns an error if the claims cannot be cast to `jwt.MapClaims`.
//   - Returns an error if the `"userId"` key is missing or not a string.
func ExtractUserIdFromTokenClaim(token *jwt.Token) (string, error) {
	// Extract userID from token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims: failed to extract user ID from token claims")
	}

	userID, ok := claims["userId"].(string)
	if !ok {
		return "", fmt.Errorf("user ID not found in token")
	}
	return userID, nil
}
