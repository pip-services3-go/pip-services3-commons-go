package errors

func NewFileError(correlationId, code, message string) *ApplicationError {
    return &ApplicationError{ 
        Category: FileError,
        CorrelationId: correlationId,
        Code: code,
        Message: message,
        Status: 500,
    }
}
