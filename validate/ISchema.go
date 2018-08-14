package validate

import "github.com/pip-services-go/pip-services-commons-go/errors"

type ISchema interface {
	Validate(value interface{}) []*ValidationResult
	ValidateAndReturnError(correlationId string, value interface{}, strict bool) *errors.ApplicationError
	ValidateAndThrowError(correlationId string, value interface{}, strict bool)
}
