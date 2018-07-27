package errors

type TApplicationErrorFactory struct{}

var ApplicationErrorFactory *TApplicationErrorFactory = &TApplicationErrorFactory{}

func (c *TApplicationErrorFactory) Create(description *ErrorDescription) *ApplicationError {
	return NewErrorFromDescription(description)
}

func NewErrorFromDescription(description *ErrorDescription) *ApplicationError {
	if description == nil {
		return nil
	}

	var err *ApplicationError = nil
	category := description.Category
	code := description.Code
	message := description.Message
	correlationId := description.CorrelationId

	// Create well-known exception type based on error category
	if Unknown == category {
		err = NewUnknownError(correlationId, code, message)
	} else if Internal == category {
		err = NewInternalError(correlationId, code, message)
	} else if Misconfiguration == category {
		err = NewConfigError(correlationId, code, message)
	} else if NoResponse == category {
		err = NewConnectionError(correlationId, code, message)
	} else if FailedInvocation == category {
		err = NewInvocationError(correlationId, code, message)
	} else if FileError == category {
		err = NewFileError(correlationId, code, message)
	} else if BadRequest == category {
		err = NewBadRequestError(correlationId, code, message)
	} else if Unauthorized == category {
		err = NewUnauthorizedError(correlationId, code, message)
	} else if Conflict == category {
		err = NewConflictError(correlationId, code, message)
	} else if NotFound == category {
		err = NewNotFoundError(correlationId, code, message)
	} else if InvalidState == category {
		err = NewInvalidStateError(correlationId, code, message)
	} else if Unsupported == category {
		err = NewUnsupportedError(correlationId, code, message)
	} else {
		err = NewUnknownError(correlationId, code, message)
		err.Category = category
		err.Status = description.Status
	}

	// Fill error with details
	err.Details = description.Details
	err.Cause = description.Cause
	err.StackTrace = description.StackTrace

	return err
}
