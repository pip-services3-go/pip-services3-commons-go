package errors

func NewConflictError(correlationId, code, message string) *ApplicationError {
    return &ApplicationError{ 
        Category: Conflict,
        CorrelationId: correlationId,
        Code: code,
        Message: message,
        Status: 409,
    }
}