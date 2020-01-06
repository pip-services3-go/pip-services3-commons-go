package validate

import "github.com/pip-services3-go/pip-services3-commons-go/v3/errors"

/*
Validation schema interface
*/
type ISchema interface {
	Validate(value interface{}) []*ValidationResult
	ValidateAndReturnError(correlationId string, value interface{}, strict bool) *errors.ApplicationError
	ValidateAndThrowError(correlationId string, value interface{}, strict bool)
}
