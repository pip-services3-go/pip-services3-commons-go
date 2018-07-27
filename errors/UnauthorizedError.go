package errors

func NewUnauthorizedError(correlationId, code, message string) *ApplicationError {
    return &ApplicationError{ 
        Category: Unauthorized,
        CorrelationId: correlationId,
        Code: code,
        Message: message,
        Status: 401,
    }
}