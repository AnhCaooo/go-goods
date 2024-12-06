// AnhCao 2024
package auth

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestVerifyToken(t *testing.T) {

}

func TestExtractUserIdFromTokenClaim(t *testing.T) {
	tests := []struct {
		name        string
		token       *jwt.Token
		expectedID  string
		expectedErr string
	}{
		{
			name: "Valid token with userId",
			token: &jwt.Token{
				Claims: jwt.MapClaims{
					"userId": "12345",
				},
			},
			expectedID:  "12345",
			expectedErr: "",
		},
		{
			name: "Missing userId in claims",
			token: &jwt.Token{
				Claims: jwt.MapClaims{},
			},
			expectedID:  "",
			expectedErr: "user ID not found in token",
		},
		{
			name: "userId is not a string",
			token: &jwt.Token{
				Claims: jwt.MapClaims{
					"userId": 12345, // Integer instead of string
				},
			},
			expectedID:  "",
			expectedErr: "user ID not found in token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			userID, err := ExtractUserIdFromTokenClaim(tt.token)

			// Assert
			if tt.expectedErr != "" {
				if err == nil || err.Error() != tt.expectedErr {
					t.Errorf("expected error: %q, got: %v", tt.expectedErr, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if userID != tt.expectedID {
				t.Errorf("expected userID: %q, got: %q", tt.expectedID, userID)
			}
		})
	}
}
