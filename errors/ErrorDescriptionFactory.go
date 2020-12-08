package errors

import (
	"fmt"
)

/*
Factory to create serializeable ErrorDescription from ApplicationException or from arbitrary errors.
The ErrorDescriptions are used to pass errors through the wire between microservices implemented in different languages. They allow to restore exceptions on the receiving side close to the original type and preserve additional information.
see
ErrorDescription
see
ApplicationError
*/
type TErrorDescriptionFactory struct{}

var ErrorDescriptionFactory = &TErrorDescriptionFactory{}

// Creates a serializable ErrorDescription from error object.
// Parameters:
//  - err error
//  an error object
// Returns *ErrorDescription
// a serializeable ErrorDescription object that describes the error.
func (c *TErrorDescriptionFactory) Create(err interface{}) *ErrorDescription {
	return NewErrorDescription(err)
}

// Creates a serializable ErrorDescription from error object.
// Parameters:
//  - err interface{}
//  an error object
// Returns *ErrorDescription
// a serializeable ErrorDescription object that describes the error.
func NewErrorDescription(err interface{}) *ErrorDescription {
	description := &ErrorDescription{
		Category: Unknown,
		Code:     "UNKNOWN",
		Status:   500,
		Message:  "Unknown error",
	}

	ex, ok := err.(*ApplicationError)
	if ok {
		description.Category = ex.Category
		description.Status = ex.Status
		description.Code = ex.Code
		description.Message = ex.Message
		description.Details = ex.Details
		description.CorrelationId = ex.CorrelationId
		description.Cause = ex.Cause
		description.StackTrace = ex.StackTrace
	} else if err != nil {
		//description.Type = err.Name
		description.Message = fmt.Sprintf("%v", err)
		//description.StackTrace = err.Stack
	}

	return description
}
