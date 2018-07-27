package errors

func NewConfigError(correlationId, code, message string) *ApplicationError {
    return &ApplicationError{ 
        Category: Misconfiguration,
        CorrelationId: correlationId,
        Code: code,
        Message: message,
        Status: 500,
    }
}
