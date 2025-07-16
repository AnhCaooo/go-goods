package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/AnhCaooo/go-goods/auth"
	goodsContext "github.com/AnhCaooo/go-goods/context"
	goodsHTTP "github.com/AnhCaooo/go-goods/http"
)

func Authenticate(next http.Handler, byPassPaths []string, jwtSecret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if shouldBypassAuthentication(r.URL.Path, byPassPaths) {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			goodsHTTP.Error(w, http.StatusForbidden, "No Authorization header provided", goodsHTTP.UnauthorizedHeader)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := auth.VerifyToken(tokenString, jwtSecret)
		if err != nil {
			goodsHTTP.Error(w, http.StatusUnauthorized, "Failed to verify token", goodsHTTP.VerifyToken)
			return
		}

		// due to 'Supabase' authentication, it stores userId via "sub" field
		userID, err := auth.ExtractValueFromTokenClaim(token, "sub")
		if err != nil {
			goodsHTTP.Error(w, http.StatusUnauthorized, "Failed to extract token", goodsHTTP.ExtractToken)
			return
		}

		sessionID, err := auth.ExtractValueFromTokenClaim(token, "session_id")
		if err != nil {
			goodsHTTP.Error(w, http.StatusUnauthorized, "Failed to extract token", goodsHTTP.ExtractToken)
			return
		}

		userCtx := goodsContext.UserContext{
			UserID:    userID,
			SessionID: sessionID,
		}

		// Add userCtx to the context
		ctx := context.WithValue(r.Context(), goodsContext.ContextKey, userCtx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// shouldBypassAuthentication checks if the request path should bypass authentication (do not need authentication)
func shouldBypassAuthentication(path string, byPassPaths []string) bool {
	for _, p := range byPassPaths {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}
