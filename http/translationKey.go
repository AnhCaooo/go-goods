package goodsHTTP

// TranslationKey represents enumerations for translation keys used across go-goods client
type TranslationKey string

const (
	InvalidRequest     TranslationKey = "error_invalid_request"
	Unauthorized       TranslationKey = "error_unauthorized"
	UnauthorizedHeader TranslationKey = "error_no_authorization_header"
	InternalServer     TranslationKey = "error_internal_server"
	VerifyToken        TranslationKey = "error_verify_token"
	ExtractToken       TranslationKey = "error_extract_token"
	NotFound           TranslationKey = "error_not_found"
)
