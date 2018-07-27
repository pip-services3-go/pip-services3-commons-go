package errors

func NewInternalError(correlationId, code, message string) *ApplicationError {
    return &ApplicationError{ 
        Category: Internal,
        CorrelationId: correlationId,
        Code: code,
        Message: message,
        Status: 500,
    }
}