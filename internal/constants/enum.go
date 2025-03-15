package constants

// Log Level
const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
	LevelFatal = "fatal"
	LevelPanic = "panic"
)

// General errors
const (
	UNPROCESSABLE_ENTITY    = 1001
	INTERNAL_EXCEPTION      = 1002
	UNAUTHORIZED            = 1003
	INVALID_QUERY_PARAMETER = 1004
	FORBIDDEN_ACCESS        = 1005
	VALIDATION_VAILED       = 1006

	EMAIL_ALREADY_REGISTERED  = 2001
	INVALID_EMAIL_OR_PASSWORD = 2002

	STORE_NOT_FOUND          = 3001
	PARENT_STORE_NOT_FOUND   = 3002
	ALL_FIELD_MUST_BE_FILLED = 3003

	USER_NOT_FOUND      = 4001
	EMAIL_IS_REQUIRED   = 4002
	EMAIL_MUST_BE_VALID = 4003

	CATEGORY_NOT_FOUND = 5001

	PRODUCT_NOT_FOUND = 6001
)
