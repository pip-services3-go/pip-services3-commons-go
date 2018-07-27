package errors

func NewBadRequestError(correlationId, code, message string) *ApplicationError {
    return &ApplicationError{ 
        Category: BadRequest,
        CorrelationId: correlationId,
        Code: code,
        Message: message,
        Status: 400,
    }
}