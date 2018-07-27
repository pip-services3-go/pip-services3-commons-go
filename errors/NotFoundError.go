package errors

func NewNotFoundError(correlationId, code, message string) *ApplicationError {
    return &ApplicationError{ 
        Category: NotFound,
        CorrelationId: correlationId,
        Code: code,
        Message: message,
        Status: 404,
    }
}