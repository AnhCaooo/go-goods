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

// ExtractValueFromTokenClaim extracts a specified value from the claims in a JWT token.
// It assumes the token uses the `jwt.MapClaims` format and that the value is stored
// under the provided `valueField` key.
//
// Parameters:
//   - token (*jwt.Token): The JWT token containing the claims.
//   - valueField (string): The key in the token claims whose value needs to be extracted.
//
// Returns:
//   - string: The extracted value if found.
//   - error: An error if the claims are invalid, or if the value is missing or not a string.
//
// Usage Example:
//
//   token, err := jwt.Parse(tokenString, keyFunc)
//   if err != nil {
//       log.Fatalf("Failed to parse token: %v", err)
//   }
//
//   userID, err := ExtractValueFromTokenClaim(token, "userId")
//   if err != nil {
//       log.Printf("Failed to extract user ID: %v", err)
//   } else {
//       log.Printf("Extracted user ID: %s", userID)
//   }
//
// Errors:
//   - Returns an error if the claims cannot be cast to `jwt.MapClaims`.
//   - Returns an error if the specified `valueField` is missing or not a string.

func ExtractValueFromTokenClaim(token *jwt.Token, valueField string) (string, error) {
	// Extract value from given field from token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	value, ok := claims[valueField].(string)
	if !ok {
		return "", fmt.Errorf("%s not found in token", valueField)
	}
	return value, nil
}
