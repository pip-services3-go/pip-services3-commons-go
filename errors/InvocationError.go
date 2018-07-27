package errors

func NewInvocationError(correlationId, code, message string) *ApplicationError {
	return &ApplicationError{
		Category:      FailedInvocation,
		CorrelationId: correlationId,
		Code:          code,
		Message:       message,
		Status:        500,
	}
}
