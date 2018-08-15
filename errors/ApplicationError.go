package errors

type ApplicationError struct {
	Message       string                 `json:"message"`
	Category      string                 `json:"category"`
	Status        int                    `json:"status"`
	Code          string                 `json:"code"`
	Details       map[string]interface{} `json:"details"`
	CorrelationId string                 `json:"correlation_id"`
	StackTrace    string                 `json:"stack_trace"`
	Cause         string                 `json:"cause"`
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func (e *ApplicationError) WithCode(code string) *ApplicationError {
	e.Code = code
	return e
}

func (e *ApplicationError) WithStatus(status int) *ApplicationError {
	e.Status = status
	return e
}

func (e *ApplicationError) WithDetails(key string, value interface{}) *ApplicationError {
	if e.Details == nil {
		e.Details = map[string]interface{}{}
	}
	e.Details[key] = value
	return e
}

func (e *ApplicationError) WithCause(cause error) *ApplicationError {
	e.Cause = cause.Error()
	return e
}

func (e *ApplicationError) WithCauseString(cause string) *ApplicationError {
	e.Cause = cause
	return e
}

func (e *ApplicationError) WithCorrelationId(correlationId string) *ApplicationError {
	e.CorrelationId = correlationId
	return e
}

func (e *ApplicationError) Wrap(err error) *ApplicationError {
	if er, ok := err.(*ApplicationError); ok == true {
		return er
	}

	e.WithCause(err)
	return e
}

func WrapError(err error, message string) *ApplicationError {
	if e, ok := err.(*ApplicationError); ok == true {
		return e
	}

	return NewError(message).WithCause(err)
}

func NewError(message string) *ApplicationError {
	if message == "" {
		message = "Unknown error"
	}
	return &ApplicationError{Code: "UNKNOWN", Message: message, Status: 500}
}
