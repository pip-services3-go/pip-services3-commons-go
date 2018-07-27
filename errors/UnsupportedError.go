package errors

func NewUnsupportedError(correlationId, code, message string) *ApplicationError {
	return &ApplicationError{
		Category:      Unsupported,
		CorrelationId: correlationId,
		Code:          code,
		Message:       message,
		Status:        500,
	}
}
