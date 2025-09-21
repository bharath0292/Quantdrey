package qerrors

type ErrorCode string

const (
	// Domain/Business Errors
	CodeValidation     ErrorCode = "VALIDATION_ERROR"     // Data validation failure or input errors
	CodeAuthentication ErrorCode = "AUTHENTICATION_ERROR" // Login, token issues, unauthorized
	CodeConflict       ErrorCode = "CONFLICT"             // Conflict errors like duplicates (e.g., unique constraint)
	CodeNotFound       ErrorCode = "NOT_FOUND"            // Resource not found errors

	// Internal / System
	CodeInternal ErrorCode = "INTERNAL_ERROR" // Unexpected internal server errors
)
