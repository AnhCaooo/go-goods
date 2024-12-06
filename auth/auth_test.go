// AnhCao 2024
package auth

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestVerifyToken(t *testing.T) {

}

func TestExtractValueFromTokenClaim(t *testing.T) {
	tests := []struct {
		name        string
		token       *jwt.Token
		valueField  string
		expectedVal string
		expectedErr string
	}{
		{
			name: "Valid token with specified field",
			token: &jwt.Token{
				Claims: jwt.MapClaims{
					"userId": "12345",
				},
			},
			valueField:  "userId",
			expectedVal: "12345",
			expectedErr: "",
		},
		{
			name: "Valid token with specified field in bigger Claims",
			token: &jwt.Token{
				Claims: jwt.MapClaims{
					"id":    "12345",
					"email": "test@mail.com",
					"phone": "",
				},
			},
			valueField:  "id",
			expectedVal: "12345",
			expectedErr: "",
		},
		{
			name: "Missing field in claims",
			token: &jwt.Token{
				Claims: jwt.MapClaims{},
			},
			valueField:  "userId",
			expectedVal: "",
			expectedErr: "userId not found in token",
		},
		{
			name: "Field value is not a string",
			token: &jwt.Token{
				Claims: jwt.MapClaims{
					"userId": 12345, // Integer instead of string
				},
			},
			valueField:  "userId",
			expectedVal: "",
			expectedErr: "userId not found in token",
		},
		{
			name: "Valid token with different field",
			token: &jwt.Token{
				Claims: jwt.MapClaims{
					"role": "admin",
				},
			},
			valueField:  "role",
			expectedVal: "admin",
			expectedErr: "",
		},
		{
			name: "Missing different field in claims",
			token: &jwt.Token{
				Claims: jwt.MapClaims{
					"role": "admin",
				},
			},
			valueField:  "userId",
			expectedVal: "",
			expectedErr: "userId not found in token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			value, err := ExtractValueFromTokenClaim(tt.token, tt.valueField)

			// Assert
			if tt.expectedErr != "" {
				if err == nil || err.Error() != tt.expectedErr {
					t.Errorf("expected error: %q, got: %v", tt.expectedErr, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if value != tt.expectedVal {
				t.Errorf("expected value: %q, got: %q", tt.expectedVal, value)
			}
		})
	}
}
