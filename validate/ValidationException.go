package validate

/*
Errors in schema validation.

Validation errors are usually generated based in ValidationResult. If using strict mode, warnings will also raise validation exceptions.
*/
import (
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

// Creates a new instance of validation exception and assigns its values.
// see
// ValidationResult
// Parameters:
//  - correlationId string
//  - message string
//  a human-readable description of the error.
//  - results: []*ValidationResult
// a list of validation results
// Returns *errors.ApplicationError
func NewValidationError(correlationId string, message string, results []*ValidationResult) *errors.ApplicationError {
	if message == "" {
		message = composeErrorMessage(results)
	}
	e := errors.NewBadRequestError(correlationId, "INVALID_DATA", message)
	if results != nil && len(results) > 0 {
		e.WithDetails("results", results)
	}
	return e
}

// Composes human readable error message based on validation results.
// see
// ValidationResult
// Parameters:
//  - results []*ValidationResult
//  a list of validation results.
// Returns string
// a composed error message.
func composeErrorMessage(results []*ValidationResult) string {
	builder := strings.Builder{}
	builder.WriteString("Validation failed")

	if results == nil || len(results) == 0 {
		return builder.String()
	}

	first := true
	for _, result := range results {
		if result.Type() == Information {
			continue
		}

		if first {
			builder.WriteString(": ")
		} else {
			builder.WriteString(", ")
		}
		builder.WriteString(result.Message())
		first = false
	}

	return builder.String()
}

// Creates a new ValidationError based on errors in validation results. If validation results have no errors, than null is returned.
// see
// ValidationResult
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - results []*ValidationResult
//  list of validation results that may contain errors
//  strict boolean
//  true to treat warnings as errors.
// 	Returns *errors.ApplicationError
// a newly created ValidationException or null if no errors in found.
func NewValidationErrorFromResults(correlationId string, results []*ValidationResult, strict bool) *errors.ApplicationError {
	hasErrors := false

	for _, result := range results {
		if result.Type() == Error {
			hasErrors = true
		}

		if strict && result.Type() == Warning {
			hasErrors = true
		}
	}

	if hasErrors {
		return NewValidationError(correlationId, "", results)
	}

	return nil
}

// Throws ValidationException based on errors in validation results. If validation results have no errors, than no exception is thrown.
// see
// ValidationResult
// see
// ValidationException
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - results []*ValidationResult
//  list of validation results that may contain errors
//  - strict bool
//  true to treat warnings as errors.
func ThrowValidationErrorIfNeeded(correlationId string, results []*ValidationResult, strict bool) {
	err := NewValidationErrorFromResults(correlationId, results, strict)
	if err != nil {
		panic(err)
	}
}
