package errors

func NewInvalidStateError(correlationId, code, message string) *ApplicationError {
    return &ApplicationError{ 
        Category: InvalidState,
        CorrelationId: correlationId,
        Code: code,
        Message: message,
        Status: 500,
    }
}