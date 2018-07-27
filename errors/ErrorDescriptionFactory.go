package errors

import "fmt"

type TErrorDescriptionFactory struct{}

var ErrorDescriptionFactory = &TErrorDescriptionFactory{}

func (c *TErrorDescriptionFactory) Create(err interface{}) *ErrorDescription {
	return NewErrorDescription(err)
}

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
