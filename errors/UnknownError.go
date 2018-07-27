package errors

func NewUnknownError(correlationId, code, message string) *ApplicationError {
	return &ApplicationError{
		Category:      Unknown,
		CorrelationId: correlationId,
		Code:          code,
		Message:       message,
		Status:        500,
	}
}
