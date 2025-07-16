package goodsContext

type key string

const (
	ContextKey key = "CONTEXT_KEY" // Context keys for reading Context values
)

// UserContext holds contextual data extracted from the request,
// such as the authenticated user's ID and session ID.
//
// It is intended to be stored in the request's context using context.WithValue,
// and retrieved later in handlers or middleware that require access to user-specific information.
//
// Usage:
//
//	// Setting the context
//	userCtx := UserContext{
//		UserID:    "12345",
//		SessionID: "abcde-session",
//	}
//	ctx := context.WithValue(r.Context(), UserContextKey, userCtx)
//	r = r.WithContext(ctx)
//
//	// Retrieving the context
//	userCtx, ok := r.Context().Value(UserContextKey).(UserContext)
//	if !ok {
//		// handle missing or malformed context
//	}
//
//	fmt.Println(userCtx.UserID, userCtx.SessionID)
type UserContext struct {
	UserID    string
	SessionID string
}
