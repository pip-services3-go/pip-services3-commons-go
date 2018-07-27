package errors

func NewConnectionError(correlationId, code, message string) *ApplicationError {
    return &ApplicationError{ 
        Category: NoResponse,
        CorrelationId: correlationId,
        Code: code,
        Message: message,
        Status: 500,
    }
}
