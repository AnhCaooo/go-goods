package goodsHTTP

// TranslationKey represents enumerations for translation keys used across go-goods client
type TranslationKey string

const (
	InvalidRequest TranslationKey = "error_invalid_request"
	Unauthorized   TranslationKey = "error_unauthorized"
	InternalServer TranslationKey = "error_internal_server"
)
